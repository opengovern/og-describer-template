package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureDataProtectionBackupVaults(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_dataprotection_backupvaults",
		Description: "Azure DataProtection BackupVaults",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetDataProtectionBackupVaults,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDataProtectionBackupVaults,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the backupvaults.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVaults.ID")},
			{
				Name:        "name",
				Description: "The name of the backupvaults.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVaults.Name")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVaults.Type"),
			},
			{
				Name:        "location",
				Description: "The location of the backup vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVaults.Location").Transform(toLower),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the backup vault resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVaults.Properties.ProvisioningState"),
			},
			{
				Name:        "resource_move_state",
				Description: "The resource move state for the backup vault.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVaults.Properties.ResourceMoveState"),
			},
			{
				Name:        "storage_settings",
				Description: "The storage settings of the backup vault.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BackupVaults.Properties.StorageSettings"),
			},
			{
				Name:        "monitoring_settings",
				Description: "The Monitoring Settings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BackupVaults.Properties.MonitoringSettings"),
			},
			{
				Name:        "identity",
				Description: "Input Managed Identity Details.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BackupVaults.Identity"),
			},
			{
				Name:        "system_data",
				Description: "Metadata pertaining to creation and last modification of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BackupVaults.SystemData"),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVaults.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.BackupVaults.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.BackupVaults.ID").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: "The Azure region where the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVaults.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: "The resource group in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupVaults.ID").Transform(extractResourceGroupFromID),
			},
		}),
	}
}
