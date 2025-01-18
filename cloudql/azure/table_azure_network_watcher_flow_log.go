package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureNetworkWatcherFlowLog(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_network_watcher_flow_log",
		Description: "Azure Network Watcher FlowLog",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"network_watcher_name", "name", "resource_group"}),
			Hydrate:    opengovernance.GetNetworkWatcherFlowLog,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListNetworkWatcherFlowLog,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the flow log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a flow log uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.ID")},
			{
				Name:        "enabled",
				Description: "Indicates whether the flow log is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.FlowLog.Properties.Enabled")},
			{
				Name:        "network_watcher_name",
				Description: "The friendly name that identifies the network watcher.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NetworkWatcherName"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the flow log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.Properties.ProvisioningState"),
			},
			{
				Name:        "type",
				Description: "The resource type of the flow log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.Type")},
			{
				Name:        "version",
				Description: "The version (revision) of the flow log.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FlowLog.Properties.Format.Version")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.Etag")},
			{
				Name:        "file_type",
				Description: "The file type of flow log. Possible values include: 'JSON'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.FlowLog.Properties.Format.Type"),
			},
			{
				Name:        "retention_policy_days",
				Description: "Specifies the number of days to retain flow log records.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FlowLog.Properties.RetentionPolicy.Days")},
			{
				Name:        "retention_policy_enabled",
				Description: "Indicates whether flow log retention is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.FlowLog.Properties.RetentionPolicy.Enabled")},
			{
				Name:        "storage_id",
				Description: "The ID of the storage account which is used to store the flow log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.Properties.StorageID")},
			{
				Name:        "target_resource_id",
				Description: "The ID of network security group to which flow log will be applied.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.Properties.TargetResourceID")},
			{
				Name:        "target_resource_guid",
				Description: "The Guid of network security group to which flow log will be applied.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.Properties.TargetResourceGUID")},
			{
				Name:        "traffic_analytics",
				Description: "Defines the configuration of flow log traffic analytics.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.FlowLog.Properties.FlowAnalyticsConfiguration.NetworkWatcherFlowAnalyticsConfiguration")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FlowLog.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.FlowLog.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.FlowLog.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTIONS

				Transform: transform.

					// Get the details of network watcher
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Create session

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS
