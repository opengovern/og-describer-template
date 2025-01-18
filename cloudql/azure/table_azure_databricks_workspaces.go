package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureDatabricksWorkspaces(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_databricks_workspaces",
		Description: "Azure Databricks Workspaces",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetDatabricksWorkspace,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDatabricksWorkspace,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the workspaces.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.ID")},
			{
				Name:        "name",
				Description: "The name of the workspaces.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Name")},
			{
				Name:        "sku",
				Description: "The SKU of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.SKU"),
			},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Type"),
			},
			{
				Name:        "location",
				Description: "The geo-location where the resource lives.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Location"),
			},
			{
				Name:        "managed_resource_group_id",
				Description: "The managed resource group ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.ManagedResourceGroupID"),
			},
			{
				Name:        "parameters",
				Description: "The workspace's custom parameters.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.Parameters"),
			},
			{
				Name:        "provisioning_state",
				Description: "The workspace provisioning state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.ProvisioningState"),
			},
			{
				Name:        "ui_definition_uri",
				Description: "The blob URI where the UI definition file is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.UIDefinitionURI"),
			},
			{
				Name:        "authorizations",
				Description: "The workspace provider authorizations.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.Authorizations"),
			},
			{
				Name:        "created_by",
				Description: "Indicates the Object ID, PUID and Application ID of entity that created the workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.CreatedBy"),
			},
			{
				Name:        "updated_by",
				Description: "Indicates the Object ID, PUID and Application ID of entity that last updated the workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.UpdatedBy"),
			},
			{
				Name:        "created_date_time",
				Description: "Specifies the date and time when the workspace is created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Workspace.Properties.CreatedDateTime").Transform(convertDateToTime),
			},
			{
				Name:        "workspace_id",
				Description: "The unique identifier of the databricks workspace in databricks control plane.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.WorkspaceID"),
			},
			{
				Name:        "workspace_url",
				Description: "The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.WorkspaceURL"),
			},
			{
				Name:        "storage_account_identity",
				Description: "The details of Managed Identity of Storage Account",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.StorageAccountIdentity"),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.Workspace.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.Workspace.ID").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}
