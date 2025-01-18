package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureAppServicePlan(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_app_service_plan",
		Description: "Azure App Service Plan",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetAppServicePlan,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAppServicePlan,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the app service plan",
				Transform:   transform.FromField("Description.Plan.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify an app service plan uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Plan.ID"),
			},
			{
				Name:        "kind",
				Description: "Contains the kind of the resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Plan.Kind")},
			{
				Name:        "type",
				Description: "The resource type of the app service plan",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Plan.Type")},
			{
				Name:        "hyper_v",
				Description: "Specify whether resource is Hyper-V container app service plan",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Plan.Properties.HyperV"), Default: false,
			},
			{
				Name:        "is_spot",
				Description: "Specify whether this App Service Plan owns spot instances, or not",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Plan.Properties.IsSpot"), Default: false,
			},
			{
				Name:        "is_xenon",
				Description: "Specify whether resource is Hyper-V container app service plan",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Plan.Properties.IsXenon"), Default: false,
			},
			{
				Name:        "maximum_elastic_worker_count",
				Description: "Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Plan.Properties.MaximumElasticWorkerCount")},
			{
				Name:        "maximum_number_of_workers",
				Description: "Maximum number of instances that can be assigned to this App Service plan",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Plan.Properties.MaximumNumberOfWorkers")},
			{
				Name:        "per_site_scaling",
				Description: "Specify whether apps assigned to this App Service plan can be scaled independently",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Plan.Properties.PerSiteScaling"), Default: false,
			},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the App Service Environment",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Plan.Properties.ProvisioningState"),
			},
			{
				Name:        "reserved",
				Description: "Specify whether the resource is Linux app service plan, or not",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Plan.Properties.Reserved"), Default: false,
			},
			{
				Name:        "sku_capacity",
				Description: "Current number of instances assigned to the resource.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Plan.SKU.Capacity")},
			{
				Name:        "sku_family",
				Description: "Family code of the resource SKU",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Plan.SKU.Family")},
			{
				Name:        "sku_name",
				Description: "Name of the resource SKU",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Plan.SKU.Name")},
			{
				Name:        "sku_size",
				Description: "Size specifier of the resource SKU",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Plan.SKU.Size")},
			{
				Name:        "sku_tier",
				Description: "Service tier of the resource SKU",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Plan.SKU.Tier")},
			{
				Name:        "status",
				Description: "App Service plan status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Plan.Properties.Status"),
			},
			{
				Name:        "apps",
				Description: "Site a web app, a mobile app backend, or an API app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Apps")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Plan.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Plan.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Plan.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Plan.Location").Transform(toLower),
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
