package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureCosmosDBRestorableDatabaseAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_cosmosdb_restorable_database_account",
		Description: "Azure Cosmos DB Restorable Database Account",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCosmosdbRestorableDatabaseAccount,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the restorable database account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Name"),
			},
			{
				Name:        "id",
				Description: "Contains ID to identify a restorable database account uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.ID"),
			},
			{
				Name:        "account_name",
				Description: "The name of the global database account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.AccountName"),
			},
			{
				Name:        "api_type",
				Description: "The API type of the restorable database account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.APIType"),
			},
			{
				Name:        "creation_time",
				Description: "The creation time of the restorable database account.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Account.Properties.CreationTime"),
			},
			{
				Name:        "deletion_time",
				Description: "The time at which the restorable database account has been deleted.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Account.Properties.DeletionTime"),
			},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Type"),
			},
			{
				Name:        "restorable_locations",
				Description: "List of regions where the database account can be restored from.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.RestorableLocations"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Name"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Location").Transform(toLower),
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
