package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGitHubArtifactDockerFile() *plugin.Table {
	return &plugin.Table{
		Name: "github_artifact_dockerfile",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListArtifactDockerFile,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"sha"}),
			Hydrate:    opengovernance.GetArtifactDockerFile,
		},
		Columns: []*plugin.Column{
			// Basic details columns
			{Name: "sha", Type: proto.ColumnType_STRING, Description: "SHA hash of the Dockerfile."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the Dockerfile."},
			{Name: "path", Type: proto.ColumnType_STRING, Description: "Path to the Dockerfile in the repository."},
			{Name: "last_updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the Dockerfile was last updated."},
			{Name: "git_url", Type: proto.ColumnType_STRING, Description: "Git URL where the Dockerfile can be accessed."},
			{Name: "html_url", Type: proto.ColumnType_STRING, Description: "HTML URL where the Dockerfile can be accessed."},
			{Name: "uri", Type: proto.ColumnType_STRING, Description: "Unique URI for the Dockerfile."},
			{Name: "dockerfile_content", Type: proto.ColumnType_STRING, Description: "Content of the Dockerfile."},
			{Name: "dockerfile_content_base64", Type: proto.ColumnType_STRING, Description: "Base64-encoded content of the Dockerfile."},
			{Name: "repository", Type: proto.ColumnType_JSON, Description: "Repository metadata associated with the Dockerfile."},
		},
	}
}
