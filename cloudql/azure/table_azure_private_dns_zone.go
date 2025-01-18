package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzurePrivateDNSZone(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_private_dns_zone",
		Description: "Azure Private DNS Zone",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetPrivateDNSZones,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPrivateDNSZones,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the Private DNS zone.",
				Transform:   transform.FromField("Description.PrivateZone.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a Private DNS zone uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateZone.ID")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateZone.Etag")},
			{
				Name:        "type",
				Description: "The resource type of the Private DNS zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateZone.Type")},
			{
				Name:        "max_number_of_record_sets",
				Description: "The maximum number of record sets that can be created in this Private DNS zone.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PrivateZone.Properties.MaxNumberOfRecordSets")},
			{
				Name:        "number_of_record_sets",
				Description: "The current number of record sets in this Private DNS zone.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PrivateZone.Properties.NumberOfRecordSets")},
			{
				Name:        "max_number_of_virtual_network_links",
				Description: "The maximum number of virtual networks that can be linked to this Private DNS zone.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PrivateZone.Properties.MaxNumberOfVirtualNetworkLinks")},
			{
				Name:        "number_of_virtual_network_links",
				Description: "The current number of virtual networks that are linked to this Private DNS zone.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PrivateZone.Properties.NumberOfVirtualNetworkLinks")},
			{
				Name:        "max_number_of_virtual_network_links_with_registration",
				Description: "The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PrivateZone.Properties.MaxNumberOfVirtualNetworkLinksWithRegistration")},
			{
				Name:        "number_of_virtual_network_links_with_registration",
				Description: "The current number of virtual networks that are linked to this Private DNS zone with registration enabled.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PrivateZone.Properties.NumberOfVirtualNetworkLinksWithRegistration")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the resource. Possible values include: `Creating`, `Updating`, `Deleting`, `Succeeded`, `Failed`, `Canceled`.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateZone.Properties.ProvisioningState")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateZone.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateZone.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateZone.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateZone.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}
