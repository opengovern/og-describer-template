package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubTeamRepository() *plugin.Table {
	return &plugin.Table{
		Name:        "github_team_repository",
		Description: "GitHub Repositories that a given team is associated with. GitHub Repositories contain all of your project's files and each file's revision history.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListTeamRepository,
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "repository_full_name", Require: plugin.Required},
			},
			Hydrate: opengovernance.GetTeamRepository,
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "permission", Type: proto.ColumnType_STRING, Description: "The permission level the team has on the repository.",
				Transform: transform.FromField("Description.Permission")},

			{Name: "team_id", Type: proto.ColumnType_INT, Description: "",
				Transform: transform.FromField("Description.TeamID")},

			{Name: "repository_full_name", Type: proto.ColumnType_STRING, Description: "",
				Transform: transform.FromField("Description.RepositoryFullName")},

			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "",
				Transform: transform.FromField("Description.CreatedAt").NullIfZero().Transform(convertTimestamp)},

			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "",
				Transform: transform.FromField("Description.UpdatedAt").NullIfZero().Transform(convertTimestamp)},
		}),
	}
}
