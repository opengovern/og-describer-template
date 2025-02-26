package describers

import (
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"golang.org/x/net/context"
	"log"
	"time"
)

func ListOrganizationRoleDefinitions(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	var values []models.Resource

	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	// Call our new function: ListRoleDefinition
	roleDefs, err := ListRoleDefinition(sdk, organizationName, false)
	if err != nil {
		log.Fatalf("Error listing role definitions: %v", err)
	}

	for _, a := range roleDefs {
		id := fmt.Sprintf("%d/%d", a.OrganizationID, a.ID)
		value := models.Resource{
			ID: id,
			Description: model.OrganizationRoleDefinitionDescription{
				Id:             a.ID,
				Name:           a.Name,
				Description:    a.Description,
				Permissions:    a.Permissions,
				OrganizationId: a.OrganizationID,
				Organization:   organizationName,
				CreatedAt:      a.CreatedAt,
				UpdatedAt:      a.UpdatedAt,
				Source:         a.Source,
				BaseRole:       a.BaseRole,
				Type:           a.Type,
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

// -----------------------------------------------------
// Data Structures
// -----------------------------------------------------

// GitHubRoleDefinition is our unified data model for **role definitions**
type GitHubRoleDefinition struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Permissions    []string  `json:"permissions"`
	OrganizationID int       `json:"organization_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Source         string    `json:"source"`
	BaseRole       *string   `json:"base_role"`
	Type           string    `json:"type"` // "organization-roles" or "custom-repository-roles"
}

// OrgRoleAPIResponse => from GET /orgs/{org}/organization-roles
type OrgRoleAPIResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Permissions []string  `json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Source      string    `json:"source"`
	BaseRole    *string   `json:"base_role"`
}

// OrgRoleListAPIResponse => top-level for org roles
type OrgRoleListAPIResponse struct {
	TotalCount int                  `json:"total_count"`
	Roles      []OrgRoleAPIResponse `json:"roles"`
}

// RepoRoleAPIResponse => from GET /orgs/{org}/custom-repository-roles
type RepoRoleAPIResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	BaseRole    string    `json:"base_role"`
	Permissions []string  `json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// RepoRoleListAPIResponse => top-level for custom repo roles
type RepoRoleListAPIResponse struct {
	TotalCount  int                   `json:"total_count"`
	CustomRoles []RepoRoleAPIResponse `json:"custom_roles"`
}

// -----------------------------------------------------
// ListRoleDefinition: fetch org ID, org roles, custom repo roles,
// unify them and return a single slice of GitHubRoleDefinition
// -----------------------------------------------------
func ListRoleDefinition(sdk *resilientbridge.ResilientBridge, org string, debug bool) ([]GitHubRoleDefinition, error) {
	// 1) Fetch org ID
	orgID, err := fetchOrganizationID(sdk, org)
	if err != nil {
		return nil, fmt.Errorf("fetchOrganizationID: %w", err)
	}

	// 2) Fetch organization-level roles
	orgRoles, err := fetchOrganizationRoleDefinitions(sdk, org, debug)
	if err != nil {
		return nil, fmt.Errorf("fetchOrganizationRoles: %w", err)
	}

	// 3) Fetch custom repository-level roles
	repoRoles, err := fetchCustomRepositoryRoles(sdk, org, debug)
	if err != nil {
		return nil, fmt.Errorf("fetchCustomRepositoryRoles: %w", err)
	}

	// 4) Merge them into a single slice of definitions
	final := unifyRoleDefinitions(orgRoles, repoRoles, orgID)

	if debug {
		log.Printf("Finished building %d total role definitions (orgRoles=%d, repoRoles=%d).",
			len(final), len(orgRoles), len(repoRoles))
	}

	return final, nil
}

func fetchOrganizationRoleDefinitions(sdk *resilientbridge.ResilientBridge, org string, debug bool) ([]OrgRoleAPIResponse, error) {
	if debug {
		log.Printf("Fetching org roles from /orgs/%s/organization-roles ...", org)
	}
	var allRoles []OrgRoleAPIResponse
	page := 1
	perPage := 100

	for {
		endpoint := fmt.Sprintf("/orgs/%s/organization-roles?per_page=%d&page=%d", org, perPage, page)
		if debug {
			log.Printf("Fetching org roles page %d ...", page)
		}
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github+json"},
		}
		resp, err := sdk.Request("github", req)
		if err != nil {
			return nil, fmt.Errorf("error on page %d: %w", page, err)
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d on page %d: %s", resp.StatusCode, page, string(resp.Data))
		}

		var parsed OrgRoleListAPIResponse
		if err := json.Unmarshal(resp.Data, &parsed); err != nil {
			return nil, fmt.Errorf("JSON parse error (page %d): %w", page, err)
		}
		allRoles = append(allRoles, parsed.Roles...)

		if debug {
			log.Printf("Fetched %d org roles on page %d", len(parsed.Roles), page)
		}

		if len(parsed.Roles) < perPage {
			// no more data
			break
		}
		page++
	}

	if debug {
		log.Printf("Total org roles fetched: %d", len(allRoles))
	}

	return allRoles, nil
}

// -----------------------------------------------------
// Helper Functions: fetch org ID, org roles, repo roles
// -----------------------------------------------------

func fetchCustomRepositoryRoles(sdk *resilientbridge.ResilientBridge, org string, debug bool) ([]RepoRoleAPIResponse, error) {
	if debug {
		log.Printf("Fetching custom repository roles from /orgs/%s/custom-repository-roles ...", org)
	}
	var allRepoRoles []RepoRoleAPIResponse
	page := 1
	perPage := 100

	for {
		endpoint := fmt.Sprintf("/orgs/%s/custom-repository-roles?per_page=%d&page=%d", org, perPage, page)
		if debug {
			log.Printf("Fetching custom repo roles page %d ...", page)
		}
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github+json"},
		}
		resp, err := sdk.Request("github", req)
		if err != nil {
			return nil, fmt.Errorf("error on page %d: %w", page, err)
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d on page %d: %s", resp.StatusCode, page, string(resp.Data))
		}

		var parsed RepoRoleListAPIResponse
		if err := json.Unmarshal(resp.Data, &parsed); err != nil {
			return nil, fmt.Errorf("JSON parse error (page %d): %w", page, err)
		}
		allRepoRoles = append(allRepoRoles, parsed.CustomRoles...)

		if debug {
			log.Printf("Fetched %d custom repo roles on page %d", len(parsed.CustomRoles), page)
		}

		if len(parsed.CustomRoles) < perPage {
			// no more data
			break
		}
		page++
	}

	if debug {
		log.Printf("Total custom repo roles fetched: %d", len(allRepoRoles))
	}
	return allRepoRoles, nil
}

// -----------------------------------------------------
// unifyRoleDefinitions merges org-level & custom repo-level roles
// -----------------------------------------------------
func unifyRoleDefinitions(
	orgRoles []OrgRoleAPIResponse,
	repoRoles []RepoRoleAPIResponse,
	orgID int,
) []GitHubRoleDefinition {

	var results []GitHubRoleDefinition

	// A) Organization roles
	for _, r := range orgRoles {
		results = append(results, GitHubRoleDefinition{
			ID:             r.ID,
			Name:           r.Name,
			Description:    r.Description,
			Permissions:    r.Permissions,
			OrganizationID: orgID,
			CreatedAt:      r.CreatedAt,
			UpdatedAt:      r.UpdatedAt,
			Source:         r.Source,   // from API
			BaseRole:       r.BaseRole, // can be nil
			Type:           "organization-roles",
		})
	}

	// B) Custom Repository roles
	for _, r := range repoRoles {
		baseRolePtr := &r.BaseRole
		results = append(results, GitHubRoleDefinition{
			ID:             r.ID,
			Name:           r.Name,
			Description:    r.Description,
			Permissions:    r.Permissions,
			OrganizationID: orgID,
			CreatedAt:      r.CreatedAt,
			UpdatedAt:      r.UpdatedAt,
			Source:         "", // not provided for custom repo roles
			BaseRole:       baseRolePtr,
			Type:           "custom-repository-roles",
		})
	}

	return results
}
