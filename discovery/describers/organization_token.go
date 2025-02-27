package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

func ListOrganizationTokens(ctx context.Context,
	githubClient model.GitHubClient,
	organizationName string,
	stream *models.StreamSender) ([]models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	var values []models.Resource

	orgID, err := fetchOrganizationID(sdk, organizationName)
	if err != nil {
		return nil, fmt.Errorf("fetchOrganizationID => %w", err)
	}

	// 2) Check if SAML enabled
	saml, err := isSamlSsoEnabled(sdk, organizationName, githubClient.Token)
	if err != nil {
		return nil, fmt.Errorf("isSamlSsoEnabled => %w", err)
	}
	if !saml {
		return nil, fmt.Errorf("SAML not enabled for org=%s", organizationName)
	}

	// 3) List SSO Auth
	rawAuths, err := listSsoAuthorizations(sdk, organizationName, githubClient.Token)
	if err != nil {
		return nil, fmt.Errorf("listSsoAuthorizations => %w", err)
	}

	// 4) Transform => final
	finalList, err := transformAll(sdk, organizationName, githubClient.Token, orgID, rawAuths)
	if err != nil {
		return nil, fmt.Errorf("transformAll => %w", err)
	}

	for _, r := range finalList {
		name := ""
		if r.Title != nil {
			name = *r.Title
		}
		value := models.Resource{
			ID:   strconv.Itoa(int(r.AuthorizedCredentialID)),
			Name: name,
			Description: model.OrganizationTokenDescription{
				AuthorizedCredentialId:        r.AuthorizedCredentialID,
				Title:                         r.Title,
				OrganizationID:                r.OrganizationID,
				Organization:                  organizationName,
				PrincipalID:                   r.PrincipleID,
				PrincipleType:                 r.PrincipleType,
				AuthorizedCredentialExpiresAt: r.AuthorizedCredentialExpiresAt,
				Login:                         r.Login,
				Scopes:                        r.Scopes,
				CredentialId:                  r.CredentialID,
				CredentialType:                r.CredentialType,
				CredentialAccessedAt:          r.CredentialAccessedAt,
				CredentialAuthorizedAt:        r.CredentialAuthorizedAt,
				TokenLastEight:                r.TokenLastEight,
				Fingerprint:                   r.Fingerprint,
			},
		}
		if stream != nil {
			if err := (*stream)(value); err != nil {
				return nil, err
			}
		} else {
			values = append(values, value)
		}
	}

	return values, nil
}

// ----------------------------------------------------
// Data Structures
// ----------------------------------------------------

// rawSsoAuth matches the raw GitHub response from
// GET /orgs/{org}/credential-authorizations
type rawSsoAuth struct {
	Login                         string     `json:"login"`
	CredentialID                  int64      `json:"credential_id"`
	CredentialType                string     `json:"credential_type"`  // "personal access token", "ssh key"
	TokenLastEight                *string    `json:"token_last_eight"` // only if personal access token
	CredentialAuthorizedAt        *time.Time `json:"credential_authorized_at"`
	Scopes                        []string   `json:"scopes"`
	Fingerprint                   *string    `json:"fingerprint"` // only if SSH key
	CredentialAccessedAt          *time.Time `json:"credential_accessed_at"`
	AuthorizedCredentialID        int64      `json:"authorized_credential_id"`
	AuthorizedCredentialTitle     *string    `json:"authorized_credential_title"`
	AuthorizedCredentialExpiresAt *time.Time `json:"authorized_credential_expires_at"`
}

// AuthorizedCredential => final JSON with new fields:
//   - organization_id (numeric org ID as string)
//   - principle_type (always "user")
//   - principle_id (numeric user ID as string)
type AuthorizedCredential struct {
	Login          string `json:"login"`
	PrincipleType  string `json:"principle_type"`  // "user"
	OrganizationID string `json:"organization_id"` // numeric org ID as string
	PrincipleID    string `json:"principle_id"`    // numeric user ID as string

	CredentialID                  int64      `json:"credential_id"`
	CredentialType                string     `json:"credential_type"` // "token" or "ssh-key"
	TokenLastEight                *string    `json:"token_last_eight"`
	CredentialAuthorizedAt        *time.Time `json:"credential_authorized_at"`
	Scopes                        []string   `json:"scopes"`
	Fingerprint                   *string    `json:"fingerprint"`
	CredentialAccessedAt          *time.Time `json:"credential_accessed_at"`
	AuthorizedCredentialID        int64      `json:"authorized_credential_id"`
	AuthorizedCredentialExpiresAt *time.Time `json:"authorized_credential_expires_at"`

	Title *string `json:"title"`
}

// Minimal user info from GET /users/{login}
type gitHubUser struct {
	ID int `json:"id"`
}

// ----------------------------------------------------
// Implementation
// ----------------------------------------------------

func newBridge(token string) *resilientbridge.ResilientBridge {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})
	return sdk
}

// isSamlSsoEnabled => do a quick test
func isSamlSsoEnabled(sdk *resilientbridge.ResilientBridge, org, token string) (bool, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: fmt.Sprintf("/orgs/%s/credential-authorizations?per_page=1", org),
		Headers: map[string]string{
			"Accept":        "application/vnd.github+json",
			"Authorization": "Bearer " + token,
		},
	}
	resp, err := requestWithChecks(sdk, req)
	if err != nil {
		return false, err
	}

	switch resp.StatusCode {
	case 200:
		return true, nil
	case 404:
		return false, nil
	case 403:
		return false, fmt.Errorf("403 => must be org owner or missing scope admin:org")
	default:
		return false, fmt.Errorf("HTTP %d => %s", resp.StatusCode, string(resp.Data))
	}
}

