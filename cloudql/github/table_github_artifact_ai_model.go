package github

import (
	opengovernance "github.com/opengovern/og-describer-github/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubArtifactAIModel() *plugin.Table {
	return &plugin.Table{
		Name:        "github_artifact_ai_model",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListArtifactAIModel,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"repository_full_name", "name"}),
			Hydrate:    opengovernance.GetArtifactAIModel,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Name"),
				Description: "name",
			},
			{
				Name:        "repository_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepositoryID"),
				Description: "repository id",
			},
			{
				Name:        "repository_full_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepositoryFullName"),
				Description: "Full name of the repository that contains the artifact.",
			},
			{
				Name:        "repository_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepositoryName"),
				Description: "repository name",
			},
			{
				Name:        "extensions",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Extensions"),
				Description: "extention",
			},
		}),
	}
}
