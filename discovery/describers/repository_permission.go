package describers

import (
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"golang.org/x/net/context"
	"strconv"
	"strings"
)

func ListRepositoryPermissions(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	var values []models.Resource

	// Build ResilientBridge
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	// 1) Collect target repos (single or multiple)
	repos, err := gatherRepos(sdk, organizationName, "", false)
	if err != nil {
		return nil, err
	}

	// 2) Filter out archived/disabled, skip forks as requested
	finalRepos, err := filterRepos(sdk, repos, organizationName, true, false, false)
	if err != nil {
		return nil, err
	}

	if len(finalRepos) == 0 {
		return nil, nil
	}

	// 3) For each repo, fetch Principals
	var allPrincipals []Principal
	for _, r := range finalRepos {
		principals, err := fetchRepoPermissions(sdk, r)
		if err != nil {
			fmt.Println(err)
			continue
		}
		allPrincipals = append(allPrincipals, principals...)
	}

	for _, p := range allPrincipals {
		value := models.Resource{
			ID:   strconv.FormatInt(int64(p.PrincipalID), 10),
			Name: p.PrincipalName,
			Description: model.RepositoryPermissionDescription{
				PrincipalName:      p.PrincipalName,
				PrincipalId:        p.PrincipalID,
				PrincipalType:      p.PrincipalType,
				RepositoryName:     p.RepositoryName,
				RepositoryFullName: p.RepositoryFullName,
				RepositoryId:       p.RepositoryID,
				Permissions:        p.Permissions,
				RoleName:           p.RoleName,
				Organization:       organizationName,
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

// --------------------------------------------------
// Minimal structures (Repo, RepoInfo, etc.)
// --------------------------------------------------

type RepoInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

type RepoTeam struct {
	Name        string          `json:"name"`
	ID          int             `json:"id"`
	Description string          `json:"description"`
	Privacy     string          `json:"privacy"`
	Permission  string          `json:"permission"`
	Permissions map[string]bool `json:"permissions"`
}

// --------------------------------------------------
// Principal => final model
// --------------------------------------------------

type Principal struct {
	PrincipalName string `json:"principal_name"`
	PrincipalID   int    `json:"principal_id"`
	PrincipalType string `json:"principal_type"` // "team" or "user"

	RepositoryName     string `json:"repository_name"`
	RepositoryFullName string `json:"repository_full_name"`
	RepositoryID       int    `json:"repository_id"`

	Permissions *map[string]bool `json:"permissions"`
	RoleName    *string          `json:"role_name"`
}

// --------------------------------------------------
// Step 3: For each Repo, gather Principals
// --------------------------------------------------

func fetchRepoPermissions(sdk *resilientbridge.ResilientBridge, r Repo) ([]Principal, error) {
	// 1) Get the ID + full_name from GET /repos/{org}/{repo}
	info, err := fetchRepoInfo(sdk, r.Owner.Login, r.Name)
	if err != nil {
		return nil, fmt.Errorf("fetchRepoInfo: %w", err)
	}

	// 2) Teams, collaborators, contributors
	teams, err := fetchRepoTeams(sdk, r.Owner.Login, r.Name)
	if err != nil {
		return nil, fmt.Errorf("fetchRepoTeams: %w", err)
	}
	collabs, err := fetchRepoCollaborators(sdk, r.Owner.Login, r.Name)
	if err != nil {
		return nil, fmt.Errorf("fetchRepoCollaborators: %w", err)
	}
	conts, err := fetchRepoContributors(sdk, r.Owner.Login, r.Name)
	if err != nil {
		return nil, fmt.Errorf("fetchRepoContributors: %w", err)
	}

	// 3) Convert to Principals
	var principals []Principal

	// (A) Teams -> Principals
	for _, t := range teams {
		roleName := t.Permission // e.g. "admin", "triage", "read"
		perms := t.Permissions
		p := Principal{
			PrincipalName:      t.Name,
			PrincipalID:        t.ID,
			PrincipalType:      "team",
			RepositoryName:     info.Name,
			RepositoryFullName: info.FullName,
			RepositoryID:       info.ID,
			Permissions:        &perms,
			RoleName:           &roleName,
		}
		principals = append(principals, p)
	}

	// (B) Collaborators -> Principals
	collabMap := make(map[string]bool)
	for _, c := range collabs {
		collabMap[strings.ToLower(c.Login)] = true
	}
	for _, c := range collabs {
		roleName := deriveRoleName(c.Permissions)
		perms := c.Permissions
		var rn *string
		if roleName != "" && roleName != "none" {
			rn = &roleName
		}
		p := Principal{
			PrincipalName:      c.Login,
			PrincipalID:        c.ID,
			PrincipalType:      "user",
			RepositoryName:     info.Name,
			RepositoryFullName: info.FullName,
			RepositoryID:       info.ID,
			Permissions:        &perms,
			RoleName:           rn,
		}
		principals = append(principals, p)
	}

	// (C) Contributors -> Principals
	// If not a collaborator, then "permissions"=null, "role_name"=null
	for _, ct := range conts {
		lower := strings.ToLower(ct.Login)
		if collabMap[lower] {
			continue
		}
		p := Principal{
			PrincipalName:      ct.Login,
			PrincipalID:        ct.ID,
			PrincipalType:      "user",
			RepositoryName:     info.Name,
			RepositoryFullName: info.FullName,
			RepositoryID:       info.ID,
			Permissions:        nil,
			RoleName:           nil,
		}
		principals = append(principals, p)
	}

	return principals, nil
}

func fetchRepoInfo(sdk *resilientbridge.ResilientBridge, org, repo string) (*RepoInfo, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s", org, repo)
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  map[string]string{"Accept": "application/vnd.github.v3+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d => %s", resp.StatusCode, string(resp.Data))
	}
	var info RepoInfo
	if e := json.Unmarshal(resp.Data, &info); e != nil {
		return nil, fmt.Errorf("unmarshal RepoInfo => %w", e)
	}
	return &info, nil
}

// --------------------------------------------------
// fetchRepoTeams, fetchRepoCollaborators, fetchRepoContributors
// --------------------------------------------------

func fetchRepoTeams(sdk *resilientbridge.ResilientBridge, org, repo string) ([]RepoTeam, error) {
	var all []RepoTeam
	page := 1
	perPage := 100
	for {
		endpoint := fmt.Sprintf("/repos/%s/%s/teams?per_page=%d&page=%d", org, repo, perPage, page)
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github.v3+json"},
		}
		resp, err := sdk.Request("github", req)
		if err != nil {
			return nil, fmt.Errorf("page %d teams error: %w", page, err)
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d on page %d: %s", resp.StatusCode, page, string(resp.Data))
		}
		var batch []RepoTeam
		if err := json.Unmarshal(resp.Data, &batch); err != nil {
			return nil, fmt.Errorf("unmarshal teams (page %d): %w", page, err)
		}
		if len(batch) == 0 {
			break
		}
		all = append(all, batch...)
		if len(batch) < perPage {
			break
		}
		page++
	}
	return all, nil
}

// --------------------------------------------------
// deriveRoleName
// --------------------------------------------------

func deriveRoleName(perms map[string]bool) string {
	switch {
	case perms["admin"]:
		return "admin"
	case perms["maintain"]:
		return "maintain"
	case perms["push"]:
		return "write"
	case perms["triage"]:
		return "triage"
	case perms["pull"]:
		return "read"
	default:
		return "none" // no recognized role
	}
}
