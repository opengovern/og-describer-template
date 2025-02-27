package describers

import (
	"context"
	"encoding/json"
	"fmt"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"log"
	"net/http"
	"strconv"

	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func GetOrganizationTeamList(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(githubClient.Token), &resilientbridge.ProviderConfig{
		UseProviderLimits: true,
		MaxRetries:        3,
		BaseBackoff:       0,
	})

	var values []models.Resource

	// 1) Confirm org => retrieve numeric org ID
	orgID, err := fetchOrganizationID(sdk, organizationName)
	if err != nil {
		return nil, fmt.Errorf("fetching org ID: %w", err)
	}

	// 2) List all teams
	rawTeams, err := listAllTeams(sdk, organizationName, githubClient.Token)
	if err != nil {
		return nil, fmt.Errorf("listing teams: %w", err)
	}

	// 3) Build final slice, fetch team-sync for each
	var final []model.TeamDescription
	for _, t := range rawTeams {
		td := model.TeamDescription{
			Name:                t.Name,
			ID:                  t.ID,
			NodeID:              t.NodeID,
			Slug:                t.Slug,
			Description:         t.Description,
			Privacy:             t.Privacy,
			NotificationSetting: t.NotificationSetting,
			URL:                 t.URL,
			HTMLURL:             t.HTMLURL,
			Permission:          t.Permission,
			MembersCount:        t.MembersCount,
			ReposCount:          t.ReposCount,
			OrganizationID:      strconv.Itoa(orgID),
			Organization:        organizationName,
			ParentTeamID:        nil, // default null
			TeamSync:            nil, // default null
		}

		// If there's a parent, fill it
		if t.Parent != nil {
			td.ParentTeamID = &t.Parent.ID
		}

		// Attempt to get team sync
		syncInfo, syncErr := getTeamSyncInfo(sdk, organizationName, t.Slug, githubClient.Token)
		if syncErr != nil {
			// Could be 403 "not externally managed," 404, 429, etc.
			log.Printf("Warning: Could not fetch team-sync for team=%q => %v", t.Slug, syncErr)
		} else {
			td.TeamSync = syncInfo
		}

		final = append(final, td)
	}

	for _, team := range final {
		value := models.Resource{
			ID:          strconv.Itoa(team.ID),
			Name:        team.Name,
			Description: team,
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

// apiTeam => shape from GET /orgs/{org}/teams, without created_at / updated_at.
type apiTeam struct {
	Name                string  `json:"name"`
	ID                  int     `json:"id"`
	NodeID              string  `json:"node_id"`
	Slug                string  `json:"slug"`
	Description         *string `json:"description"`
	Privacy             string  `json:"privacy"`
	NotificationSetting *string `json:"notification_setting"`
	URL                 string  `json:"url"`
	HTMLURL             string  `json:"html_url"`
	Permission          string  `json:"permission"`
	MembersCount        int     `json:"members_count"`
	ReposCount          int     `json:"repos_count"`

	// Parent => only used to extract its ID
	Parent *struct {
		ID int `json:"id"`
	} `json:"parent"`
}

// ----------------------------------------------------
// 1) List All Teams in Org (Paginated)
// ----------------------------------------------------

func listAllTeams(sdk *resilientbridge.ResilientBridge, org, token string) ([]apiTeam, error) {
	var all []apiTeam
	page := 1
	perPage := 100

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
			return nil, fmt.Errorf("page %d => %w", page, err)
		}
		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("HTTP %d page %d => %s", resp.StatusCode, page, string(resp.Data))
		}

		var batch []apiTeam
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

// ----------------------------------------------------
// 2) Get Team Sync => GET /orgs/{org}/teams/{slug}/team-sync/group-mappings
// ----------------------------------------------------

func getTeamSyncInfo(sdk *resilientbridge.ResilientBridge, org, slug, token string) (*model.TeamIdpGroups, error) {
	endpoint := fmt.Sprintf("/orgs/%s/teams/%s/team-sync/group-mappings", org, slug)
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
		return nil, fmt.Errorf("requesting team-sync: %w", err)
	}

	if resp.StatusCode == http.StatusOK {
		var payload model.TeamIdpGroups
		if err := json.Unmarshal(resp.Data, &payload); err != nil {
			return nil, fmt.Errorf("unmarshal team-sync groups: %w", err)
		}
		return &payload, nil
	}

	// If not 200 => return error, caller sets TeamSync=null
	return nil, fmt.Errorf("HTTP %d => %s", resp.StatusCode, string(resp.Data))
}

// ----------------------------------------------------
// 4) ResilientBridge => No Auto-Retries
// ----------------------------------------------------

func newBridgeNoRetries() *resilientbridge.ResilientBridge {
	sdk := resilientbridge.NewResilientBridge()
	sdk.RegisterProvider("github", adapters.NewGitHubAdapter(""), &resilientbridge.ProviderConfig{
		UseProviderLimits: false,
		MaxRetries:        0,
		BaseBackoff:       0,
	})
	return sdk
}
