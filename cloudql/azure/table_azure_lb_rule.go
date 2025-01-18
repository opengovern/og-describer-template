package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureLoadBalancerRule(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_lb_rule",
		Description: "Azure Load Balancer Rule",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"load_balancer_name", "name", "resource_group"}),
			Hydrate:    opengovernance.GetLoadBalancerRule,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLoadBalancerRule,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource that is unique within the set of load balancing rules used by the load balancer. This name can be used to access the resource.",
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

				Transform: transform.FromField("Description.LoadBalancerName")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the load balancing rule resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.",
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
				Name:        "backend_address_pool_id",
				Description: "A reference to a pool of DIPs. Inbound traffic is randomly load balanced across IPs in the backend IPs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Properties.BackendAddressPool.ID")},
			{
				Name:        "backend_port",
				Description: "The port used for internal connections on the endpoint. Acceptable values are between 0 and 65535. Note that value 0 enables 'Any Port'.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Rule.Properties.BackendPort")},
			{
				Name:        "disable_outbound_snat",
				Description: "Configures SNAT for the VMs in the backend pool to use the publicIP address specified in the frontend of the load balancing rule.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Rule.Properties.DisableOutboundSnat"), Default: false,
			},
			{
				Name:        "enable_floating_ip",
				Description: "Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Rule.Properties.EnableFloatingIP"), Default: false,
			},
			{
				Name:        "enable_tcp_reset",
				Description: "Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Rule.Properties.EnableTCPReset"), Default: false,
			},
			{
				Name:        "frontend_ip_configuration_id",
				Description: "A reference to frontend IP addresses.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Properties.FrontendIPConfiguration.ID")},
			{
				Name:        "frontend_port",
				Description: "The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values are between 0 and 65534. Note that value 0 enables 'Any Port'.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Rule.Properties.FrontendPort")},
			{
				Name:        "idle_timeout_in_minutes",
				Description: "The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Rule.Properties.IdleTimeoutInMinutes")},
			{
				Name:        "load_distribution",
				Description: "The load distribution policy for this rule. Possible values include: 'Default', 'SourceIP', 'SourceIPProtocol'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Properties.LoadDistribution")},
			{
				Name:        "probe_id",
				Description: "The reference to the load balancer probe used by the load balancing rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Properties.Probe.ID")},
			{
				Name:        "protocol",
				Description: "The reference to the transport protocol used by the load balancing rule. Possible values include: 'UDP', 'TCP', 'All'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Rule.Properties.Protocol")},
			{
				Name:        "backend_address_pools",
				Description: "An array of references to pool of DIPs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Rule.Properties.BackendAddressPools")},

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
