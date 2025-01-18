package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureCdnProfiles(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_cdn_profiles",
		Description: "Azure Cdn Profiles",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetCDNProfile,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCDNProfile,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the profiles.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Profiles.ID")},
			{
				Name:        "name",
				Description: "The name of the profiles.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Profile.Name")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Profile.Type"),
			},
			{
				Name:        "location",
				Description: "The location of the CDN front door profile.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Profile.Location").Transform(toLower),
			},
			{
				Name:        "sku_name",
				Description: "Name of the pricing tier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Profile.SKU.Name"),
			},
			{
				Name:        "kind",
				Description: "Kind of the profile. Used by portal to differentiate traditional CDN profile and new AFD profile.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Profile.Kind"),
			},
			{
				Name:        "resource_state",
				Description: "Resource status of the CDN front door profile.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Profile.Properties.ResourceState").Transform(transform.ToString),
			},
			{
				Name:        "provisioning_state",
				Description: "Provisioning status of the CDN front door profile.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Profile.Properties.ProvisioningState"),
			},
			{
				Name:        "front_door_id",
				Description: "The ID of the front door.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Profile.Properties.FrontDoorID"),
			},
			{
				Name:        "origin_response_timeout_seconds",
				Description: "Send and receive timeout on forwarding request to the origin. When timeout is reached, the request fails and returns.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Profile.Properties.OriginResponseTimeoutSeconds"),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Profile.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.Profile.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.Profile.ID").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: "The Azure region where the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Profile.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: "The resource group in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}
