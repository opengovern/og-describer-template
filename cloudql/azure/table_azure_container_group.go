package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureContainerGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_container_group",
		Description: "Azure Container Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetContainerInstanceContainerGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListContainerInstanceContainerGroup,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerGroup.Name")},
			{
				Name:        "id",
				Description: "The resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerGroup.ID")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerGroup.Type")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the container group. This only appears in the response.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.ProvisioningState"),
			},
			{
				Name:        "restart_policy",
				Description: "Restart policy for all containers within the container group. Possible values include: 'ContainerGroupRestartPolicyAlways', 'ContainerGroupRestartPolicyOnFailure', 'ContainerGroupRestartPolicyNever'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.RestartPolicy"),
			},

			{
				Name:        "sku",
				Description: "The SKU for a container group. Possible values include: 'ContainerGroupSkuStandard', 'ContainerGroupSkuDedicated'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.SKU"),
			},
			{
				Name:        "os_type",
				Description: "The operating system type required by the containers in the container group. Possible values include: 'OperatingSystemTypesWindows', 'OperatingSystemTypesLinux'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.OSType"),
			},
			{
				Name:        "encryption_properties",
				Description: "The encryption settings of container registry.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.EncryptionProperties"),
			},
			{
				Name:        "containers",
				Description: "The containers within the container group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.Containers"),
			},
			{
				Name:        "ip_address",
				Description: "The IP address type of the container group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.IPAddress"),
			},
			{
				Name:        "volumes",
				Description: "The instance view of the container group. Only valid in response.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.Volumes"),
			},
			{
				Name:        "instance_view",
				Description: "The instance view of the container group. Only valid in response.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.InstanceView"),
			},
			{
				Name:        "diagnostics",
				Description: "The diagnostic information for a container group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.Diagnostics"),
			},
			{
				Name:        "subnet_ids",
				Description: "The subnet resource IDs for a container group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.SubnetIDs"),
			},
			{
				Name:        "dns_config",
				Description: "The DNS config information for a container group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.DNSConfig"),
			},
			{
				Name:        "init_containers",
				Description: "The init containers for a container group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.InitContainers"),
			},
			{
				Name:        "image_registry_credentials",
				Description: "The image registry credentials by which the container group is created from.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Properties.ImageRegistryCredentials"),
			},
			{
				Name:        "identity",
				Description: "The identity of the container group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Identity"),
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerGroup.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ContainerGroup.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerGroup.Location").Transform(toLower),
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
