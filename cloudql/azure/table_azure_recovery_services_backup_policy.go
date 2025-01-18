package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION ////

func tableAzureRecoveryServicesBackupPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_recovery_services_backup_policy",
		Description: "Azure Recovery Services Backup Policy",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRecoveryServicesBackupPolicy,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "404"}),
			},
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "vault_name",
					Require: plugin.Optional,
				},
				{
					Name:    "resource_group",
					Require: plugin.Optional,
				},
			},
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the table service",
				Transform:   transform.FromField("Description.Policy.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a table service uniquely",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.ID"),
			},
			{
				Name:        "vault_name",
				Description: "Backup policy vault name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VaultName"),
			},
			{
				Name:        "properties",
				Description: "Backup policy properties",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Properties"),
			},
			// Azure standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Policy.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Policy.ID").Transform(idToAkas),
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
