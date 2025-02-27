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

func ListOrganizationApps(ctx context.Context,
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

	endpoint := fmt.Sprintf("/orgs/%s/installations", organizationName)
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

	var appsResponse OrganizationAppsResponse
	if err := json.Unmarshal(resp.Data, &appsResponse); err != nil {
		return nil, fmt.Errorf("error decoding repos list: %w", err)
	}

	orgID, err := fetchOrganizationID(sdk, organizationName)
	if err != nil {
		return nil, fmt.Errorf("fetching org ID: %w", err)
	}

	for _, app := range appsResponse.Installations {
		account := model.Account{
			Login:             app.Account.Login,
			ID:                app.Account.ID,
			NodeID:            app.Account.NodeID,
			AvatarURL:         app.Account.AvatarURL,
			GravatarID:        app.Account.GravatarID,
			URL:               app.Account.URL,
			HTMLURL:           app.Account.HTMLURL,
			FollowersURL:      app.Account.FollowersURL,
			FollowingURL:      app.Account.FollowingURL,
			GistsURL:          app.Account.GistsURL,
			StarredURL:        app.Account.StarredURL,
			SubscriptionsURL:  app.Account.SubscriptionsURL,
			OrganizationsURL:  app.Account.OrganizationsURL,
			ReposURL:          app.Account.ReposURL,
			EventsURL:         app.Account.EventsURL,
			ReceivedEventsURL: app.Account.ReceivedEventsURL,
			Type:              app.Account.Type,
			UserViewType:      app.Account.UserViewType,
			SiteAdmin:         app.Account.SiteAdmin,
		}
		var suspendedBy *model.Account
		if app.SuspendedBy != nil {
			suspendedBy = &model.Account{
				Login:             app.SuspendedBy.Login,
				ID:                app.SuspendedBy.ID,
				NodeID:            app.SuspendedBy.NodeID,
				AvatarURL:         app.SuspendedBy.AvatarURL,
				GravatarID:        app.SuspendedBy.GravatarID,
				URL:               app.SuspendedBy.URL,
				HTMLURL:           app.SuspendedBy.HTMLURL,
				FollowersURL:      app.SuspendedBy.FollowersURL,
				FollowingURL:      app.SuspendedBy.FollowingURL,
				GistsURL:          app.SuspendedBy.GistsURL,
				StarredURL:        app.SuspendedBy.StarredURL,
				SubscriptionsURL:  app.SuspendedBy.SubscriptionsURL,
				OrganizationsURL:  app.SuspendedBy.OrganizationsURL,
				ReposURL:          app.SuspendedBy.ReposURL,
				EventsURL:         app.SuspendedBy.EventsURL,
				ReceivedEventsURL: app.SuspendedBy.ReceivedEventsURL,
				Type:              app.SuspendedBy.Type,
				UserViewType:      app.SuspendedBy.UserViewType,
				SiteAdmin:         app.SuspendedBy.SiteAdmin,
			}
		}
		value := models.Resource{
			ID:   strconv.Itoa(int(app.ID)),
			Name: app.AppSlug,
			Description: model.OrganizationAppDescription{
				Organization:           organizationName,
				OrganizationID:         orgID,
				ID:                     app.ID,
				ClientID:               app.ClientID,
				Account:                account,
				RepositorySelection:    app.RepositorySelection,
				AccessTokensURL:        app.AccessTokensURL,
				RepositoriesURL:        app.RepositoriesURL,
				HTMLURL:                app.HTMLURL,
				AppID:                  app.AppID,
				AppSlug:                app.AppSlug,
				TargetID:               app.TargetID,
				TargetType:             app.TargetType,
				Permissions:            app.Permissions,
				Events:                 app.Events,
				CreatedAt:              app.CreatedAt,
				UpdatedAt:              app.UpdatedAt,
				SingleFileName:         app.SingleFileName,
				HasMultipleSingleFiles: app.HasMultipleSingleFiles,
				SingleFilePaths:        app.SingleFilePaths,
				SuspendedBy:            suspendedBy,
				SuspendedAt:            app.SuspendedAt,
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

type OrganizationAppsResponse struct {
	Installations []OrganizationApp `json:"installations"`
}

type OrganizationApp struct {
	ID                     int64             `json:"id"`
	ClientID               string            `json:"client_id"`
	Account                AccountJSON       `json:"account"`
	RepositorySelection    string            `json:"repository_selection"`
	AccessTokensURL        string            `json:"access_tokens_url"`
	RepositoriesURL        string            `json:"repositories_url"`
	HTMLURL                string            `json:"html_url"`
	AppID                  int               `json:"app_id"`
	AppSlug                string            `json:"app_slug"`
	TargetID               int64             `json:"target_id"`
	TargetType             string            `json:"target_type"`
	Permissions            map[string]string `json:"permissions"`
	Events                 []string          `json:"events"`
	CreatedAt              time.Time         `json:"created_at"`
	UpdatedAt              time.Time         `json:"updated_at"`
	SingleFileName         *string           `json:"single_file_name"`
	HasMultipleSingleFiles bool              `json:"has_multiple_single_files"`
	SingleFilePaths        []string          `json:"single_file_paths"`
	SuspendedBy            *AccountJSON      `json:"suspended_by"`
	SuspendedAt            *time.Time        `json:"suspended_at"`
}

type AccountJSON struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	UserViewType      string `json:"user_view_type"`
	SiteAdmin         bool   `json:"site_admin"`
}
