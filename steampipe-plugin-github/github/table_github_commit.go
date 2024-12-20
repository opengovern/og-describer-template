package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGitHubCommit() *plugin.Table {
	return &plugin.Table{
		Name:        "github_commit",
		Description: "GitHub Commits bundle project files for download by users.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCommit,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.AllColumns([]string{"id"}),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			Hydrate:           opengovernance.GetCommit,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier (SHA) of the commit.",
			},
			{
				Name:        "message",
				Type:        proto.ColumnType_STRING,
				Description: "Commit message.",
			},
			{
				Name:        "author_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Author.Name"),
				Description: "Name of the author of the commit.",
			},
			{
				Name:        "author_email",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Author.Email"),
				Description: "Email address of the author.",
			},
			{
				Name:        "date",
				Type:        proto.ColumnType_STRING,
				Description: "Date of the commit.",
			},
			{
				Name:        "comment_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of comments on the commit.",
			},
			{
				Name:        "html_url",
				Type:        proto.ColumnType_STRING,
				Description: "URL of the commit on the repository.",
			},
			{
				Name:        "is_verified",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates if the commit is verified.",
			},
			{
				Name:        "changes",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Changes"),
				Description: "Details of the changes made in the commit.",
			},
			{
				Name:        "files",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Files"),
				Description: "List of files changed in the commit.",
			},
			{
				Name:        "pull_requests",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("PullRequests"),
				Description: "List of associated pull request IDs.",
			},
			{
				Name:        "target",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Target"),
				Description: "Target details of the commit (e.g., branch, organization, repository).",
			},
			{
				Name:        "additional_details",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AdditionalDetails"),
				Description: "Additional details about the commit.",
			},
		},
	}
}
