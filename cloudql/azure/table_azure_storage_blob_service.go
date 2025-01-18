package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureStorageBlobService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_storage_blob_service",
		Description: "Azure Storage Blob Service",
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetStorageBlobService,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound"}),
			},
			KeyColumns: plugin.AllColumns([]string{"id"}),
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStorageBlobService,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the blob",
				Transform:   transform.FromField("Description.BlobService.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a blob uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BlobService.ID"),
			},
			{
				Name:        "storage_account_name",
				Description: "A unique read-only string that changes whenever the resource is updated",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccountName")},
			{
				Name:        "type",
				Description: "Type of the resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BlobService.Type")},
			{
				Name:        "sku_name",
				Description: "The sku name",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.BlobService.SKU.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Contains the sku tier",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.BlobService.SKU.Tier"),
			},
			{
				Name:        "automatic_snapshot_policy_enabled",
				Description: "Specifies whether automatic snapshot creation is enabled, or not",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.BlobService.BlobServiceProperties.AutomaticSnapshotPolicyEnabled"), Default: false,
			},
			{
				Name:        "change_feed_enabled",
				Description: "Specifies whether change feed event logging is enabled for the Blob service",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.BlobService.BlobServiceProperties.ChangeFeed.Enabled"), Default: false,
			},
			{
				Name:        "default_service_version",
				Description: "Indicates the default version to use for requests to the Blob service if an incoming request’s version is not specified",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BlobService.BlobServiceProperties.DefaultServiceVersion")},
			{
				Name:        "is_versioning_enabled",
				Description: "Specifies whether the versioning is enabled, or not",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.BlobService.BlobServiceProperties.IsVersioningEnabled"), Default: false,
			},
			{
				Name:        "container_delete_retention_policy",
				Description: "The blob service properties for container soft delete",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BlobService.BlobServiceProperties.ContainerDeleteRetentionPolicy")},
			{
				Name:        "cors_rules",
				Description: "A list of CORS rules for a storage account’s Blob service",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BlobService.BlobServiceProperties.Cors.CorsRules")},
			{
				Name:        "delete_retention_policy",
				Description: "The blob service properties for blob soft delete",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BlobService.BlobServiceProperties.ContainerDeleteRetentionPolicy")},
			{
				Name:        "restore_policy",
				Description: "The blob service properties for blob restore policy",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BlobService.BlobServiceProperties.RestorePolicy")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BlobService.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.BlobService.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
