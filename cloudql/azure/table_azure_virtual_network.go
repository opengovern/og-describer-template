package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureVirtualNetwork(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_virtual_network",
		Description: "Azure Virtual Network",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetVirtualNetwork,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListVirtualNetwork,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the virtual network",
				Transform:   transform.FromField("Description.VirtualNetwork.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a virtual network uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualNetwork.ID")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualNetwork.Etag")},
			{
				Name:        "type",
				Description: "Type of the resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualNetwork.Type")},
			{
				Name:        "enable_ddos_protection",
				Description: "Indicates if DDoS protection is enabled for all the protected resources in the virtual network",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualNetwork.Properties.EnableDdosProtection")},
			{
				Name:        "enable_vm_protection",
				Description: "Indicates if VM protection is enabled for all the subnets in the virtual network",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VirtualNetwork.Properties.EnableVMProtection")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the virtual network resource",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetwork.Properties.ProvisioningState"),
			},
			{
				Name:        "resource_guid",
				Description: "The resourceGuid property of the Virtual Network resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualNetwork.Properties.ResourceGUID")},
			{
				Name:        "address_prefixes",
				Description: "A list of address blocks reserved for this virtual network in CIDR notation",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualNetwork.Properties.AddressSpace.AddressPrefixes")},
			{
				Name:        "network_peerings",
				Description: "A list of peerings in a Virtual Network",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualNetwork.Properties.VirtualNetworkPeerings")},
			{
				Name:        "subnets",
				Description: "A list of subnets in a Virtual Network",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.VirtualNetwork.Properties.Subnets")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualNetwork.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualNetwork.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.VirtualNetwork.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetwork.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// FETCH FUNCTIONS ////

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
