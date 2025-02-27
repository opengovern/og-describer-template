package describers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"log"
	"strconv"
)

func GetAllTeamsMembers(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	var values []models.Resource

	//
	allTeams, err := fetchAllTeams(sdk, organizationName, githubClient.Token)
	if err != nil {
		log.Fatalf("fetchAllTeams => %v", err)
	}
	if len(allTeams) == 0 {
		return nil, nil
	}

	// Attempt to get orgID from the first team, fallback if missing
	orgID := allTeams[0].Organization.ID
	if orgID == 0 {
		orgID, err = fetchOrganizationID(sdk, organizationName)
		if err != nil {
			log.Fatalf("fetchOrganizationID => %v", err)
		}
	}

	// For each team, fetch members, stream them
	for _, t := range allTeams {
		members, err := ListTeamMembers(sdk, organizationName, orgID, &t, githubClient.Token)
		if err != nil {
			log.Printf("Warning: ListTeamMembers(%s) => %v", t.Slug, err)
			continue
		}
		for _, m := range members {
			id := fmt.Sprintf("%s/%s", m.TeamID, m.MemberPrincipalID)
			value := models.Resource{
				ID: id,
				Description: model.TeamMemberDescription{
					Organization:        organizationName,
					TeamID:              m.TeamID,
					MemberPrincipalID:   m.MemberPrincipalID,
					MemberPrincipalType: m.MemberPrincipalType,
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
	}
	return values, nil
}

// GitHubTeam => minimal info from /orgs/:org/teams or /orgs/:org/teams/:slug
// plus the optional "organization.id" if your token has permission
type GitHubTeam struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`

	// If authorized, GitHub returns the organization object here
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
}

// GitHubTeamMemberRaw => raw data from GET /orgs/:org/teams/:slug/members
type GitHubTeamMemberRaw struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
	Type  string `json:"type"` // typically "User"
}

// TeamMember => final structure we print (IDs as strings, no org or slug)
type TeamMember struct {
	TeamID              string `json:"team-id"`      // numeric team ID as string
	MemberPrincipalType string `json:"member-type"`  // always "User"
	MemberPrincipalID   string `json:"principal-id"` // numeric user ID as string
}

// fetchAllTeams => GET /orgs/{org}/teams
func fetchAllTeams(sdk *resilientbridge.ResilientBridge, org, token string) ([]GitHubTeam, error) {
	page := 1
	perPage := 100
	var all []GitHubTeam

	for {
		endpoint := fmt.Sprintf("/orgs/%s/teams?per_page=%d&page=%d", org, perPage, page)
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers: map[string]string{
				"Accept":        "application/vnd.github+json",
				"Authorization": "Bearer " + token,
			},
		}
		resp, err := requestWithChecks(sdk, req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d => %s", resp.StatusCode, string(resp.Data))
		}

		var batch []GitHubTeam
		if err := json.Unmarshal(resp.Data, &batch); err != nil {
			return nil, fmt.Errorf("unmarshal teams => %w", err)
		}
		all = append(all, batch...)

		if len(batch) < perPage {
			break
		}
		page++
	}
	return all, nil
}

// fetchTeamBySlug => GET /orgs/{org}/teams/{slug}
func fetchTeamBySlug(sdk *resilientbridge.ResilientBridge, org, teamSlug, token string) (*GitHubTeam, error) {
	if teamSlug == "" {
		return nil, errors.New("team slug cannot be empty")
	}
	endpoint := fmt.Sprintf("/orgs/%s/teams/%s", org, teamSlug)
	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: endpoint,
		Headers: map[string]string{
			"Accept":        "application/vnd.github+json",
			"Authorization": "Bearer " + token,
		},
	}

	resp, err := requestWithChecks(sdk, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("team %q not found or insufficient permission", teamSlug)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d => %s", resp.StatusCode, string(resp.Data))
	}

	var team GitHubTeam
	if err := json.Unmarshal(resp.Data, &team); err != nil {
		return nil, fmt.Errorf("unmarshal team => %w", err)
	}
	return &team, nil
}

// ListTeamMembers => fetch a single team's members
// orgID is used only to reduce calls if we need fallback. It's not output.
func ListTeamMembers(
	sdk *resilientbridge.ResilientBridge,
	org string,
	orgID int,
	team *GitHubTeam,
	token string,
) ([]TeamMember, error) {

	if team == nil {
		return nil, errors.New("ListTeamMembers called with nil team")
	}

	page := 1
	perPage := 100
	var final []TeamMember

	for {
		endpoint := fmt.Sprintf("/orgs/%s/teams/%s/members?per_page=%d&page=%d", org, team.Slug, perPage, page)
		req := &resilientbridge.NormalizedRequest{
			Method:   "GET",
			Endpoint: endpoint,
			Headers: map[string]string{
				"Accept":        "application/vnd.github+json",
				"Authorization": "Bearer " + token,
			},
		}

		resp, err := requestWithChecks(sdk, req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode == 404 {
			return nil, fmt.Errorf("team %q not found or insufficient permission", team.Slug)
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d => %s", resp.StatusCode, string(resp.Data))
		}

		var batch []GitHubTeamMemberRaw
		if err := json.Unmarshal(resp.Data, &batch); err != nil {
			return nil, fmt.Errorf("unmarshal members => %w", err)
		}

		// Transform raw => final (IDs as strings, no org/slug in output)
		for _, u := range batch {
			member := TeamMember{
				TeamID:              strconv.Itoa(team.ID),
				MemberPrincipalType: "User",
				MemberPrincipalID:   strconv.Itoa(u.ID),
			}
			final = append(final, member)
		}

		if len(batch) < perPage {
			break
		}
		page++
	}

	return final, nil
}
