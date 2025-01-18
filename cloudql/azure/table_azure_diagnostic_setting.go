package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureDiagnosticSetting(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_diagnostic_setting",
		Description: "Azure Diagnostic Setting",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetDiagnosticSetting,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDiagnosticSetting,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the resource.",
				Transform:   transform.FromField("Description.DiagnosticSettingsResource.Name")},
			{
				Name:        "id",
				Description: "The resource Id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource.ID")},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource.Type")},
			{
				Name:        "storage_account_id",
				Description: "The resource ID of the storage account to which you would like to send Diagnostic Logs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource.Properties.StorageAccountID")},
			{
				Name:        "service_bus_rule_id",
				Description: "The service bus rule Id of the diagnostic setting.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource.Properties.ServiceBusRuleID")},
			{
				Name:        "event_hub_authorization_rule_id",
				Description: "The resource Id for the event hub authorization rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource.Properties.EventHubAuthorizationRuleID")},
			{
				Name:        "event_hub_name",
				Description: "The name of the event hub.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource.Properties.EventHubName")},
			{
				Name:        "workspace_id",
				Description: "The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource.Properties.WorkspaceID")},
			{
				Name:        "log_analytics_destination_type",
				Description: "A string indicating whether the export to Log Analytics should use the default destinatio type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource.Properties.LogAnalyticsDestinationType")},
			{
				Name:        "metrics",
				Description: "The list of metric settings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource.Properties.Metrics")},
			{
				Name:        "logs",
				Description: "The list of logs settings.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.DiagnosticSettingsResource.Properties.Logs")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.DiagnosticSettingsResource.ID").Transform(idToAkas),
			},

			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup")},
		}),
	}
}
