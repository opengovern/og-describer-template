package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureDBforPostgreSQLFlexibleServers(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_postgresql_flexible_server",
		Description: "Azure PostgreSQL Flexible Server",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetPostgresqlFlexibleServer,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPostgresqlFlexibleServer,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the flexibleservers.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.ID")},
			{
				Name:        "name",
				Description: "The name of the flexibleservers.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Name")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Type"),
			},
			{
				Name:        "location",
				Description: "The geo-location where the resource lives.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Location"),
			},
			{
				Name:        "sku",
				Description: "The SKU (pricing tier) of the server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Server.SKU"),
			},
			{
				Name:        "server_properties",
				Description: "Properties of the server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Server.ServerProperties"),
			},
			{
				Name:        "flexible_server_configurations",
				Description: "The server configurations(parameters) details of the server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Server.ServerConfigurations"),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.Server.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.Server.ID").Transform(idToAkas),
			},
		}),
	}
}

// func tableAzurePostgreSqlFlexibleServer(_ context.Context) *plugin.Table {
// 	return &plugin.Table{
// 		Name:        "azure_postgresql_flexible_server",
// 		Description: "Azure PostgreSQL Flexible Server",
// 		Get: &plugin.GetConfig{
// 			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
// 			Hydrate:    getPostgreSqlFlexibleServer,
// 			IgnoreConfig: &plugin.IgnoreConfig{
// 				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
// 			},
// 		},
// 		List: &plugin.ListConfig{
// 			ParentHydrate: listResourceGroups,
// 			Hydrate:       listPostgreSqlFlexibleServers,
// 		},
// 		Columns: azureColumns([]*plugin.Column{
// 			{
// 				Name:        "name",
// 				Description: "The name of the resource.",
// 				Type:        proto.ColumnType_STRING,
// 			},
// 			{
// 				Name:        "id",
// 				Description: "Fully qualified resource ID for the resource.",
// 				Type:        proto.ColumnType_STRING,
// 				Transform:   transform.FromGo(),
// 			},
// 			{
// 				Name:        "type",
// 				Description: "The type of the resource.",
// 				Type:        proto.ColumnType_STRING,
// 			},
// 			{
// 				Name:        "location",
// 				Description: "The geo-location where the resource lives.",
// 				Type:        proto.ColumnType_STRING,
// 			},
// 			// We have raised a support request for this as SystemData is always null
// 			// {
// 			// 	Name:        "system_data",
// 			// 	Description: "The system data for the server.",
// 			// 	Type:        proto.ColumnType_JSON,
// 			// },
// 			{
// 				Name:        "sku",
// 				Description: "The SKU (pricing tier) of the server.",
// 				Type:        proto.ColumnType_JSON,
// 			},
// 			{
// 				Name:        "server_properties",
// 				Description: "Properties of the server.",
// 				Type:        proto.ColumnType_JSON,
// 				Transform:   transform.FromField("ServerProperties").Transform(extractPostgresFlexibleServerProperties),
// 			},
// 			{
// 				Name:        "flexible_server_configurations",
// 				Description: "The server configurations(parameters) details of the server.",
// 				Type:        proto.ColumnType_JSON,
// 				Hydrate:     listPostgreSQLFlexibleServersConfigurations,
// 				Transform:   transform.FromValue(),
// 			},
// 			// Steampipe standard columns
// 			{
// 				Name:        "title",
// 				Description: ColumnDescriptionTitle,
// 				Type:        proto.ColumnType_STRING,
// 				Transform:   transform.FromField("ResourceName"),
// 			},
// 			{
// 				Name:        "tags",
// 				Description: ColumnDescriptionTags,
// 				Type:        proto.ColumnType_JSON,
// 			},
// 			{
// 				Name:        "akas",
// 				Description: ColumnDescriptionAkas,
// 				Type:        proto.ColumnType_JSON,
// 				Transform:   transform.FromField("ResourceID").Transform(idToAkas),
// 			},
// 			{
// 				Name:        "resource_group",
// 				Description: ColumnDescriptionResourceGroup,
// 				Type:        proto.ColumnType_STRING,
// 				Transform:   transform.FromField("ResourceID").Transform(extractResourceGroupFromID),
// 			},
// 		}),
// 	}
// }
