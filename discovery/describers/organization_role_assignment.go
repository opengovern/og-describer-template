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

	// 3) For each role, gather assigned teams and users
	var assignments []GitHubRoleAssignment

	for _, role := range orgRoles {
		teams, err := fetchTeamsAssignedToRole(sdk, organizationName, role.ID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		teamIDs := make([]int, 0, len(teams))
		for _, t := range teams {
			teamIDs = append(teamIDs, t.ID)
		}

		users, err := fetchUsersAssignedToRole(sdk, organizationName, role.ID)
		if err != nil {
			return nil, err
		}
		userIDs := make([]int, 0, len(users))
		for _, u := range users {
			userIDs = append(userIDs, u.ID)
		}

		assignments = append(assignments, GitHubRoleAssignment{
			RoleID:         role.ID,
			OrganizationID: orgID,
			CreatedAt:      role.CreatedAt,
			UpdatedAt:      role.UpdatedAt,
			ListOfTeams:    teamIDs,
			ListOfUsers:    userIDs,
		})
	}

	for _, a := range assignments {
		id := fmt.Sprintf("%d/%d", a.OrganizationID, a.RoleID)
		value := models.Resource{
			ID: id,
			Description: model.OrganizationRoleAssignmentDescription{
				RoleId:         a.RoleID,
				OrganizationId: a.OrganizationID,
				Organization:   organizationName,
				ListOfTeams:    a.ListOfTeams,
				ListOfUsers:    a.ListOfUsers,
				CreatedAt:      a.CreatedAt,
				UpdatedAt:      a.UpdatedAt,
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

// GitHubRoleAssignment captures which teams and which users are assigned to a single org role
type GitHubRoleAssignment struct {
	RoleID         int       `json:"role_id"`
	OrganizationID int       `json:"organization_id"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`

	ListOfTeams []int `json:"list_of_teams"`
	ListOfUsers []int `json:"list_of_users"`
}

// Minimal structs for teams and users assigned to roles
type GitHubTeam struct {
	ID int `json:"id"`
}

type GitHubUser struct {
	ID int `json:"id"`
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
