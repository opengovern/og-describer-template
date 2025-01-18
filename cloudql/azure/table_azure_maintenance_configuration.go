package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureMaintenanceConfiguration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_maintenance_configuration",
		Description: "Azure Maintenance Configuration.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"resource_group", "name"}),
			Hydrate:    opengovernance.GetMaintenanceConfiguration,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMaintenanceConfiguration,
		},
		Columns: azureColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Fully qualified identifier of the resource.",
				Transform:   transform.FromField("Description.MaintenanceConfiguration.ID"),
			},
			{
				Name:        "name",
				Description: "Name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.Name"),
			},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.Type"),
			},
			{
				Name:        "namespace",
				Description: "Gets or sets namespace of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.Properties.Namespace"),
			},
			{
				Name:        "visibility",
				Description: "The visibility of the configuration. The default value is 'Custom'. Possible values include: 'VisibilityCustom', 'VisibilityPublic'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.Properties.Visibility"),
			},
			{
				Name:        "maintenance_scope",
				Description: "The maintenanceScope of the configuration. Possible values include: 'ScopeHost', 'ScopeOSImage', 'ScopeExtension', 'ScopeInGuestPatch', 'ScopeSQLDB', 'ScopeSQLManagedInstance'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.Properties.MaintenanceScope"),
			},
			{
				Name:        "created_at",
				Description: "The timestamp of resource creation (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.SystemData.CreatedAt").Transform(convertDateToTime),
			},
			{
				Name:        "created_by",
				Description: "The identity that created the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.SystemData.CreatedBy"),
			},
			{
				Name:        "created_by_type",
				Description: "The type of identity that created the resource. Possible values include: 'CreatedByTypeUser', 'CreatedByTypeApplication', 'CreatedByTypeManagedIdentity', 'CreatedByTypeKey'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.SystemData.CreatedByType"),
			},
			{
				Name:        "last_modified_at",
				Description: "The timestamp of resource last modification (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.SystemData.LastModifiedAt").Transform(convertDateToTime),
			},
			{
				Name:        "last_modified_by",
				Description: "The identity that last modified the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.SystemData.LastModifiedBy"),
			},
			{
				Name:        "last_modified_by_type",
				Description: "The type of identity that last modified the resource. Possible values include: 'CreatedByTypeUser', 'CreatedByTypeApplication', 'CreatedByTypeManagedIdentity', 'CreatedByTypeKey'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.SystemData.LastModifiedByType"),
			},
			{
				Name:        "extension_properties",
				Description: "Gets or sets extensionProperties of the maintenanceConfiguration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("CDescription.MaintenanceConfiguration.Properties.ExtensionProperties"),
			},
			{
				Name:        "window",
				Description: "Definition of a MaintenanceWindow.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.Properties.Window"),
			},
			{
				Name:        "system_data",
				Description: "Azure Resource Manager metadata containing createdBy and modifiedBy information.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.SystemData"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.Name"),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.Tags"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.MaintenanceConfiguration.Location").Transform(toLower),
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
