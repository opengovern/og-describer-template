package describers

import (
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"golang.org/x/net/context"
	"time"
)

func ListOrganizationRoleAssignments(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	var values []models.Resource

	// Initialize Resilient Bridge once
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	// 1) Fetch org ID
	orgID, err := fetchOrganizationID(sdk, organizationName)
	if err != nil {
		return nil, err
	}

	// 2) Fetch org roles (just enough info to get assignments)
	orgRoles, err := fetchOrganizationRoles(sdk, organizationName)
	if err != nil {
		return nil, err
	}

	// 3) For each role, gather assigned teams and users, then produce output
	var allAssignments []GitHubOrganizationRoleAssignment

	for _, role := range orgRoles {
		teams, err := fetchTeamsAssignedToRole(sdk, organizationName, role.ID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Create one JSON object per assigned team
		for _, t := range teams {
			allAssignments = append(allAssignments, GitHubOrganizationRoleAssignment{
				RoleID:         role.ID,
				OrganizationID: orgID,
				PrincipalType:  "team",
				PrincipalID:    t.ID,
			})
		}

		users, err := fetchUsersAssignedToRole(sdk, organizationName, role.ID)
		if err != nil {
			return nil, err
		}
		// Create one JSON object per assigned user
		for _, u := range users {
			allAssignments = append(allAssignments, GitHubOrganizationRoleAssignment{
				RoleID:         role.ID,
				OrganizationID: orgID,
				PrincipalType:  "user",
				PrincipalID:    u.ID,
			})
		}
	}

	for _, a := range allAssignments {
		id := fmt.Sprintf("%d/%d", a.OrganizationID, a.RoleID)
		value := models.Resource{
			ID: id,
			Description: model.OrganizationRoleAssignmentDescription{
				RoleId:         a.RoleID,
				OrganizationId: a.OrganizationID,
				Organization:   organizationName,
				PrincipalType:  a.PrincipalType,
				PrincipalId:    a.PrincipalID,
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

// OrgInfo is for GET /orgs/{org}
type OrgInfo struct {
	ID int `json:"id"`
}

// We only keep minimal information needed to fetch role assignments:
type OrgRoleForAssignment struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type orgRoleListResponseForAssignment struct {
	Roles []OrgRoleForAssignment `json:"roles"`
}

type GitHubUser struct {
	ID int `json:"id"`
}

// GitHubOrganizationRoleAssignment is the final output format:
// one object per "principal" (team or user) assigned to a role.
type GitHubOrganizationRoleAssignment struct {
	RoleID         int    `json:"role_id"`
	OrganizationID int    `json:"organization_id"`
	PrincipalType  string `json:"principal_type"` // "team" or "user"
	PrincipalID    int    `json:"principal_id"`
}

// -----------------------------------------------------
// Implementation: Org ID, Roles, Assignments
// -----------------------------------------------------

func fetchOrganizationID(sdk *resilientbridge.ResilientBridge, org string) (int, error) {
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: "/orgs/" + org,
		Headers:  map[string]string{"Accept": "application/vnd.github+json"},
	}
	resp, err := sdk.Request("github", req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode >= 400 {
		return 0, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(resp.Data))
	}
	var info OrgInfo
	if err := json.Unmarshal(resp.Data, &info); err != nil {
		return 0, fmt.Errorf("error parsing org info: %w", err)
	}
	return info.ID, nil
}

func fetchOrganizationRoles(sdk *resilientbridge.ResilientBridge, org string) ([]OrgRoleForAssignment, error) {
	var allRoles []OrgRoleForAssignment
	page := 1
	perPage := 100

	for {
		endpoint := fmt.Sprintf("/orgs/%s/organization-roles?per_page=%d&page=%d", org, perPage, page)
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
		var parsed orgRoleListResponseForAssignment
		if err := json.Unmarshal(resp.Data, &parsed); err != nil {
			return nil, fmt.Errorf("JSON parse error (page %d): %w", page, err)
		}
		allRoles = append(allRoles, parsed.Roles...)
		if len(parsed.Roles) < perPage {
			break
		}
		page++
	}
	return allRoles, nil
}

func fetchTeamsAssignedToRole(sdk *resilientbridge.ResilientBridge, org string, roleID int) ([]GitHubTeam, error) {
	var allTeams []GitHubTeam
	page := 1
	perPage := 100
	for {
		endpoint := fmt.Sprintf("/orgs/%s/organization-roles/%d/teams?per_page=%d&page=%d", org, roleID, perPage, page)
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github+json"},
		}
		resp, err := sdk.Request("github", req)
		if err != nil {
			return nil, fmt.Errorf("fetch teams (role %d, page %d): %w", roleID, page, err)
		}
		if resp.StatusCode == 404 {
			return nil, fmt.Errorf("404 for teams on role %d", roleID)
		} else if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d fetching teams for role %d: %s", resp.StatusCode, roleID, string(resp.Data))
		}

		var batch []GitHubTeam
		if err := json.Unmarshal(resp.Data, &batch); err != nil {
			return nil, fmt.Errorf("JSON parse error (role %d, page %d): %w", roleID, page, err)
		}
		allTeams = append(allTeams, batch...)
		if len(batch) < perPage {
			break
		}
		page++
	}
	return allTeams, nil
}

func fetchUsersAssignedToRole(sdk *resilientbridge.ResilientBridge, org string, roleID int) ([]GitHubUser, error) {
	var allUsers []GitHubUser
	page := 1
	perPage := 100
	for {
		endpoint := fmt.Sprintf("/orgs/%s/organization-roles/%d/users?per_page=%d&page=%d", org, roleID, perPage, page)
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers:  map[string]string{"Accept": "application/vnd.github+json"},
		}
		resp, err := sdk.Request("github", req)
		if err != nil {
			return nil, fmt.Errorf("fetch users (role %d, page %d): %w", roleID, page, err)
		}
		if resp.StatusCode == 404 {
			return nil, fmt.Errorf("404 for users on role %d", roleID)
		} else if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d fetching users for role %d: %s", resp.StatusCode, roleID, string(resp.Data))
		}

		var batch []GitHubUser
		if err := json.Unmarshal(resp.Data, &batch); err != nil {
			return nil, fmt.Errorf("JSON parse error (role %d, page %d): %w", roleID, page, err)
		}
		allUsers = append(allUsers, batch...)
		if len(batch) < perPage {
			break
		}
		page++
	}
	return allUsers, nil
}
