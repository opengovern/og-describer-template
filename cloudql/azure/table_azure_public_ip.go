package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzurePublicIP(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_public_ip",
		Description: "Azure Public IP",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetPublicIPAddress,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPublicIPAddress,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the public ip",
				Transform:   transform.FromField("Description.PublicIPAddress.Name"),
			},
			{
				Name:        "id",
				Description: "Contains ID to identify a public ip uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.ID"),
			},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Etag"),
			},
			{
				Name:        "type",
				Description: "The resource type of the public ip",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Type"),
			},
			{
				Name:        "provisioning_state",
				Description: "The resource type of the public ip",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.ProvisioningState"),
			},
			{
				Name:        "ddos_custom_policy_id",
				Description: "The DDoS custom policy associated with the public IP",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.DdosSettings.DdosCustomPolicy.ID"),
			},
			{
				Name:        "ddos_settings_protected_ip",
				Description: "Indicates whether DDoS protection is enabled on the public IP, or not",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.DdosSettings.ProtectedIP"),
			},
			{
				Name:        "ddos_settings_protection_coverage",
				Description: "The DDoS protection policy customizability of the public IP",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.DdosSettings.ProtectionCoverage"),
			},
			{
				Name:        "dns_settings_domain_name_label",
				Description: "Contains the domain name label",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.DNSSettings.DomainNameLabel"),
			},
			{
				Name:        "dns_settings_fqdn",
				Description: "The Fully Qualified Domain Name of the A DNS record associated with the public IP",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.DNSSettings.Fqdn"),
			},
			{
				Name:        "dns_settings_reverse_fqdn",
				Description: "Contains the reverse FQDN",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.DNSSettings.ReverseFqdn"),
			},
			{
				Name:        "idle_timeout_in_minutes",
				Description: "The idle timeout of the public IP address",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.IdleTimeoutInMinutes"),
			},
			{
				Name:        "ip_address",
				Description: "The IP address associated with the public IP address resource",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.IPAddress"),
			},
			{
				Name:        "ip_configuration_id",
				Description: "Contains the IP configuration ID",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.IPConfiguration.ID"),
			},
			{
				Name:        "public_ip_address_version",
				Description: "Contains the public IP address version",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.PublicIPAddressVersion"),
			},
			{
				Name:        "public_ip_allocation_method",
				Description: "Contains the public IP address allocation method",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.PublicIPAllocationMethod"),
			},
			{
				Name:        "public_ip_prefix_id",
				Description: "The Public IP Prefix this Public IP Address should be allocated from",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.PublicIPPrefix.ID"),
			},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the public ip resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Properties.ResourceGUID"),
			},
			{
				Name:        "sku_name",
				Description: "Name of a public IP address SKU",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.SKU.Name"),
			},
			{
				Name:        "ip_tags",
				Description: "A list of tags associated with the public IP address",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PublicIPAddress.PublicIPAddressPropertiesFormat.IPTags"),
			},
			{
				Name:        "zones",
				Description: "A collection of availability zones denoting the IP allocated for the resource needs to come from",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PublicIPAddress.Zones"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Name"),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PublicIPAddress.Tags"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PublicIPAddress.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PublicIPAddress.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}
