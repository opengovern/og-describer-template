package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGitHubRubyGemsPackage() *plugin.Table {
	return &plugin.Table{
		Name: "github_ruby_gems_package",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPackage,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetPackage,
		},
		Columns: []*plugin.Column{
			// Basic details columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier for the package."},
			{Name: "registryId", Type: proto.ColumnType_STRING, Description: "Registry ID associated with the package."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the package."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL where the package can be accessed."},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the package was created."},
			{Name: "updatedAt", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the package was last updated."},
		},
	}
}
