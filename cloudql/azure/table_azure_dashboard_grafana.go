package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureDashboardGrafana(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_dashboard_grafana",
		Description: "Azure Dashboard Grafana",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetDashboardGrafana,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDashboardGrafana,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the grafana.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Grafana.ID")},
			{
				Name:        "name",
				Description: "The name of the grafana.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Grafana.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Grafana.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Grafana.ID").Transform(idToAkas), // or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
			},
		}),
	}
}
