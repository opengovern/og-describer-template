package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureMachineLearningWorkspace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_machine_learning_workspace",
		Description: "Azure Machine Learning Workspace",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetMachineLearningWorkspace,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "Invalid input"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMachineLearningWorkspace,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Name")},
			{
				Name:        "friendly_name",
				Description: "The friendly name for this workspace. This name in mutable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.FriendlyName")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.ID")},
			{
				Name:        "provisioning_state",
				Description: "The current deployment state of workspace resource, The provisioningState is to indicate states for resource provisioning. Possible values include: 'Unknown', 'Updating', 'Creating', 'Deleting', 'Succeeded', 'Failed', 'Canceled'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.ProvisioningState")},
			{
				Name:        "creation_time",
				Description: "The creation time for this workspace resource.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Workspace.SystemData.CreatedAt").Transform(convertDateToTime),
			},
			{
				Name:        "workspace_id",
				Description: "The immutable id associated with this workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.WorkspaceID")},
			{
				Name:        "application_insights",
				Description: "ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.ApplicationInsights")},
			{
				Name:        "container_registry",
				Description: "ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.ContainerRegistry")},
			{
				Name:        "description",
				Description: "The description of this workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.Description")},
			{
				Name:        "discovery_url",
				Description: "ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.DiscoveryURL")},
			{
				Name:        "hbi_workspace",
				Description: "The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Workspace.Properties.HbiWorkspace")},
			{
				Name:        "key_vault",
				Description: "ARM id of the key vault associated with this workspace, This cannot be changed once the workspace has been created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.KeyVault")},
			{
				Name:        "location",
				Description: "The location of the resource. This cannot be changed after the resource is created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Location")},
			{
				Name:        "service_provisioned_resource_group",
				Description: "The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.ServiceProvisionedResourceGroup")},
			{
				Name:        "sku_name",
				Description: "Name of the sku.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.SKU.Name")},
			{
				Name:        "sku_tier",
				Description: "Tier of the sku like Basic or Enterprise.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.SKU.Tier")},
			{
				Name:        "storage_account",
				Description: "ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.StorageAccount")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Type")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the azure ML workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "encryption",
				Description: "The encryption settings of Azure ML workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.Encryption")},
			{
				Name:        "identity",
				Description: "The identity of the resource.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Workspace.Identity")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Workspace.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.

					//// HYDRATE FUNCTIONS
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Create session

// Return nil, if no input provide

// Create session

// If we return the API response directly, the output only gives
// the contents of DiagnosticSettings
