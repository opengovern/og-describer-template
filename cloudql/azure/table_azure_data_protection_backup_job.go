package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION ////

func tableAzureDataProtectionBackupJob(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_data_protection_backup_job",
		Description: "Azure Data Protection Backup Job",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDataProtectionJob,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "404"}),
			},
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Resource name associated with the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Name"),
			},
			{
				Name:        "vault_name",
				Description: "The data protection vault name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VaultName"),
			},
			{
				Name:        "id",
				Description: "Resource ID represents the complete path to the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.ID"),
			},
			{
				Name:        "type",
				Description: "Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Type"),
			},
			{
				Name:        "activity_id",
				Description: "Job Activity Id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.ActivityID"),
			},
			{
				Name:        "backup_instance_friendly_name",
				Description: "Name of the Backup Instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.BackupInstanceFriendlyName"),
			},
			{
				Name:        "data_source_id",
				Description: "ARM ID of the DataSource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.DataSourceID"),
			},
			{
				Name:        "data_source_location",
				Description: "Location of the DataSource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.DataSourceLocation"),
			},
			{
				Name:        "data_source_name",
				Description: "User Friendly Name of the DataSource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.DataSourceName"),
			},
			{
				Name:        "data_source_type",
				Description: "Type of DataSource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.DataSourceType"),
			},
			{
				Name:        "is_user_triggered",
				Description: "Indicates whether the job is adhoc(true) or scheduled(false).",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.IsUserTriggered"),
			},
			{
				Name:        "operation",
				Description: "Type of Job i.e. Backup:full/log/diff ;Restore:ALR/OLR; Tiering:Backup/Archive ; Management:ConfigureProtection/UnConfigure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.Operation"),
			},
			{
				Name:        "operation_category",
				Description: "Indicates the type of Job i.e. Backup/Restore/Tiering/Management.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.OperationCategory"),
			},
			{
				Name:        "progress_enabled",
				Description: "Indicates whether progress is enabled for the job.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.ProgressEnabled"),
			},
			{
				Name:        "source_resource_group",
				Description: "Resource Group Name of the Datasource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.SourceResourceGroup"),
			},
			{
				Name:        "source_subscription_id",
				Description: "SubscriptionId corresponding to the DataSource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.SourceSubscriptionID"),
			},
			{
				Name:        "start_time",
				Description: "StartTime of the job (in UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.StartTime"),
			},
			{
				Name:        "status",
				Description: "Status of the job like InProgress/Success/Failed/Cancelled/SuccessWithWarning.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.Status"),
			},
			{
				Name:        "data_source_set_name",
				Description: "Data Source Set Name of the DataSource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.DataSourceSetName"),
			},
			{
				Name:        "destination_data_store_name",
				Description: "Destination Data Store Name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.DestinationDataStoreName"),
			},
			{
				Name:        "duration",
				Description: "Total run time of the job. ISO 8601 format.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.Duration"),
			},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.Etag"),
			},
			{
				Name:        "source_data_store_name",
				Description: "Source Data Store Name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.SourceDataStoreName"),
			},
			{
				Name:        "backup_instance_id",
				Description: "ARM ID of the Backup Instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.BackupInstanceID"),
			},
			{
				Name:        "end_time",
				Description: "EndTime of the job (in UTC).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.EndTime").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "policy_id",
				Description: "ARM ID of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.PolicyID"),
			},
			{
				Name:        "policy_name",
				Description: "Name of the policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.PolicyName"),
			},
			{
				Name:        "progress_url",
				Description: "Url which contains job's progress.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.ProgressURL"),
			},
			{
				Name:        "restore_type",
				Description: "Indicates the sub type of operation i.e. in case of Restore it can be ALR/OLR.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.RestoreType"),
			},
			{
				Name:        "error_details",
				Description: "A List, detailing the errors related to the job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.ErrorDetails"),
			},
			{
				Name:        "supported_actions",
				Description: "List of supported actions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.SupportedActions"),
			},
			{
				Name:        "extended_info",
				Description: "Extended Information about the job.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataProtectionJob.Properties.ExtendedInfo"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataProtectionJob.Name"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataProtectionJob.ID").Transform(idToAkas),
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
