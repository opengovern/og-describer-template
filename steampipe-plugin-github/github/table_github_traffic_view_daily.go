package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubTrafficViewDaily() *plugin.Table {
	return &plugin.Table{
		Name:        "github_traffic_view_daily",
		Description: "Daily traffic view over the last 14 days for the given repository.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListTrafficViewDaily,
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "repository_full_name", Type: proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepositoryFullName"),
				Description: "Full name of the repository that contains the branch."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Timestamp").NullIfZero().Transform(convertTimestamp),
				Description: "Date for the view data."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "View count for the day.",
				Transform: transform.FromField("Description.Count")},
			{Name: "uniques", Type: proto.ColumnType_INT, Description: "Unique viewer count for the day.",
				Transform: transform.FromField("Description.Uniques")},
		}),
	}
}
