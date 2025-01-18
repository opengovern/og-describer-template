package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureLoadBalancerProbe(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_lb_probe",
		Description: "Azure Load Balancer Probe",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"load_balancer_name", "name", "resource_group"}),
			Hydrate:    opengovernance.GetLoadBalancerProbe,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLoadBalancerProbe,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource that is unique within the set of probes used by the load balancer. This name can be used to access the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Probe.Name")},
			{
				Name:        "id",
				Description: "The resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Probe.ID"),
			},
			{
				Name:        "load_balancer_name",
				Description: "The friendly name that identifies the load balancer.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LoadBalancerName")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the probe resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Probe.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Probe.Type")},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Probe.Etag")},
			{
				Name:        "interval_in_seconds",
				Description: "The interval, in seconds, for how frequently to probe the endpoint for health status. Typically, the interval is slightly less than half the allocated timeout period (in seconds) which allows two full probes before taking the instance out of rotation. The default value is 15, the minimum value is 5.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Probe.Properties.IntervalInSeconds")},
			{
				Name:        "number_of_probes",
				Description: "The number of probes where if no response, will result in stopping further traffic from being delivered to the endpoint. This values allows endpoints to be taken out of rotation faster or slower than the typical times used in Azure.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Probe.Properties.NumberOfProbes")},
			{
				Name:        "port",
				Description: "The port for communicating the probe. Possible values range from 1 to 65535, inclusive.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Probe.Properties.Port")},
			{
				Name:        "protocol",
				Description: "The protocol of the end point. If 'Tcp' is specified, a received ACK is required for the probe to be successful. If 'Http' or 'Https' is specified, a 200 OK response from the specifies URI is required for the probe to be successful. Possible values include: 'HTTP', 'TCP', 'HTTPS'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Probe.Properties.Protocol")},
			{
				Name:        "request_path",
				Description: "The URI used for requesting health status from the VM. Path is required if a protocol is set to http. Otherwise, it is not allowed. There is no default value.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Probe.Properties.RequestPath")},
			{
				Name:        "load_balancing_rules",
				Description: "The load balancer rules that use this probe.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Probe.Properties.LoadBalancingRules")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Probe.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Probe.ID").Transform(idToAkas),
			},

			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
