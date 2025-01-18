package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureComputeImage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_image",
		Description: "Azure Compute Image",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetComputeImage,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeImage,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the image",
				Transform:   transform.FromField("Description.Image.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a image uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.ID"),
			},
			{
				Name:        "type",
				Description: "Type of the resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Type")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the image resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Properties.ProvisioningState")},
			{
				Name:        "hyper_v_generation",
				Description: "Gets the HyperVGenerationType of the VirtualMachine created from the image",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Image.Properties.HyperVGeneration"),
			},
			{
				Name:        "source_virtual_machine_id",
				Description: "Contains the id of the virtual machine",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Properties.SourceVirtualMachine.ID")},
			{
				Name:        "storage_profile_os_disk_blob_uri",
				Description: "Contains uri of the virtual hard disk",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Properties.StorageProfile.OSDisk.BlobURI")},
			{
				Name:        "storage_profile_os_disk_caching",
				Description: "Specifies the caching requirements",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Image.Properties.StorageProfile.OSDisk.Caching"),
			},
			{
				Name:        "storage_profile_os_disk_encryption_set",
				Description: "Specifies the customer managed disk encryption set resource id for the managed image disk",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Properties.StorageProfile.OSDisk.DiskEncryptionSet.ID")},
			{
				Name:        "storage_profile_os_disk_managed_disk_id",
				Description: "Contains the id of the managed disk",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Properties.StorageProfile.OSDisk.ManagedDisk.ID")},
			{
				Name:        "storage_profile_os_disk_size_gb",
				Description: "Specifies the size of empty data disks in gigabytes",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Image.Properties.StorageProfile.OSDisk.DiskSizeGB")},
			{
				Name:        "storage_profile_os_disk_snapshot_id",
				Description: "Contains the id of the snapshot",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Properties.StorageProfile.OSDisk.Snapshot.ID")},
			{
				Name:        "storage_profile_os_disk_storage_account_type",
				Description: "Specifies the storage account type for the managed disk",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Image.Properties.StorageProfile.OSDisk.StorageAccountType"),
			},
			{
				Name:        "storage_profile_os_disk_state",
				Description: "Contains state of the OS",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Image.Properties.StorageProfile.OSDisk.OSState"),
			},
			{
				Name:        "storage_profile_os_disk_type",
				Description: "Specifies the type of the OS that is included in the disk if creating a VM from a custom image",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Image.Properties.StorageProfile.OSDisk.OSType"),
			},
			{
				Name:        "storage_profile_zone_resilient",
				Description: "Specifies whether an image is zone resilient or not",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Image.Properties.StorageProfile.ZoneResilient"), Default: false,
			},
			{
				Name:        "storage_profile_data_disks",
				Description: "A list of parameters that are used to add a data disk to a virtual machine",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.Properties.StorageProfile.DataDisks")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Image.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Image.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Image.ID").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Image.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
