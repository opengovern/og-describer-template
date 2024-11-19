package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubTrafficViewWeekly() *plugin.Table {
	return &plugin.Table{
		Name:        "github_traffic_view_weekly",
		Description: "Weekly traffic view over the last 14 days for the given repository.",
		List: &plugin.ListConfig{
			KeyColumns:        plugin.SingleColumn("repository_full_name"),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           opengovernance.ListTrafficViewWeekly,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "repository_full_name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepositoryFullName"),
				Description: "Full name of the repository that contains the branch."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Timestamp").Transform(convertTimestamp),
				Description: "Date for the view data."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "View count for the day.",
				Transform: transform.FromField("Description.Count")},
			{Name: "uniques", Type: proto.ColumnType_INT, Description: "Unique viewer count for the day.",
				Transform: transform.FromField("Description.Uniques")},
		}),
	}
}
