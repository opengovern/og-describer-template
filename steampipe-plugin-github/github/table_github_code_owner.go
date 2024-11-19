package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func gitHubCodeOwnerColumns() []*plugin.Column {
	return []*plugin.Column{
		// Top columns
		{Name: "repository_full_name", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.RepositoryFullName"),
			Description: "The full name of the repository, including the owner and repo name."},
		// Other columns
		{Name: "line", Type: proto.ColumnType_INT,
			Description: "The rule's line number in the CODEOWNERS file.",
			Transform:   transform.FromField("Description.LineNumber"),
		},
		{Name: "pattern", Type: proto.ColumnType_STRING,
			Transform:   transform.FromField("Description.Pattern"),
			Description: "The pattern used to identify what code a team, or an individual is responsible for"},
		{Name: "users",
			Transform: transform.FromField("Description.Users"),
			Type:      proto.ColumnType_JSON, Description: "Users responsible for code in the repo"},
		{Name: "teams",
			Transform: transform.FromField("Description.Teams"),
			Type:      proto.ColumnType_JSON, Description: "Teams responsible for code in the repo"},
		{Name: "pre_comments",
			Transform: transform.FromField("Description.PreComments"),
			Type:      proto.ColumnType_JSON, Description: "Specifies the comments added above a key."},
		{Name: "line_comment",
			Transform: transform.FromField("Description.LineComment"),
			Type:      proto.ColumnType_STRING, Description: "Specifies the comment following the node and before empty lines."},
	}
}

func tableGitHubCodeOwner() *plugin.Table {
	return &plugin.Table{
		Name:        "github_code_owner",
		Description: "Individuals or teams that are responsible for code in a repository.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCodeOwner,
		},
		Columns: commonColumns(gitHubCodeOwnerColumns()),
	}
}
