package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSubnet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_subnet",
		Description: "Azure Subnet",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "virtual_network_name", "resource_group"}),
			Hydrate:    opengovernance.GetSubnet,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "NotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSubnet,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the subnet.",
				Transform:   transform.FromField("Description.Subnet.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a subnet uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.ID"),
			},
			{
				Name:        "virtual_network_name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name of the virtual network in which the subnet is created.",
				Transform:   transform.FromField("Description.VirtualNetworkName")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.Etag")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Default:     "Microsoft.Network/subnets",
				Transform:   transform.FromField("Description.Subnet.Properties.RouteTable.Type")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the subnet resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Subnet.Properties.ProvisioningState"),
			},
			{
				Name:        "address_prefix",
				Description: "Contains the address prefix for the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.Properties.AddressPrefix")},
			{
				Name:        "nat_gateway_id",
				Description: "The ID of the Nat gateway associated with the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.Properties.NatGateway.ID")},
			{
				Name:        "network_security_group_id",
				Description: "Network security group associated with the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.Properties.NetworkSecurityGroup.ID")},
			{
				Name:        "private_endpoint_network_policies",
				Description: "Enable or Disable apply network policies on private end point in the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.Properties.PrivateEndpointNetworkPolicies")},
			{
				Name:        "private_link_service_network_policies",
				Description: "Enable or Disable apply network policies on private link service in the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.Properties.PrivateLinkServiceNetworkPolicies")},
			{
				Name:        "route_table_id",
				Description: "Route table associated with the subnet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.Properties.RouteTable.ID")},
			{
				Name:        "delegations",
				Description: "A list of references to the delegations on the subnet.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subnet.Properties.Delegations")},
			{
				Name:        "ip_configurations",
				Description: "IP Configuration details in a subnet.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subnet.Properties.IPConfigurations")},
			{
				Name:        "service_endpoints",
				Description: "A list of service endpoints.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subnet.Properties.ServiceEndpoints")},
			{
				Name:        "service_endpoint_policies",
				Description: "A list of service endpoint policies.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Subnet.Properties.ServiceEndpointPolicies")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subnet.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Subnet.ID").Transform(idToAkas),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.

					// Get the details of virtual network
					FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data
