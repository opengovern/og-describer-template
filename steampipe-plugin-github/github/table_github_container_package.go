package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGitHubContainerPackage() *plugin.Table {
	return &plugin.Table{
		Name: "github_container_package",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPackage,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    opengovernance.GetPackage,
		},
		Columns: []*plugin.Column{
			// Basic details columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the package."},
			{Name: "digest", Type: proto.ColumnType_STRING, Description: "Digest of the package."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL where the package can be accessed."},
			{Name: "package_uri", Type: proto.ColumnType_STRING, Description: "URI of the package."},
			{Name: "package_html_url", Type: proto.ColumnType_STRING, Description: "HTML URL of the package."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the package was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the package was last updated."},
			{Name: "html_url", Type: proto.ColumnType_STRING, Description: "HTML URL for the package."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the package."},
			{Name: "media_type", Type: proto.ColumnType_STRING, Description: "Media type of the package."},
			{Name: "total_size", Type: proto.ColumnType_INT, Description: "Total size of the package."},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Metadata of the package."},
			{Name: "manifest", Type: proto.ColumnType_JSON, Description: "Manifest of the package."},
		},
	}
}
