package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureFirewall(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_firewall",
		Description: "Azure Firewall",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetNetworkAzureFirewall,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListNetworkAzureFirewall,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the firewall",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AzureFirewall.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a firewall uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AzureFirewall.ID")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AzureFirewall.Etag")},
			{
				Name:        "type",
				Description: "The resource type of the firewall",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AzureFirewall.Type")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the firewall resource",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AzureFirewall.Properties.ProvisioningState"),
			},
			{
				Name:        "firewall_policy_id",
				Description: "The firewallPolicy associated with this azure firewall",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AzureFirewall.Properties.FirewallPolicy.ID")},
			{
				Name:        "hub_private_ip_address",
				Description: "Private IP Address associated with azure firewall",
				Type:        proto.ColumnType_IPADDR,

				Transform: transform.FromField("Description.AzureFirewall.Properties.HubIPAddresses.PrivateIPAddress")},
			{
				Name:        "hub_public_ip_address_count",
				Description: "The number of Public IP addresses associated with azure firewall",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.AzureFirewall.Properties.HubIPAddresses.PublicIPs.Count")},
			{
				Name:        "sku_name",
				Description: "Name of an Azure Firewall SKU",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AzureFirewall.Properties.SKU.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Tier of an Azure Firewall",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AzureFirewall.Properties.SKU.Tier"),
			},
			{
				Name:        "threat_intel_mode",
				Description: "The operation mode for Threat Intelligence",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AzureFirewall.Properties.ThreatIntelMode"),
			},
			{
				Name:        "virtual_hub_id",
				Description: "The virtualHub to which the firewall belongs",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AzureFirewall.Properties.VirtualHub.ID")},
			{
				Name:        "additional_properties",
				Description: "A collection of additional properties used to further config this azure firewall",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.AzureFirewall.Properties.AdditionalProperties")},
			{
				Name:        "application_rule_collections",
				Description: "A collection of application rule collections used by Azure Firewall",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.AzureFirewall.Properties.ApplicationRuleCollections")},
			{
				Name:        "availability_zones",
				Description: "A collection of availability zones denoting where the resource needs to come from",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.AzureFirewall.Zones")},
			{
				Name:        "hub_public_ip_addresses",
				Description: "A collection of Public IP addresses associated with azure firewall or IP addresses to be retained",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.AzureFirewall.Properties.HubIPAddresses.PublicIPs.Addresses")},
			{
				Name:        "ip_configurations",
				Description: "A collection of IP configuration of the Azure Firewall resource",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(ipConfigurationData),
			},
			{
				Name:        "ip_groups",
				Description: "A collection of IpGroups associated with AzureFirewall",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.AzureFirewall.Properties.IPGroups")},
			{
				Name:        "nat_rule_collections",
				Description: "A collection of NAT rule collections used by Azure Firewall",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.AzureFirewall.Properties.NatRuleCollections")},
			{
				Name:        "network_rule_collections",
				Description: "A collection of network rule collections used by Azure Firewall",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.AzureFirewall.Properties.NetworkRuleCollections")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AzureFirewall.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AzureFirewall.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.AzureFirewall.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AzureFirewall.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				// Check if context has been cancelled or if the limit has been hit (if specified)
				// if there is a limit, it will return the number of rows required to reach this limit
				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.ResourceGroup")},
		}),
	}
}

//// HYDRATE FUNCTIONS ////

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

//// Transform Functions

func ipConfigurationData(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.NetworkAzureFirewall).Description.AzureFirewall

	var output []map[string]interface{}
	for _, firewall := range data.Properties.IPConfigurations {
		objectMap := make(map[string]interface{})
		if firewall.Properties.PrivateIPAddress != nil {
			objectMap["privateIPAddress"] = firewall.Properties.PrivateIPAddress
		}
		if firewall.Properties.PublicIPAddress != nil {
			objectMap["publicIPAddress"] = firewall.Properties.PublicIPAddress
		}
		if firewall.Properties.Subnet != nil {
			objectMap["subnet"] = firewall.Properties.Subnet
		}
		if firewall.Properties.ProvisioningState != nil {
			if *firewall.Properties.ProvisioningState != "" {
				objectMap["provisioningState"] = firewall.Properties.ProvisioningState
			}
		}
		output = append(output, objectMap)
	}
	return output, nil
}
