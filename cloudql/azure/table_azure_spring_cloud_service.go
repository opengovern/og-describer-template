package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSpringCloudService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_spring_cloud_service",
		Description: "Azure Spring Cloud Service",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetSpringCloudService,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSpringCloudService,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Name")},
			{
				Name:        "id",
				Description: "Fully qualified resource Id for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.ID")},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the Service. Possible values include: 'Creating', 'Updating', 'Deleting', 'Deleted', 'Succeeded', 'Failed', 'Moving', 'Moved', 'MoveFailed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Properties.AppType")},
			//{
			//	Name:        "service_id",
			//	Description: "Service instance entity GUID which uniquely identifies a created resource.",
			//	Type:        proto.ColumnType_STRING,
			//	Transform:   transform.FromField("Description.ServiceResource.ID")},
			//{
			//	Name:        "sku_name",
			//	Description: "Name of the Sku.",
			//	Type:        proto.ColumnType_STRING,
			//	Transform:   transform.FromField("Description.ServiceResource.SKU.Name")},
			//{
			//	Name:        "sku_tier",
			//	Description: "Tier of the Sku.",
			//	Type:        proto.ColumnType_STRING,
			//	Transform:   transform.FromField("Description.ServiceResource.SKU.Tier")},
			//{
			//	Name:        "sku_capacity",
			//	Description: "Current capacity of the target resource.",
			//	Type:        proto.ColumnType_INT,
			//	Transform:   transform.FromField("Description.ServiceResource.SKU.Capacity")},
			{
				Name:        "version",
				Description: "Version of the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Properties.SpringBootVersion")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource")},
			//{
			//	Name:        "network_profile",
			//	Description: "Network profile of the service.",
			//	Type:        proto.ColumnType_JSON,
			//	Transform:   transform.From(extractSpringCloudServiceNetworkProfile),
			//},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.App.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.App.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.App.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Site.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

type SpringCloudServiceNetworkProfile struct {
	ServiceRuntimeSubnetID             *string
	AppSubnetID                        *string
	ServiceCidr                        *string
	ServiceRuntimeNetworkResourceGroup *string
	AppNetworkResourceGroup            *string
	OutboundPublicIPs                  *[]string
}

//// LIST FUNCTION

// Get the details of the resource group

//// HYDRATE FUNCTIONS

// Handle empty name or resourceGroup

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

// Create session

// If we return the API response directly, the output does not provide
// all the contents of DiagnosticSettings

//// TRANSFORM FUNCTION

// If we return the API response directly, the output does not provide
// all the properties of NetworkProfile
func extractSpringCloudServiceNetworkProfile(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	workspace := d.HydrateItem.(opengovernance.SpringCloudService).Description.Site
	return workspace.Properties, nil
}
