package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func sharedRepositoryColumns() []*plugin.Column {
	return commonColumns([]*plugin.Column{
		{
			Name:        "id",
			Type:        proto.ColumnType_INT,
			Description: "Unique identifier of the GitHub repository.",
		},
		{
			Name:        "node_id",
			Type:        proto.ColumnType_STRING,
			Description: "Node ID of the repository.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the repository.",
		},
		{
			Name:        "name_with_owner",
			Type:        proto.ColumnType_STRING,
			Description: "Full name of the repository including the owner.",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description of the repository.",
		},
		{
			Name:        "created_at",
			Type:        proto.ColumnType_STRING,
			Description: "Timestamp when the repository was created.",
		},
		{
			Name:        "updated_at",
			Type:        proto.ColumnType_STRING,
			Description: "Timestamp when the repository was last updated.",
		},
		{
			Name:        "pushed_at",
			Type:        proto.ColumnType_STRING,
			Description: "Timestamp when the repository was last pushed.",
		},
		{
			Name:        "is_active",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the repository is active.",
		},
		{
			Name:        "is_empty",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the repository is empty.",
		},
		{
			Name:        "is_fork",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the repository is a fork.",
		},
		{
			Name:        "is_security_policy_enabled",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the repository has a security policy enabled.",
		},
		{
			Name:        "owner",
			Type:        proto.ColumnType_JSON,
			Description: "Owner details of the repository.",
		},
		{
			Name:        "homepage_url",
			Type:        proto.ColumnType_STRING,
			Description: "Homepage URL of the repository.",
		},
		{
			Name:        "license_info",
			Type:        proto.ColumnType_JSON,
			Description: "License information of the repository.",
		},
		{
			Name:        "topics",
			Type:        proto.ColumnType_JSON,
			Description: "List of topics associated with the repository.",
		},
		{
			Name:        "visibility",
			Type:        proto.ColumnType_STRING,
			Description: "Visibility status of the repository.",
		},
		{
			Name:        "default_branch_ref",
			Type:        proto.ColumnType_JSON,
			Description: "Details of the default branch of the repository.",
		},
		{
			Name:        "permissions",
			Type:        proto.ColumnType_JSON,
			Description: "Permissions associated with the repository.",
		},
		{
			Name:        "organization",
			Type:        proto.ColumnType_JSON,
			Description: "Organization details of the repository.",
		},
		{
			Name:        "parent",
			Type:        proto.ColumnType_JSON,
			Description: "Parent repository details if the repository is forked.",
		},
		{
			Name:        "source",
			Type:        proto.ColumnType_JSON,
			Description: "Source repository details if the repository is forked.",
		},
		{
			Name:        "languages",
			Type:        proto.ColumnType_JSON,
			Description: "Languages used in the repository along with their usage statistics.",
		},
		{
			Name:        "repo_settings",
			Type:        proto.ColumnType_JSON,
			Description: "Settings of the repository.",
		},
		{
			Name:        "security_settings",
			Type:        proto.ColumnType_JSON,
			Description: "Security settings of the repository.",
		},
		{
			Name:        "repo_urls",
			Type:        proto.ColumnType_JSON,
			Description: "Repository URLs for different purposes (e.g., clone URLs).",
		},
		{
			Name:        "metrics",
			Type:        proto.ColumnType_JSON,
			Description: "Metrics and statistics of the repository.",
		},
	})
}

func tableGitHubRepository() *plugin.Table {
	return &plugin.Table{
		Name:        "github_repository",
		Description: "GitHub Repositories contain all of your project's files and each file's revision history.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRepository,
		},
		Columns: commonColumns(sharedRepositoryColumns()),
	}
}
