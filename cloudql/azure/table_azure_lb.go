package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureLoadBalancer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_lb",
		Description: "Azure Load Balancer",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetLoadBalancer,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLoadBalancer,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Name")},
			{
				Name:        "id",
				Description: "The resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.ID")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the load balancer resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Type")},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Etag")},
			{
				Name:        "extended_location_name",
				Description: "The name of the extended location.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.ExtendedLocation.Name")},
			{
				Name:        "extended_location_type",
				Description: "The type of the extended location. Possible values include: 'ExtendedLocationTypesEdgeZone'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LoadBalancer.ExtendedLocation.Type"),
			},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the load balancer resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.ResourceGUID")},
			{
				Name:        "sku_name",
				Description: "Name of the load balancer SKU. Possible values include: 'Basic', 'Standard', 'Gateway'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LoadBalancer.SKU.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Tier of the load balancer SKU. Possible values include: 'Regional', 'Global'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LoadBalancer.SKU.Tier"),
			},
			{
				Name:        "backend_address_pools",
				Description: "Collection of backend address pools used by the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.BackendAddressPools")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSetting")},
			{
				Name:        "frontend_ip_configurations",
				Description: "Object representing the frontend IPs to be used for the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.FrontendIPConfigurations")},
			{
				Name:        "inbound_nat_pools",
				Description: "Defines an external port range for inbound NAT to a single backend port on NICs associated with the load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on the Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.InboundNatPools")},
			{
				Name:        "inbound_nat_rules",
				Description: "Collection of inbound NAT Rules used by the load balancer. Defining inbound NAT rules on the load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.InboundNatRules")},
			{
				Name:        "load_balancing_rules",
				Description: "Object collection representing the load balancing rules Gets the provisioning.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.LoadBalancingRules")},
			{
				Name:        "outbound_rules",
				Description: "The outbound rules.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.OutboundRules")},
			{
				Name:        "probes",
				Description: "Collection of probe objects used in the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Properties.Probes")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.LoadBalancer.ID").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LoadBalancer.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
