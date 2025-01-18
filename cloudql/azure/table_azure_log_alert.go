package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureLogAlert(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_log_alert",
		Description: "Azure Log Alert",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetLogAlert,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLogAlert,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the resource.",
				Transform:   transform.FromField("Description.ActivityLogAlertResource.Name")},
			{
				Name:        "id",
				Description: "The resource Id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActivityLogAlertResource.ID")},
			{
				Name:        "type",
				Description: "Type of the resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActivityLogAlertResource.Type")},
			{
				Name:        "enabled",
				Description: "Indicates whether this activity log alert is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ActivityLogAlertResource.Properties.Enabled")},
			{
				Name:        "description",
				Description: "A description of this activity log alert.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActivityLogAlertResource.Properties.Description")},
			{
				Name:        "location",
				Description: "The location of the resource. Since Azure Activity Log Alerts is a global service, the location of the rules should always be 'global'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActivityLogAlertResource.Location")},
			{
				Name:        "scopes",
				Description: "A list of resourceIds that will be used as prefixes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ActivityLogAlertResource.Properties.Scopes")},
			{
				Name:        "condition",
				Description: "The condition that will cause this alert to activate.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ActivityLogAlertResource.Properties.Condition")},
			{
				Name:        "actions",
				Description: "The actions that will activate when the condition is met.",
				Type:        proto.ColumnType_STRING,

				// Steampipe standard columns
				Transform: transform.FromField("Description.ActivityLogAlertResource.Properties.Actions")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActivityLogAlertResource.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ActivityLogAlertResource.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.ActivityLogAlertResource.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ActivityLogAlertResource.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.ResourceGroup")},
		}),
	}
}

//// HYDRATE FUNCTIONS
