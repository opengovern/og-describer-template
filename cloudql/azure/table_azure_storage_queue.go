package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureStorageQueue(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_storage_queue",
		Description: "Azure Storage Queue",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"storage_account_name", "name", "resource_group"}),
			Hydrate:    opengovernance.GetStorageQueue,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "QueueNotFound", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStorageQueue,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the queue.",
				Transform:   transform.FromField("Description.Queue.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a queue uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Queue.ID"),
			},
			{
				Name:        "storage_account_name",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccountName")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Queue.Type")},
			{
				Name:        "metadata",
				Description: "A name-value pair that represents queue metadata.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Queue.QueueProperties.Metadata")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Queue.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Queue.ID").Transform(idToAkas),
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
