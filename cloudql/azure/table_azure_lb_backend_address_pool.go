package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureLoadBalancerBackendAddressPool(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_lb_backend_address_pool",
		Description: "Azure Load Balancer Backend Address Pool",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"load_balancer_name", "name", "resource_group"}),
			Hydrate:    opengovernance.GetLoadBalancerBackendAddressPool,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLoadBalancerBackendAddressPool,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.Name")},
			{
				Name:        "id",
				Description: "The resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.ID")},
			{
				Name:        "load_balancer_name",
				Description: "The friendly name that identifies the load balancer.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LoadBalancer.Name")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the backend address pool resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.Type")},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.Etag")},
			{
				Name:        "outbound_rule_id",
				Description: "A reference to an outbound rule that uses this backend address pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.Properties.OutboundRule.ID"),
			},
			{
				Name:        "backend_ip_configurations",
				Description: "An array of references to IP addresses defined in network interfaces.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pool.Properties.BackendIPConfigurations")},
			{
				Name:        "gateway_load_balancer_tunnel_interface",
				Description: "An array of gateway load balancer tunnel interfaces.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pool.Properties.TunnelInterfaces")},
			{
				Name:        "load_balancer_backend_addresses",
				Description: "An array of backend addresses.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Pool.Properties.LoadBalancerBackendAddresses")},
			{
				Name:        "load_balancing_rules",
				Description: "An array of references to load balancing rules that use this backend address pool.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.LoadBalancingRules")},
			{
				Name:        "outbound_rules",
				Description: "An array of references to outbound rules that use this backend address pool.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.OutboundRules")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pool.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Pool.ID").Transform(idToAkas),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
