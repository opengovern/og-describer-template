package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSynapseWorkspace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_synapse_workspace",
		Description: "Azure Synapse Workspace",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetSynapseWorkspace,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSynapseWorkspace,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Name")},
			{
				Name:        "id",
				Description: "Fully qualified resource ID for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.ID")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Type")},
			{
				Name:        "adla_resource_id",
				Description: "The ADLA resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.AdlaResourceID")},
			{
				Name:        "managed_resource_group_name",
				Description: "The managed resource group of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.ManagedResourceGroupName")},
			{
				Name:        "managed_virtual_network",
				Description: "A managed virtual network for the workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.ManagedVirtualNetwork")},
			{
				Name:        "public_network_access",
				Description: "Pubic network access to workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.PublicNetworkAccess")},
			{
				Name:        "sql_administrator_login",
				Description: "Login for workspace SQL active directory administrator.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.SQLAdministratorLogin")},
			{
				Name:        "sql_administrator_login_password",
				Description: "The SQL administrator login password of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Properties.SQLAdministratorLoginPassword")},
			{
				Name:        "connectivity_endpoints",
				Description: "Connectivity endpoints of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.ConnectivityEndpoints")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "default_data_lake_storage",
				Description: "Workspace default data lake storage account details.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.DefaultDataLakeStorage")},
			{
				Name:        "encryption",
				Description: "The encryption details of the workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractSynapseWorkspaceEncryption),
			},
			{
				Name:        "extra_properties",
				Description: "Workspace level configs and feature flags.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.ExtraProperties")},
			{
				Name:        "identity",
				Description: "The identity of the workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Identity")},
			{
				Name:        "managed_virtual_network_settings",
				Description: "Managed virtual network settings of the workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.ManagedVirtualNetworkSettings")},
			{
				Name:        "private_endpoint_connections",
				Description: "Private endpoint connections to the workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractSynapseWorkspacePrivateEndpointConnections),
			},
			{
				Name:        "purview_configuration",
				Description: "Purview configuration of the workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.PurviewConfiguration")},
			{
				Name:        "virtual_network_profile",
				Description: "Virtual network profile of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.VirtualNetworkProfile")},
			{
				Name:        "workspace_managed_sql_server_vulnerability_assessments",
				Description: "The vulnerability assessments details of the workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServerVulnerabilityAssessments")},
			{
				Name:        "workspace_repository_configuration",
				Description: "Git integration settings of the workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Properties.WorkspaceRepositoryConfiguration")},
			{
				Name:        "workspace_uid",
				Description: "The unique identifier of the workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:
				// Steampipe standard columns
				transform.FromField("Description.Workspace.Properties.WorkspaceUID")},

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
				Transform:   transform.FromField("ResourceID").Transform(idToAkas),
			},

			// Azure standard columns
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

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

type SynapseWorkspaceEncryption struct {
	DoubleEncryptionEnabled *bool
	CmkStatus               *string
	CmkKey                  interface{}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

// Handle empty name or resourceGroup

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

// Create session

// If we return the API response directly, the output does not provide all
// the contents of DiagnosticSettings

//// TRANSFORM FUNCTIONS

// If we return the API response directly, the output will not provide all the properties of PrivateEndpointConnections
func extractSynapseWorkspacePrivateEndpointConnections(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	workspace := d.HydrateItem.(opengovernance.SynapseWorkspace).Description.Workspace
	var properties []map[string]interface{}

	if workspace.Properties.PrivateEndpointConnections != nil {
		for _, i := range workspace.Properties.PrivateEndpointConnections {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.ID != nil {
				objectMap["name"] = i.Name
			}
			if i.ID != nil {
				objectMap["type"] = i.Type
			}
			if i.Properties != nil {
				if i.Properties.PrivateEndpoint != nil {
					objectMap["privateEndpointPropertyId"] = i.Properties.PrivateEndpoint.ID
				}
				if i.Properties.PrivateLinkServiceConnectionState != nil {
					if i.Properties.PrivateLinkServiceConnectionState.ActionsRequired != nil {
						objectMap["privateLinkServiceConnectionStateActionsRequired"] = i.Properties.PrivateLinkServiceConnectionState.ActionsRequired
					}
					if i.Properties.PrivateLinkServiceConnectionState.Status != nil {
						objectMap["privateLinkServiceConnectionStateStatus"] = i.Properties.PrivateLinkServiceConnectionState.Status
					}
					if i.Properties.PrivateLinkServiceConnectionState.Description != nil {
						objectMap["privateLinkServiceConnectionStateDescription"] = i.Properties.PrivateLinkServiceConnectionState.Description
					}
				}
				if i.Properties.ProvisioningState != nil {
					objectMap["provisioningState"] = i.Properties.ProvisioningState
				}
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of Encryption
func extractSynapseWorkspaceEncryption(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	workspace := d.HydrateItem.(opengovernance.SynapseWorkspace).Description.Workspace
	var properties SynapseWorkspaceEncryption

	if workspace.Properties.Encryption != nil {
		if workspace.Properties.Encryption.DoubleEncryptionEnabled != nil {
			properties.DoubleEncryptionEnabled = workspace.Properties.Encryption.DoubleEncryptionEnabled
		}
		if workspace.Properties.Encryption.Cmk != nil {
			if workspace.Properties.Encryption.Cmk.Status != nil {
				properties.CmkStatus = workspace.Properties.Encryption.Cmk.Status
			}
			if workspace.Properties.Encryption.Cmk.Key != nil {
				properties.CmkKey = workspace.Properties.Encryption.Cmk.Key
			}
		}
	}

	return properties, nil
}
