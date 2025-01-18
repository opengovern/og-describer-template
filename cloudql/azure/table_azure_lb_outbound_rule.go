package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureLoadBalancerOutboundRule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_lb_outbound_rule",
		Description: "Azure Load Balancer Outbound Rule",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"load_balancer_name", "name", "resource_group"}),
			Hydrate:    opengovernance.GetLoadBalancerOutboundRule,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLoadBalancerOutboundRule,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource that is unique within the set of outbound rules used by the load balancer. This name can be used to access the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Name")},
			{
				Name:        "id",
				Description: "The resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.ID"),
			},
			{
				Name:        "load_balancer_name",
				Description: "The friendly name that identifies the load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancerName")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the outbound rule resource. Possible values include: 'ProvisioningStateSucceeded', 'ProvisioningStateUpdating', 'ProvisioningStateDeleting', 'ProvisioningStateFailed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Type")},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Etag")},
			{
				Name:        "allocated_outbound_ports",
				Description: "The number of outbound ports to be used for NAT.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Rule.Properties.AllocatedOutboundPorts")},
			{
				Name:        "enable_tcp_reset",
				Description: "Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Rule.Properties.EnableTCPReset"), Default: false,
			},
			{
				Name:        "idle_timeout_in_minutes",
				Description: "The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Rule.Properties.IdleTimeoutInMinutes")},
			{
				Name:        "protocol",
				Description: "The protocol for the outbound rule in load balancer. Possible values include: 'LoadBalancerOutboundRuleProtocolTCP', 'LoadBalancerOutboundRuleProtocolUDP', 'LoadBalancerOutboundRuleProtocolAll'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Properties.Protocol")},
			{
				Name:        "backend_address_pools",
				Description: "A reference to a pool of DIPs. Outbound traffic is randomly load balanced across IPs in the backend IPs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.Properties.BackendAddressPool")},
			{
				Name:        "frontend_ip_configurations",
				Description: "The Frontend IP addresses of the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.Properties.FrontendIPConfigurations")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Rule.ID").Transform(idToAkas),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
