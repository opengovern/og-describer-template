package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"

	"github.com/opengovern/og-describer-github/steampipe-plugin-github/github/models"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubStargazer() *plugin.Table {
	return &plugin.Table{
		Name:        "github_stargazer",
		Description: "Stargazers are users who have starred the repository.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStargazer,
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "repository_full_name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepoFullName"),
				Description: "Full name of the repository that contains the stargazer."},
			{Name: "starred_at", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.StarredAt").NullIfZero().Transform(convertTimestamp),
				Description: "Time when the stargazer was created."},
			{Name: "user_login", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UserLogin"),
				Description: "The login name of the user who starred the repository."},
			{Name: "user_detail", Type: proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.UserDetail"),
				Description: "Details of the user who starred the repository."},
		}),
	}
}

type Stargazer struct {
	StarredAt models.NullableTime `graphql:"starredAt @include(if:$includeStargazerStarredAt)" json:"starred_at"`
	Node      models.BasicUser    `graphql:"node @include(if:$includeStargazerNode)" json:"ndoe"`
}
