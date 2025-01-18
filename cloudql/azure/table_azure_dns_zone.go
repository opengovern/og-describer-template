package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureDNSZone(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_dns_zone",
		Description: "Azure DNS Zone",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetDNSZones,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDNSZones,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the DNS zone.",
				Transform:   transform.FromField("Description.Zone.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a DNS zone uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Zone.ID")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Zone.Etag")},
			{
				Name:        "type",
				Description: "The resource type of the DNS zone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Zone.Type")},
			{
				Name:        "max_number_of_record_sets",
				Description: "The maximum number of record sets that can be created in this DNS zone.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Zone.Properties.MaxNumberOfRecordSets")},
			{
				Name:        "max_number_of_records_per_record_set",
				Description: "The maximum number of records per record set that can be created in this DNS zone.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Zone.Properties.MaxNumberOfRecordsPerRecordSet")},
			{
				Name:        "number_of_record_sets",
				Description: "The current number of record sets in this DNS zone.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Zone.Properties.NumberOfRecordSets"),
			},
			{
				Name:        "name_servers",
				Description: "The name servers for this DNS zone.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Zone.Properties.NameServers")},
			{
				Name:        "zone_type",
				Description: "The type of this DNS zone (always `Public`, see `azure_private_dns_zone` table for private DNS zones).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Zone.Properties.ZoneType")},
			{
				Name:        "registration_virtual_networks",
				Description: "A list of references to virtual networks that register hostnames in this DNS zone.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Zone.Properties.RegistrationVirtualNetworks")},
			{
				Name:        "resolution_virtual_networks",
				Description: "A list of references to virtual networks that resolve records in this DNS zone.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Zone.Properties.ResolutionVirtualNetworks")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Zone.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Zone.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Zone.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Zone.Location").Transform(toLower),
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
