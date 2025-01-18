package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureMonitorLogProfile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_monitor_log_profile",
		Description: "Azure Monitor Log Profile",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name"}),
			Hydrate:    opengovernance.GetMonitorLogProfile,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMonitorLogProfile,
		},
		Columns: azureColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Azure resource Id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfile.ID"),
			},
			{
				Name:        "name",
				Description: "Azure resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfile.Name"),
			},
			{
				Name:        "type",
				Description: "Azure resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfile.Type"),
			},
			{
				Name:        "location",
				Description: "The resource location.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfile.Location"),
			},
			{
				Name:        "storage_account_id",
				Description: "The resource id of the storage account to which you would like to send the Activity Log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfile.Properties.StorageAccountID"),
			},
			{
				Name:        "service_bus_rule_id",
				Description: "The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfile.Properties.ServiceBusRuleID"),
			},
			{
				Name:        "locations",
				Description: "List of regions for which Activity Log events should be stored or streamed. It is a comma separated list of valid ARM locations including the 'global' location.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LogProfile.Properties.Locations"),
			},
			{
				Name:        "categories",
				Description: "The categories of the logs. These categories are created as is convenient to the user.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LogProfile.Properties.Categories"),
			},
			{
				Name:        "retention_policy",
				Description: "The retention policy for the events in the log.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LogProfile.Properties.RetentionPolicy"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfile.Name"),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LogProfile.Tags"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfile.ID").Transform(idToAkas),
			},
		}),
	}
}
