package github

import (
	opengovernance "github.com/opengovern/og-describer-github/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableGitHubBlob() *plugin.Table {
	return &plugin.Table{
		Name:        "github_blob",
		Description: "Gets a blob from a repository.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBlob,
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "repository_full_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RepoFullName"),
				Description: "Full name of the repository that contains the blob.",
			},
			{
				Name:        "blob_sha",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SHA"),
				Description: "SHA1 of the blob."},
			// Other columns
			{
				Name:        "node_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NodeID"),
				Description: "The node ID of the blob.",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.URL"),
				Description: "URL of the blob.",
			},
			{
				Name:        "content",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Content"),
				Description: "The encoded content of the blob.",
			},
			{
				Name:        "encoding",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Encoding"),
				Description: "The encoding of the blob.",
			},
			{
				Name:        "size",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Size"),
				Description: "Size of the blob.",
			},
		},
	}
}
