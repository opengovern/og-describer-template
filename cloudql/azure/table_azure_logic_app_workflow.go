package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureLogicAppWorkflow(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_logic_app_workflow",
		Description: "Azure Logic App Workflow",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetLogicAppWorkflow,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "400"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLogicAppWorkflow,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workflow.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workflow.ID")},
			{
				Name:        "state",
				Description: "The state of the workflow.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workflow.Properties.State")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the workflow.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workflow.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workflow.Type")},
			{
				Name:        "access_endpoint",
				Description: "The access endpoint of the workflow.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workflow.Properties.AccessEndpoint")},
			{
				Name:        "created_time",
				Description: "The time when workflow was created.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Workflow.Properties.CreatedTime").Transform(convertDateToTime),
			},
			{
				Name:        "changed_time",
				Description: "Specifies the time, the workflow was updated.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Workflow.Properties.ChangedTime").Transform(convertDateToTime),
			},
			{
				Name:        "sku_name",
				Description: "The sku name.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Workflow.Properties.SKU.Name"),
			},
			{
				Name:        "version",
				Description: "Version of the workflow.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workflow.Properties.Version")},
			{
				Name:        "access_control",
				Description: "The access control configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workflow.Properties.AccessControl")},
			{
				Name:        "definition",
				Description: "The workflow defination.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workflow.Properties.Definition")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the workflow.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "endpoints_configuration",
				Description: "The endpoints configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workflow.Properties.EndpointsConfiguration")},
			{
				Name:        "integration_account",
				Description: "The integration account of the workflow.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workflow.Properties.IntegrationAccount")},
			{
				Name:        "integration_service_environment",
				Description: "The integration service environment of the workflow.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workflow.Properties.IntegrationServiceEnvironment")},
			{
				Name:        "parameters",
				Description: "The workflow parameters.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workflow.Properties.Parameters")},
			{
				Name:        "sku_plan",
				Description: "The sku Plan.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Workflow.Properties.SKU.Plan")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workflow.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workflow.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Workflow.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Workflow.Location").Transform(toLower),
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

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS

// Create session

// Return nil, if no input provide

// Create session

// If we return the API response directly, the output only gives
// the contents of DiagnosticSettings
