package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION ////

func tableAzureRecoveryServicesBackupJob(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_recovery_services_backup_job",
		Description: "Azure Recovery Services Backup Job",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRecoveryServicesBackupJob,
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
				Description: "Resource name associated with the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.Name")},
			{
				Name:        "vault_name",
				Description: "The recovery vault name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VaultName")},
			{
				Name:        "id",
				Description: "Resource ID represents the complete path to the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.ID")},
			{
				Name:        "type",
				Description: "Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.Type")},
			{
				Name:        "etag",
				Description: "Optional ETag.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.ETag")},
			// JSON fields
			{
				Name:        "properties",
				Description: "JobResource properties.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Properties")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.Name")},

			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Job.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Job.Location").Transform(toLower),
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
