package azure

import (
	"context"
	"strings"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureAppServiceEnvironment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_app_service_environment",
		Description: "Azure App Service Environment",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetAppServiceEnvironment,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAppServiceEnvironment,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the app service environment",
				Transform:   transform.FromField("Description.AppServiceEnvironmentResource.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify an app service environment uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppServiceEnvironmentResource.ID")},
			{
				Name:        "kind",
				Description: "Contains the kind of the resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppServiceEnvironmentResource.Kind")},
			{
				Name:        "type",
				Description: "The resource type of the app service environment",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppServiceEnvironmentResource.Type")},
			{
				Name:        "status",
				Description: "Current status of the App Service Environment",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AppServiceEnvironmentResource.Properties.Status"),
			},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the App Service Environment",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AppServiceEnvironmentResource.Properties.ProvisioningState"),
			},
			{
				Name:        "default_front_end_scale_factor",
				Description: "Default Scale Factor for FrontEnds",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AppServiceEnvironmentResource.Properties.FrontEndScaleFactor")},
			{
				Name:        "dynamic_cache_enabled",
				Description: "Indicates whether the dynamic cache is enabled or not",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AppServiceEnvironmentResource.Properties.EnableAcceleratedNetworking"), //TODO-Saleh ? Set this regarding to the last azure sdk version description
				Default:     false,
			},
			{
				Name:        "front_end_scale_factor",
				Description: "Scale factor for front-ends",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.AppServiceEnvironmentResource.Properties.FrontEndScaleFactor")},
			{
				Name:        "has_linux_workers",
				Description: "Indicates whether an ASE has linux workers or not",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.AppServiceEnvironmentResource.Properties.HasLinuxWorkers"), Default: false,
			},
			{
				Name:        "internal_load_balancing_mode",
				Description: "Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AppServiceEnvironmentResource.Properties.InternalLoadBalancingMode"),
			},
			{
				Name:        "is_healthy_environment",
				Description: "Indicates whether the App Service Environment is healthy",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.AppServiceEnvironmentResource.Properties.EnvironmentIsHealthy"), Default: false, //TODO-Saleh ?
			},
			{
				Name:        "suspended",
				Description: "Indicates whether the App Service Environment is suspended or not",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.AppServiceEnvironmentResource.Properties.Suspended"), Default: false,
			},
			{
				Name:        "vnet_name",
				Description: "Name of the Virtual Network for the App Service Environment",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getVNSubnetName)},
			{
				Name:        "vnet_resource_group_name",
				Description: "Name of the resource group where the virtual network is created",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getVNResourceGroupName),
			},
			{
				Name:        "vnet_subnet_name",
				Description: "Name of the subnet of the virtual network",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getVNName)},
			{
				Name:        "cluster_settings",
				Description: "Custom settings for changing the behavior of the App Service Environment.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.AppServiceEnvironmentResource.Properties.ClusterSettings")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppServiceEnvironmentResource.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AppServiceEnvironmentResource.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.AppServiceEnvironmentResource.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AppServiceEnvironmentResource.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// FETCH FUNCTIONS ////

				Transform: transform.FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}

func getVNResourceGroupName(ctx context.Context, d *transform.TransformData) (any, error) {
	virtualNetwork := d.HydrateItem.(opengovernance.AppServiceEnvironment).Description.AppServiceEnvironmentResource.Properties.VirtualNetwork
	resourceGroup := strings.Split(*virtualNetwork.ID, "/")[4]
	return resourceGroup, nil
}

func getVNName(ctx context.Context, d *transform.TransformData) (any, error) {
	virtualNetwork := d.HydrateItem.(opengovernance.AppServiceEnvironment).Description.AppServiceEnvironmentResource.Properties.VirtualNetwork
	split := strings.Split(*virtualNetwork.ID, "/")
	name := split[len(split)-3]
	return name, nil
}

func getVNSubnetName(ctx context.Context, d *transform.TransformData) (any, error) {
	virtualNetwork := d.HydrateItem.(opengovernance.AppServiceEnvironment).Description.AppServiceEnvironmentResource.Properties.VirtualNetwork
	split := strings.Split(*virtualNetwork.ID, "/")
	subnet := split[len(split)-1]
	return subnet, nil
}
