package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION ////

func tableAzureComputeSnapshot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_snapshot",
		Description: "Azure Compute Snapshot",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetComputeSnapshots,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeSnapshots,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the snapshot",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.Name")},
			{
				Name:        "id",
				Description: "The unique id identifying the resource in subscription",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.ID")},
			{
				Name:        "type",
				Description: "The type of the resource in Azure",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.Type")},
			{
				Name:        "provisioning_state",
				Description: "The disk provisioning state",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Snapshot.Properties.ProvisioningState")},
			{
				Name:        "create_option",
				Description: "Specifies the possible sources of a disk's creation",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.CreationData.CreateOption"),
			},
			{
				Name:        "disk_access_id",
				Description: "ARM id of the DiskAccess resource for using private endpoints on disks",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.DiskAccessID")},
			{
				Name:        "disk_encryption_set_id",
				Description: "ResourceId of the disk encryption set to use for enabling encryption at rest",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.Encryption.DiskEncryptionSetID")},
			{
				Name:        "disk_size_bytes",
				Description: "The size of the disk in bytes",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.Snapshot.Properties.DiskSizeBytes")},
			{
				Name:        "disk_size_gb",
				Description: "The size of the disk to create",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.Snapshot.Properties.DiskSizeGB")},
			{
				Name:        "encryption_setting_collection_enabled",
				Description: "Specifies whether the encryption is enables, or not",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Snapshot.Properties.EncryptionSettingsCollection.Enabled")},
			{
				Name:        "encryption_setting_version",
				Description: "Describes what type of encryption is used for the disks",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.EncryptionSettingsCollection.EncryptionSettingsVersion")},
			{
				Name:        "encryption_type",
				Description: "The type of the encryption",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.Encryption.Type"),
			},
			{
				Name:        "gallery_image_reference_id",
				Description: "A relative uri containing either a Platform Image Repository or user image reference",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.CreationData.GalleryImageReference.ID")},
			{
				Name:        "gallery_reference_lun",
				Description: "Specifies the index that indicates which of the data disks in the image to use",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.Snapshot.Properties.CreationData.GalleryImageReference.Lun")},
			{
				Name:        "hyperv_generation",
				Description: "Specifies the hypervisor generation of the Virtual Machine",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.HyperVGeneration"),
			},
			{
				Name:        "image_reference_id",
				Description: "A relative uri containing either a Platform Image Repository or user image reference",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.CreationData.ImageReference.ID")},
			{
				Name:        "image_reference_lun",
				Description: "Specifies the index that indicates which of the data disks in the image to use",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.Snapshot.Properties.CreationData.ImageReference.Lun")},
			{
				Name:        "incremental",
				Description: "Specifies whether a snapshot is incremental, or not",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Snapshot.Properties.Incremental")},
			{
				Name:        "network_access_policy",
				Description: "Contains the type of access",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.NetworkAccessPolicy"),
			},
			{
				Name:        "os_type",
				Description: "Contains the type of operating system",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.OSType"),
			},
			{
				Name:        "sku_name",
				Description: "The snapshot sku name",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.SKU.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The sku tier",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.SKU.Tier")},
			{
				Name:        "source_resource_id",
				Description: "ARM id of the source snapshot or disk",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.CreationData.SourceResourceID")},
			{
				Name:        "source_unique_id",
				Description: "An unique id identifying the source of this resource",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.CreationData.SourceUniqueID")},
			{
				Name:        "source_uri",
				Description: "An URI of a blob to be imported into a managed disk",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.CreationData.SourceURI")},
			{
				Name:        "storage_account_id",
				Description: "The Azure Resource Manager identifier of the storage account containing the blob to import as a disk",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.CreationData.StorageAccountID")},
			{
				Name:        "time_created",
				Description: "The time when the snapshot was created",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Snapshot.Properties.TimeCreated").Transform(convertDateToTime),
			},
			{
				Name:        "unique_id",
				Description: "An unique Guid identifying the resource",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Properties.UniqueID")},
			{
				Name:        "upload_size_bytes",
				Description: "The size of the contents of the upload including the VHD footer",
				Type:        proto.ColumnType_INT,

				Transform: transform.FromField("Description.Snapshot.Properties.CreationData.UploadSizeBytes")},
			{
				Name:        "encryption_settings",
				Description: "A list of encryption settings, one for each disk volume",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Snapshot.Properties.EncryptionSettingsCollection.EncryptionSettings")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Snapshot.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Snapshot.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Snapshot.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION ////

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
