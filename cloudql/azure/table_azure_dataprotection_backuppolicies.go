package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureDataProtectionBackupPolicies(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_dataprotection_backuppolicies",
		Description: "Azure DataProtection BackupPolicies",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetDataProtectionBackupVaultsBackupPolicies,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDataProtectionBackupVaultsBackupPolicies,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the backuppolicies.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupPolicies.ID")},
			{
				Name:        "name",
				Description: "The name of the backuppolicies.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupPolicies.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BackupPolicies.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.BackupPolicies.ID").Transform(idToAkas),
			},
		}),
	}
}
