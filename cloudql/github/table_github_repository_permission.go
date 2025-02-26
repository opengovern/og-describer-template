package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubRepositoryPermission() *plugin.Table {
	return &plugin.Table{
		Name:        "github_repository_permission",
		Description: "Get the software bill of materials (SBOM) for a repository.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRepositoryPermission,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "principal_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrincipalName"),
				Description: "The full name of the repository (login/repo-name).",
			},
			{
				Name:        "principal_id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PrincipalId"),
				Description: "The full name of the repository (login/repo-name).",
			},
			{
				Name:        "principal_type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrincipalType"),
				Description: "The full name of the repository (login/repo-name).",
			},
			{
				Name:        "repository_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepositoryName"),
				Description: "The full name of the repository (login/repo-name).",
			},
			{
				Name:        "repository_full_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepositoryFullName"),
				Description: "The full name of the repository (login/repo-name).",
			},
			{
				Name:        "repository_id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.RepositoryId"),
				Description: "The full name of the repository (login/repo-name).",
			},
			{
				Name:        "permissions",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Permissions"),
				Description: "The full name of the repository (login/repo-name).",
			},
			{
				Name:        "role_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleName"),
				Description: "The full name of the repository (login/repo-name).",
			},
		}),
	}
}
