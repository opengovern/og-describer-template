package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"strconv"
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

	endpoint := fmt.Sprintf("/orgs/%s/credential-authorizations", organizationName)
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}

	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, fmt.Errorf("error fetching repos: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(resp.Data))
	}

	// Decode into a slice of generic maps. We'll only parse name, archived, disabled, etc.
	var tokensResponse []OrganizationToken
	if err := json.Unmarshal(resp.Data, &tokensResponse); err != nil {
		return nil, fmt.Errorf("error decoding repos list: %w", err)
	}

	fmt.Println(tokensResponse)

	for _, r := range tokensResponse {
		value := models.Resource{
			ID:   strconv.Itoa(int(r.AuthorizedCredentialId)),
			Name: r.AuthorizedCredentialTitle,
			Description: model.OrganizationTokenDescription{
				AuthorizedCredentialId:        r.AuthorizedCredentialId,
				AuthorizedCredentialTitle:     r.AuthorizedCredentialTitle,
				AuthorizedCredentialNote:      r.AuthorizedCredentialNote,
				AuthorizedCredentialExpiresAt: r.AuthorizedCredentialExpiresAt,
				Login:                         r.Login,
				Scopes:                        r.Scopes,
				CredentialId:                  r.CredentialId,
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

type OrganizationToken struct {
	AuthorizedCredentialId        int64     `json:"authorized_credential_id"`
	AuthorizedCredentialTitle     string    `json:"authorized_credential_title"`
	AuthorizedCredentialNote      string    `json:"authorized_credential_note"`
	AuthorizedCredentialExpiresAt time.Time `json:"authorized_credential_expires_at"`
	Login                         string    `json:"login"`
	Scopes                        []string  `json:"scopes"`
	CredentialId                  int64     `json:"credential_id"`
	CredentialType                string    `json:"credential_type"`
	CredentialAccessedAt          time.Time `json:"credential_accessed_at"`
	CredentialAuthorizedAt        time.Time `json:"credential_authorized_at"`
	TokenLastEight                string    `json:"token_last_eight"` // Change from time.Time to string
	Fingerprint                   string    `json:"fingerprint"`
}
