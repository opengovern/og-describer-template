package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureCostManagementCostByResourceType(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_costmanagement_costbyresourcetype",
		Description: "Azure CostManagement CostByResourceType",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCostManagementCostByResourceType,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the costbyresourcetype.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ResourceID")},
			{
				Name:        "usage_date",
				Description: "Usage date",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.CostManagementCostByResourceType.UsageDate")},
			{
				Name:        "cost",
				Description: "Cost",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.CostManagementCostByResourceType.Cost")},
			{
				Name:        "resource_type",
				Description: "Resource type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata.ResourceType")},
			{
				Name:        "service_name",
				Description: "Service Name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CostManagementCostByResourceType.ServiceName")},
			{
				Name:        "publisher_type",
				Description: "Publisher Type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CostManagementCostByResourceType.PublisherType")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.CostManagementCostByResourceType").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
