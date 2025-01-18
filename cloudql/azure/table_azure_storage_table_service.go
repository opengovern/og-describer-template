package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureStorageTableService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_storage_table_service",
		Description: "Azure Storage Table Service",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"storage_account_name", "resource_group"}),
			Hydrate:    opengovernance.GetStorageTableService,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStorageTableService,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the table service",
				Transform:   transform.FromField("Description.TableService.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a table service uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TableService.ID"),
			},
			{
				Name:        "storage_account_name",
				Description: "An unique read-only string that changes whenever the resource is updated",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AccountName")},
			{
				Name:        "type",
				Description: "Type of the resource",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.TableService.Type")},
			{
				Name:        "cors_rules",
				Description: "A list of CORS rules",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TableService.TableServiceProperties.Cors.CorsRules")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TableService.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.TableService.ID").Transform(idToAkas),
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

				Transform: transform.FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}
