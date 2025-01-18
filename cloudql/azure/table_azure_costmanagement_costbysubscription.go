package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureCostManagementCostBySubscription(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_costmanagement_costbysubscription",
		Description: "Azure CostManagement CostBySubscription",
		Get: &plugin.GetConfig{
			Hydrate:    opengovernance.GetCostManagementCostBySubscription,
			KeyColumns: plugin.OptionalColumns([]string{"id"}),
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCostManagementCostBySubscription,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the costbysubscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ResourceID")},
			{
				Name:        "name",
				Description: "The name of the costbysubscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CostManagementCostBySubscription")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.CostManagementCostBySubscription").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
