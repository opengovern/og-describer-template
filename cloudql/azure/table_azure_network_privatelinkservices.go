package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureNetworkPrivateLinkServices(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_network_privatelinkservices",
		Description: "Azure Network PrivateLinkServices",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetPrivateLinkService,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPrivateLinkService,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the privatelinkservices.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateLinkServices.ID")},
			{
				Name:        "name",
				Description: "The name of the privatelinkservices.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateLinkService.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateLinkService.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.PrivateLinkService.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.PrivateLinkService.ID").Transform(idToAkas),
			},
		}),
	}
}
