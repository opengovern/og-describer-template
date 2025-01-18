package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureLogProfile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_log_profile",
		Description: "Azure Log Profile",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetLogProfile,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLogProfile,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the resource.",
				Transform:   transform.FromField("Description.LogProfileResource.Name")},
			{
				Name:        "id",
				Description: "The resource Id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfileResource.ID")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfileResource.Type")},
			{
				Name:        "location",
				Description: "Specifies the name of the region, the resource is created at.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfileResource.Location")},
			{
				Name:        "storage_account_id",
				Description: "The resource id of the storage account to which you would like to send the Activity Log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfileResource.Properties.StorageAccountID")},
			{
				Name:        "service_bus_rule_id",
				Description: "The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfileResource.Properties.ServiceBusRuleID")},
			{
				Name:        "log_event_location",
				Description: "List of regions for which Activity Log events should be stored or streamed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LogProfileResource.Properties.Locations")},
			{
				Name:        "categories",
				Description: "The categories of the logs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LogProfileResource.Properties.Categories")},
			{
				Name:        "retention_policy",
				Description: "The retention policy for the events in the log.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.LogProfileResource.Properties.RetentionPolicy")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LogProfileResource.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LogProfileResource.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.LogProfileResource.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LogProfileResource.Location").Transform(toLower),
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
