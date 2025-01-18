package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureBlueprintBlueprints(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_blueprint_blueprints",
		Description: "Azure Blueprint Blueprints",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetBlueprint,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBlueprint,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the blueprints.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blueprints.ID")},
			{
				Name:        "name",
				Description: "The name of the blueprints.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blueprint.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Blueprint.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.Blueprint.ID").Transform(idToAkas),
			},
		}),
	}
}
