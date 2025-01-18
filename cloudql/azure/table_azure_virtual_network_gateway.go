package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureVirtualNetworkGateway(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_virtual_network_gateway",
		Description: "Azure Virtual Network Gateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetVirtualNetworkGateway,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListVirtualNetworkGateway,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the virtual network gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualNetworkGateway.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a virtual network gateway uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualNetworkGateway.ID")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualNetworkGateway.Type")},
			{
				Name:        "gateway_type",
				Description: "The type of this virtual network gateway. Possible values include: 'Vpn', 'ExpressRoute'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.GatewayType"),
			},
			{
				Name:        "vpn_type",
				Description: "The type of this virtual network gateway. Valid values are: 'PolicyBased', 'RouteBased'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.VPNType"),
			},
			{
				Name:        "vpn_gateway_generation",
				Description: "The generation for this virtual network gateway. Must be None if gatewayType is not VPN. Valid values are: 'None', 'Generation1', 'Generation2'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.VPNGatewayGeneration"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the virtual network gateway resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.ProvisioningState"),
			},
			{
				Name:        "active_active",
				Description: "Indicates whether virtual network gateway configured with active-active mode, or not. If true, each Azure gateway instance will have a unique public IP address, and each will establish an IPsec/IKE S2S VPN tunnel to your on-premises VPN device specified in your local network gateway and connection.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.Active")},
			{
				Name:        "enable_bgp",
				Description: "Indicates whether BGP is enabled for this virtual network gateway, or not.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.EnableBgp")},
			{
				Name:        "enable_dns_forwarding",
				Description: "Indicates whether DNS forwarding is enabled, or not.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.EnableDNSForwarding")},
			{
				Name:        "enable_private_ip_address",
				Description: "Indicates whether private IP needs to be enabled on this gateway for connections or not.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.EnablePrivateIPAddress")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualNetworkGateway.Etag")},
			{
				Name:        "gateway_default_site",
				Description: "The reference to the LocalNetworkGateway resource, which represents local network site having default routes. Assign Null value in case of removing existing default site setting.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.GatewayDefaultSite.ID")},
			{
				Name:        "inbound_dns_forwarding_endpoint",
				Description: "The IP address allocated by the gateway to which dns requests can be sent.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.InboundDNSForwardingEndpoint")},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the virtual network gateway resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.ResourceGUID")},
			{
				Name:        "sku_name",
				Description: "Gateway SKU name.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.SKU.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Gateway SKU tier.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.SKU.Tier"),
			},
			{
				Name:        "sku_capacity",
				Description: "Gateway SKU capacity.",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.SKU.Capacity")},
			{
				Name:        "bgp_settings",
				Description: "Virtual network gateway's BGP speaker settings.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.BgpSettings")},
			{
				Name:        "custom_routes_address_prefixes",
				Description: "A list of address blocks reserved for this virtual network in CIDR notation.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.CustomRoutes.AddressPrefixes")},
			{
				Name:        "gateway_connections",
				Description: "A list of virtual network gateway connection resources that exists in a resource group.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.VirtualNetworkGatewayConnection")},
			{
				Name:        "ip_configurations",
				Description: "IP configurations for virtual network gateway.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.IPConfigurations")},
			{
				Name:        "vpn_client_configuration",
				Description: "The reference to the VpnClientConfiguration resource which represents the P2S VpnClient configurations.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Properties.VPNClientConfiguration")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualNetworkGateway.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.VirtualNetworkGateway.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualNetworkGateway.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
