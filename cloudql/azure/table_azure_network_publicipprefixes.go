package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureNetworkPublicIPPrefixes(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_network_publicipprefixes",
		Description: "Azure Network PublicIPPrefixes",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetPublicIPPrefix,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPublicIPPrefix,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the publicipprefixes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPPrefixes.ID")},
			{
				Name:        "name",
				Description: "The name of the publicipprefixes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPPrefix.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPPrefix.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.PublicIPPrefix.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.PublicIPPrefix.ID").Transform(idToAkas),
			},
		}),
	}
}
