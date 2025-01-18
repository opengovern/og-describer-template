package azure

import (
	"context"
	"strings"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureStorageBlob(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_storage_blob",
		Description: "Azure Storage Blob",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStorageBlob,
		},
		Columns: azureOGColumns([]*plugin.Column{
			// Basic info
			{
				Name:        "name",
				Description: "The friendly name that identifies the blob.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Name")},
			{
				Name:        "storage_account_name",
				Description: "The friendly name that identifies the storage account, in which the blob is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccountName")},
			{
				Name:        "container_name",
				Description: "The friendly name that identifies the container, in which the blob has been uploaded.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ContainerName")},
			{
				Name:        "type",
				Description: "Specifies the type of the blob.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Blob.Properties.BlobType"),
			},
			{
				Name:        "is_snapshot",
				Description: "Specifies whether the resource is snapshot of a blob, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.IsSnapshot")},
			{
				Name:        "access_tier",
				Description: "The tier of the blob.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Blob.Properties.AccessTier"),
			},
			{
				Name:        "creation_time",
				Description: "Indicates the time, when the blob was uploaded.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Blob.Properties.CreationTime")},
			{
				Name:        "deleted",
				Description: "Specifies whether the blob was deleted, or not.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Blob.Deleted"), Default: false,
			},
			{
				Name:        "deleted_time",
				Description: "Specifies the deletion time of blob container.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Blob.Properties.DeletedTime")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Blob.Properties.Etag"),
			},
			{
				Name:        "last_modified",
				Description: "Specifies the date and time the container was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Blob.Properties.LastModified")},
			{
				Name:        "snapshot",
				Description: "Specifies the time, when the snapshot is taken.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Snapshot")},
			{
				Name:        "version_id",
				Description: "Specifies the version id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.VersionID")},
			{
				Name:        "server_encrypted",
				Description: "Indicates whether the blob is encrypted on the server, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Blob.Properties.ServerEncrypted")},
			{
				Name:        "encryption_scope",
				Description: "The name of the encryption scope under which the blob is encrypted.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.EncryptionScope")},
			{
				Name:        "encryption_key_sha256",
				Description: "The SHA-256 hash of the provided encryption key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.CustomerProvidedKeySha256")},
			{
				Name:        "is_current_version",
				Description: "Specifies whether the blob container was deleted, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Blob.IsCurrentVersion")},
			{
				Name:        "access_tier_change_time",
				Description: "Species the time, when the access tier has been updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Blob.Properties.AccessTierChangeTime")},
			{
				Name:        "access_tier_inferred",
				Description: "Indicates whether the access tier was inferred by the service.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Blob.Properties.AccessTierInferred")},
			{
				Name:        "blob_sequence_number",
				Description: "Specifies the sequence number for page blob used for coordinating concurrent writes.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Blob.Properties.BlobSequenceNumber")},
			{
				Name:        "content_length",
				Description: "Specifies the size of the content returned.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Blob.Properties.ContentLength")},
			{
				Name:        "cache_control",
				Description: "Indicates the cache control specified for the blob.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.CacheControl")},
			{
				Name:        "content_disposition",
				Description: "Specifies additional information about how to process the response payload, and also can be used to attach additional metadata.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.ContentDisposition")},
			{
				Name:        "content_encoding",
				Description: "Indicates content encoding specified for the blob.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.ContentEncoding")},
			{
				Name:        "content_language",
				Description: "Indicates content language specified for the blob.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.ContentLanguage")},
			{
				Name:        "content_md5",
				Description: "If the content_md5 has been set for the blob, this response header is stored so that the client can check for message content integrity.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Blob.Properties.ContentMD5")},
			{
				Name:        "content_type",
				Description: "Specifies the content type specified for the blob. If no content type was specified, the default content type is application/octet-stream.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.ContentType")},
			{
				Name:        "copy_completion_time",
				Description: "Conclusion time of the last attempted Copy Blob operation where this blob was the destination blob.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Blob.Properties.CopyCompletionTime")},
			{
				Name:        "copy_id",
				Description: "A String identifier for the last attempted Copy Blob operation where this blob was the destination blob.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.CopyID")},
			{
				Name:        "copy_progress",
				Description: "Contains the number of bytes copied and the total bytes in the source in the last attempted Copy Blob operation where this blob was the destination blob.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.CopyProgress")},
			{
				Name:        "copy_source",
				Description: "An URL up to 2 KB in length that specifies the source blob used in the last attempted Copy Blob operation where this blob was the destination blob.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.CopySource")},
			{
				Name:        "copy_status",
				Description: "Specifies the state of the copy operation identified by Copy ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.CopyStatus")},
			{
				Name:        "copy_status_description",
				Description: "Describes cause of fatal or non-fatal copy operation failure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.CopyStatusDescription")},
			{
				Name:        "destination_snapshot",
				Description: "Included if the blob is incremental copy blob or incremental copy snapshot, if x-ms-copy-status is success. Snapshot time of the last successful incremental copy snapshot for this blob.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Properties.DestinationSnapshot")},
			{
				Name:        "lease_duration",
				Description: "Specifies whether the lease is of infinite or fixed duration, when a blob is leased.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Blob.Properties.LeaseDuration"),
			},
			{
				Name:        "lease_state",
				Description: "Specifies lease state of the blob.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Blob.Properties.LeaseState"),
			},
			{
				Name:        "lease_status",
				Description: "Specifies the lease status of the blob.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Blob.Properties.LeaseStatus"),
			},
			{
				Name:        "incremental_copy",
				Description: "Copies the snapshot of the source page blob to a destination page blob. The snapshot is copied such that only the differential changes between the previously copied snapshot are transferred to the destination.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Blob.Properties.IncrementalCopy")},
			{
				Name:        "is_sealed",
				Description: "Indicate if the append blob is sealed or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Blob.Properties.IsSealed")},
			{
				Name:        "remaining_retention_days",
				Description: "The number of days that the blob will be retained before being permanently deleted by the service.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Blob.Properties.RemainingRetentionDays")},
			{
				Name:        "archive_status",
				Description: "Specifies the archive status of the blob.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Blob.Properties.ArchiveStatus"),
			},
			{
				Name:        "blob_tag_set",
				Description: "A list of blob tags.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Blob.BlobTags.BlobTagSet")},
			{
				Name:        "metadata",
				Description: "A name-value pair to associate with the container as metadata.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Blob.Metadata")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blob.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(blobDataToAka),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Metadata.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.

					//// TRANSFORM FUNCTIONS
					FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}

func blobDataToAka(_ context.Context, d *transform.TransformData) (interface{}, error) {
	blob := d.HydrateItem.(opengovernance.StorageBlob)

	// Build resource aka
	akas := []string{"azure:///subscriptions/" + blob.Metadata.SubscriptionID + "/resourceGroups/" + blob.Description.ResourceGroup + "/providers/Microsoft.Storage/storageAccounts/" + blob.Description.AccountName + "/blobServices/default/containers/" + blob.Description.ContainerName + "/blobs/" + blob.Description.Blob.Name, "azure:///subscriptions/" + blob.Metadata.SubscriptionID + "/resourcegroups/" + strings.ToLower(blob.Description.ResourceGroup) + "/providers/microsoft.storage/storageaccounts/" + strings.ToLower(blob.Description.AccountName) + "/blobservices/default/containers/" + strings.ToLower(blob.Description.ContainerName) + "/blobs/" + strings.ToLower(blob.Description.Blob.Name)}

	return akas, nil
}
