package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzurePrivateEndpoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_private_endpoint",
		Description: "Azure Private Endpoint",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetPrivateEndpoint,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			ParentHydrate: opengovernance.ListResourceGroup,
			Hydrate:       opengovernance.ListPrivateEndpoint,
			KeyColumns:    plugin.OptionalColumns([]string{"resource_group"}),
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the private endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateEndpoint.Name"),
			},
			{
				Name:        "id",
				Description: "The ID of the private endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateEndpoint.ID"),
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateEndpoint.Etag"),
			},
			{
				Name:        "type",
				Description: "The type of the private endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateEndpoint.Type"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the private endpoint resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateEndpoint.Properties.ProvisioningState").Transform(transform.ToString),
			},
			{
				Name:        "custom_network_interface_name",
				Description: "The custom name of the network interface attached to the private endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateEndpoint.Properties.CustomNetworkInterfaceName"),
			},
			{
				Name:        "location",
				Description: "The location of the private endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateEndpoint.Location").Transform(toLower),
			},
			{
				Name:        "extended_location",
				Description: "The extended location of the private endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateEndpoint.ExtendedLocation").Transform(toLower),
			},
			{
				Name:        "subnet",
				Description: "The ID of the subnet from which the private IP will be allocated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateEndpoint.Properties.Subnet"),
			},
			{
				Name:        "network_interfaces",
				Description: "An array of references to the network interfaces created for this private endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateEndpoint.Properties.NetworkInterfaces"),
			},
			{
				Name:        "private_link_service_connections",
				Description: "A grouping of information about the connection to the remote resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateEndpoint.Properties.PrivateLinkServiceConnections"),
			},
			{
				Name:        "manual_private_link_service_connections",
				Description: "A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateEndpoint.Properties.ManualPrivateLinkServiceConnections"),
			},
			{
				Name:        "custom_dns_configs",
				Description: "An array of custom DNS configurations.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateEndpoint.Properties..CustomDNSConfigs"),
			},
			{
				Name:        "application_security_groups",
				Description: "Application security groups in which the private endpoint IP configuration is included.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateEndpoint.Properties.ApplicationSecurityGroups"),
			},
			{
				Name:        "ip_configurations",
				Description: "A list of IP configurations of the private endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateEndpoint.Properties.IPConfigurations"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateEndpoint.Name"),
			},
			{
				Name:        "tags",
				Description: "Tags associated with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateEndpoint.Tags"),
			},
			{
				Name:        "akas",
				Description: "Array of globally unique identifier strings (also known as) for the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateEndpoint.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: "The Azure region where the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrivateEndpoint.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: "The resource group in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}
