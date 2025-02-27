package describers

import (
	"context"
	"fmt"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"log"
	"strconv"

	"github.com/opengovern/og-describer-github/discovery/pkg/models"
	model "github.com/opengovern/og-describer-github/discovery/provider"
)

func GetTeamList(ctx context.Context, githubClient model.GitHubClient, organizationName string, stream *models.StreamSender) ([]models.Resource, error) {
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
