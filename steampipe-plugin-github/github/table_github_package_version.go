package github

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableGitHubPackageVersion() *plugin.Table {
	return &plugin.Table{
		Name:        "github_package_version",
		Description: "Details of package versions, including version, size, digest, and download count.",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			// Basic details columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique identifier for the package version."},
			{Name: "packageType", Type: proto.ColumnType_STRING, Description: "Type of the package (e.g., binary, source, etc.)."},
			{Name: "packageId", Type: proto.ColumnType_INT, Description: "ID of the associated package."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "Version of the package."},
			{Name: "digest", Type: proto.ColumnType_STRING, Description: "Digest of the package version, if available."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Size of the package version in bytes."},
			{Name: "createdAt", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the package version was created."},
			{Name: "updatedAt", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp when the package version was last updated."},
			{Name: "downloadCount", Type: proto.ColumnType_INT, Description: "Number of times the package version has been downloaded."},
		},
	}
}
