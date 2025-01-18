package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureStorageShareFile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_storage_share_file",
		Description: "Azure Storage Share File",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group", "storage_account_name"}),
			Hydrate:    opengovernance.GetStorageFileShare,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStorageFileShare,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the resource.",
				Transform:   transform.FromField("Description.FileShare.Name")},
			{
				Name:        "storage_account_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the storage account.",
				Transform:   transform.FromField("Description.AccountName")},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Fully qualified resource ID for the resource.",
				Transform:   transform.FromField("Description.FileShare.ID"),
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the resource.",
				Transform:   transform.FromField("Description.FileShare.Type")},
			{
				Name:        "access_tier",
				Type:        proto.ColumnType_STRING,
				Description: "Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool.",
				Transform:   transform.FromField("Description.FileShare.Properties.AccessTier")},
			{
				Name:        "access_tier_change_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Indicates the last modification time for share access tier.",

				Transform: transform.FromField("Description.FileShare.Properties.AccessTierChangeTime").Transform(convertDateToTime),
			},
			{
				Name:        "access_tier_status",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates if there is a pending transition for access tier.",
				Transform:   transform.FromField("Description.FileShare.Properties.AccessTierStatus")},
			{
				Name:        "last_modified_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Returns the date and time the share was last modified.",

				Transform: transform.FromField("Description.FileShare.Properties.LastModifiedTime").Transform(convertDateToTime),
			},
			{
				Name:        "deleted",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the share was deleted.",
				Transform:   transform.FromField("Description.FileShare.Properties.Deleted")},
			{
				Name:        "deleted_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The deleted time if the share was deleted.",

				Transform: transform.FromField("Description.FileShare.Properties.DeletedTime").Transform(convertDateToTime),
			},
			{
				Name:        "enabled_protocols",
				Type:        proto.ColumnType_STRING,
				Description: "The authentication protocol that is used for the file share. Can only be specified when creating a share. Possible values include: 'SMB', 'NFS'.",
				Transform:   transform.FromField("Description.FileShare.Properties.EnabledProtocols")},
			{
				Name:        "remaining_retention_days",
				Type:        proto.ColumnType_INT,
				Description: "Remaining retention days for share that was soft deleted.",
				Transform:   transform.FromField("Description.FileShare.Properties.RemainingRetentionDays")},
			{
				Name:        "root_squash",
				Type:        proto.ColumnType_STRING,
				Description: "The property is for NFS share only. The default is NoRootSquash. Possible values include: 'NoRootSquash', 'RootSquash', 'AllSquash'.",
				Transform:   transform.FromField("Description.FileShare.Properties.RootSquash")},
			{
				Name:        "share_quota",
				Type:        proto.ColumnType_INT,
				Description: "The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400.",
				Transform:   transform.FromField("Description.FileShare.Properties.ShareQuota")},
			{
				Name:        "share_usage_bytes",
				Type:        proto.ColumnType_INT,
				Description: "The approximate size of the data stored on the share. Note that this value may not include all recently created or recently resized files.",
				Transform:   transform.FromField("Description.FileShare.Properties.ShareUsageBytes")},
			{
				Name:        "version",
				Type:        proto.ColumnType_STRING,
				Description: "The version of the share.",
				Transform:   transform.FromField("Description.FileShare.Properties.Version")},
			{
				Name:        "metadata",
				Type:        proto.ColumnType_JSON,
				Description: "A name-value pair to associate with the share as metadata.",
				Transform:   transform.FromField("Description.FileShare.Properties.Metadata")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FileShare.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.FileShare.ID").Transform(idToAkas),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}
