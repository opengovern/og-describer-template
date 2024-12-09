package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGitHubPackageVersion() *plugin.Table {
	return &plugin.Table{
		Name:        "github_package_version",
		Description: "Details of package versions, including version, size, digest, and download count.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPackageVersion,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id"}),
			Hydrate:    opengovernance.GetPackageVersion,
		},
		Columns: []*plugin.Column{
			// Basic details columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the package version."},
			{Name: "name", Type: proto.ColumnType_INT, Description: "name and version of the package"},
			{Name: "versionUri", Type: proto.ColumnType_STRING, Description: "version uri of the package."},
			{Name: "digest", Type: proto.ColumnType_STRING, Description: "Digest of the package version"},
			{Name: "packageName", Type: proto.ColumnType_INT, Description: "name of the package"},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the package version was created."},
			{Name: "updatedAt", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the package version was last updated."},
		},
	}
}
