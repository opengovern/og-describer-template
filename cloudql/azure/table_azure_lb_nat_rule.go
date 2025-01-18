package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureLoadBalancerNatRule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_lb_nat_rule",
		Description: "Azure Load Balancer Nat Rule",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"load_balancer_name", "name", "resource_group"}),
			Hydrate:    opengovernance.GetLoadBalancerNatRule,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLoadBalancerNatRule,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Name"),
			},
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
				Transform:   transform.FromField("Description.LoadBalancerName"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the inbound NAT rule resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Properties.ProvisioningState"),
			},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Type"),
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Etag"),
			},
			{
				Name:        "backend_port",
				Description: "The port used for the internal endpoint. Acceptable values range from 1 to 65535.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Rule.Properties.BackendPort"),
			},
			{
				Name:        "enable_floating_ip",
				Description: "Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Rule.Properties.EnableFloatingIP"),
				Default:     false,
			},
			{
				Name:        "frontend_port",
				Description: "The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Rule.Properties.FrontendPort"),
			},
			{
				Name:        "enable_tcp_reset",
				Description: "Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Rule.Properties.EnableTCPReset"),
				Default:     false,
			},
			{
				Name:        "idle_timeout_in_minutes",
				Description: "The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Rule.Properties.IdleTimeoutInMinutes"),
			},
			{
				Name:        "protocol",
				Description: "The reference to the transport protocol used by the load balancing rule. Possible values include: 'TransportProtocolUDP', 'TransportProtocolTCP', 'TransportProtocolAll'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Properties.Protocol"),
			},
			{
				Name:        "backend_ip_configuration",
				Description: "A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.Properties.BackendIPConfiguration"),
			},
			{
				Name:        "frontend_ip_configuration",
				Description: "A reference to frontend IP addresses.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.Properties.FrontendIPConfiguration"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Name"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}
