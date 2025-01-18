package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureAutoscaleSetting(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_autoscale_setting",
		Description: "Azure Autoscale Setting",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetAutoscaleSetting,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAutoscaleSetting,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the resource.",
				Transform:   transform.FromField("Description.AutoscaleSettingsResource.Name")},
			{
				Name:        "id",
				Description: "The resource Id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoscaleSettingsResource.ID")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoscaleSettingsResource.Type")},
			{
				Name:        "profiles",
				Description: "Autoscale setting profiles",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoscaleSettingsResource.Properties.Profiles")},
			{
				Name:        "enabled",
				Description: "Whether the autoscale setting is enabled or not.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoscaleSettingsResource.Properties.Enabled")},
			{
				Name:        "notifications",
				Description: "Autoscale setting notifications settings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutoscaleSettingsResource.Properties.Notifications")},
			{
				Name:        "target_resource_location",
				Description: "Autoscale setting target resource location.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoscaleSettingsResource.Properties.TargetResourceLocation")},
			{
				Name:        "target_resource_uri",
				Description: "Autoscale setting target resource uri.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoscaleSettingsResource.Properties.TargetResourceURI")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AutoscaleSettingsResource.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.AutoscaleSettingsResource.ID").Transform(idToAkas),
			},

			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup")},
		}),
	}
}