// listSsoAuthorizations => fetch all
func listSsoAuthorizations(sdk *resilientbridge.ResilientBridge, org, token string) ([]rawSsoAuth, error) {
	var all []rawSsoAuth
	page := 1
	perPage := 100

	for {
		ep := fmt.Sprintf("/orgs/%s/credential-authorizations?per_page=%d&page=%d", org, perPage, page)
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: ep,
			Headers: map[string]string{
				"Accept":        "application/vnd.github+json",
				"Authorization": "Bearer " + token,
			},
		}
		resp, err := requestWithChecks(sdk, req)
		if err != nil {
			return nil, fmt.Errorf("page %d => %w", page, err)
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d => %s", resp.StatusCode, string(resp.Data))
		}

		var batch []rawSsoAuth
		if err := json.Unmarshal(resp.Data, &batch); err != nil {
			return nil, fmt.Errorf("unmarshal page %d => %w", page, err)
		}
		all = append(all, batch...)

		if len(batch) < perPage {
			break
		}
		page++
	}

	return all, nil
}

// transformAll => for each raw record, fetch user ID, fill in org ID
func transformAll(sdk *resilientbridge.ResilientBridge, org, token string, orgID int, rawAuths []rawSsoAuth) ([]AuthorizedCredential, error) {
	var finalList []AuthorizedCredential

	// We'll do login->userID lookups once
	userCache := make(map[string]string)
	var mu sync.Mutex

	fetchUserIDCached := func(login string) string {
		mu.Lock()
		val, found := userCache[login]
		mu.Unlock()
		if found {
			return val
		}
		// Do actual fetch
		numeric, err := fetchUserID(sdk, login, token)
		if err != nil {
			log.Printf("Warning: cannot fetch user ID for login=%q => %v", login, err)
			return ""
		}
		userIDStr := fmt.Sprintf("%d", numeric)
		mu.Lock()
		userCache[login] = userIDStr
		mu.Unlock()
		return userIDStr
	}

	orgIDStr := fmt.Sprintf("%d", orgID)

	for _, r := range rawAuths {
		userID := fetchUserIDCached(r.Login)
		ctype := mapCredentialType(r.CredentialType)

		var mergedTitle *string
		if r.AuthorizedCredentialTitle != nil && *r.AuthorizedCredentialTitle != "" {
			mergedTitle = r.AuthorizedCredentialTitle
		}

		final := AuthorizedCredential{
			Login:          r.Login,
			PrincipleType:  "user",
			OrganizationID: orgIDStr,
			PrincipleID:    userID,

			CredentialID:                  r.CredentialID,
			CredentialType:                ctype,
			TokenLastEight:                r.TokenLastEight,
			CredentialAuthorizedAt:        r.CredentialAuthorizedAt,
			Scopes:                        r.Scopes,
			Fingerprint:                   r.Fingerprint,
			CredentialAccessedAt:          r.CredentialAccessedAt,
			AuthorizedCredentialID:        r.AuthorizedCredentialID,
			AuthorizedCredentialExpiresAt: r.AuthorizedCredentialExpiresAt,
			Title:                         mergedTitle,
		}
		finalList = append(finalList, final)
	}

	return finalList, nil
}

// fetchUserID => /users/{login} => numeric ID
func fetchUserID(sdk *resilientbridge.ResilientBridge, login, token string) (int, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: "/users/" + login,
		Headers: map[string]string{
			"Accept":        "application/vnd.github+json",
			"Authorization": "Bearer " + token,
		},
	}
	resp, err := requestWithChecks(sdk, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode == 404 {
		return 0, fmt.Errorf("user not found: %s", login)
	}
	if resp.StatusCode >= 400 {
		return 0, fmt.Errorf("HTTP %d => %s", resp.StatusCode, string(resp.Data))
	}

	var user gitHubUser
	if err := json.Unmarshal(resp.Data, &user); err != nil {
		return 0, fmt.Errorf("unmarshal user => %w", err)
	}
	return user.ID, nil
}

// mapCredentialType => convert "personal access token" => "token", "ssh key" => "ssh-key"
func mapCredentialType(s string) string {
	low := strings.ToLower(s)
	if strings.Contains(low, "personal access token") {
		return "token"
	} else if strings.Contains(low, "ssh key") {
		return "ssh-key"
	}
	return s
}

// requestWithChecks => detect 403 vs rate limit
func requestWithChecks(sdk *resilientbridge.ResilientBridge, req *resilientbridge.NormalizedRequest) (*resilientbridge.NormalizedResponse, error) {
	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 403 || resp.StatusCode == 429 {
		remain := resp.Headers["X-RateLimit-Remaining"]
		if remain == "0" {
			return nil, fmt.Errorf("rate limit exceeded")
		}
		return nil, fmt.Errorf("HTTP %d => permission or SAML restriction. Body=%s",
			resp.StatusCode, string(resp.Data))
	}
	return resp, nil
}
