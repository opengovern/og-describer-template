package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION ////

func tableAzureComputeAvailabilitySet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_availability_set",
		Description: "Azure Compute Availability Set",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetComputeAvailabilitySet,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeAvailabilitySet,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the availability set",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilitySet.Name")},
			{
				Name:        "id",
				Description: "The unique id identifying the resource in subscription",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilitySet.ID")},
			{
				Name:        "type",
				Description: "The type of the resource in Azure",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AvailabilitySet.Type")},
			{
				Name:        "platform_fault_domain_count",
				Description: "Contains the fault domain count",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.AvailabilitySet.Properties.PlatformFaultDomainCount")},
			{
				Name:        "platform_update_domain_count",
				Description: "Contains the update domain count",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.AvailabilitySet.Properties.PlatformUpdateDomainCount")},
			{
				Name:        "proximity_placement_group_id",
				Description: "Specifies information about the proximity placement group that the availability set should be assigned to",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AvailabilitySet.Properties.ProximityPlacementGroup.ID")},
			{
				Name:        "sku_capacity",
				Description: "Specifies the number of virtual machines in the scale set",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.AvailabilitySet.SKU.Capacity")},
			{
				Name:        "sku_name",
				Description: "The availability sets sku name",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AvailabilitySet.SKU.Name")},
			{
				Name:        "sku_tier",
				Description: "Specifies the tier of virtual machines in a scale set",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AvailabilitySet.SKU.Tier")},
			{
				Name:        "status",
				Description: "The resource status information",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.AvailabilitySet.Properties.Statuses")},
			{
				Name:        "virtual_machines",
				Description: "A list of references to all virtual machines in the availability set",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.AvailabilitySet.Properties.VirtualMachines")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AvailabilitySet.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AvailabilitySet.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.AvailabilitySet.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AvailabilitySet.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS ////

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data
