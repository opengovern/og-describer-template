package azure

import (
	"context"
	"strings"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSqlDatabase(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_sql_database",
		Description: "Azure SQL Database",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "server_name", "resource_group"}),
			Hydrate:    opengovernance.GetSqlDatabase,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSqlDatabase,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the database.",
				Transform:   transform.FromField("Description.Database.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a database uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.ID")},
			{
				Name:        "server_name",
				Description: "The name of the parent server of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(idToServerName),
			},
			{
				Name:        "status",
				Description: "The status of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.Status")},
			{
				Name:        "type",
				Description: "Type of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Type")},
			{
				Name:        "collation",
				Description: "The collation of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.Collation")},
			//{
			//	Name:        "containment_state",
			//	Description: "The containment state of the database.",
			//	Type:        proto.ColumnType_INT,
			//	Transform:   transform.FromField("Description.Database.Properties.ContainmentState")}, // not correct
			{
				Name:        "creation_date",
				Description: "The creation date of the database.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Database.Properties.CreationDate").Transform(convertDateToTime),
			},
			{
				Name:        "current_service_objective_id",
				Description: "The current service level objective ID of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.CurrentServiceObjectiveName")},
			{
				Name:        "database_id",
				Description: "The ID of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.DatabaseID")},
			{
				Name:        "default_secondary_location",
				Description: "The default secondary region for this database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.DefaultSecondaryLocation")},
			{
				Name:        "earliest_restore_date",
				Description: "This records the earliest start date and time that restore is available for this database.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Database.Properties.EarliestRestoreDate").Transform(convertDateToTime),
			},
			{
				Name:        "edition",
				Description: "The edition of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.RequestedServiceObjectiveName")},
			{
				Name:        "elastic_pool_name",
				Description: "The name of the elastic pool the database is in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.ElasticPoolID")},
			{
				Name:        "failover_group_id",
				Description: "The resource identifier of the failover group containing this database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.FailoverGroupID")},
			{
				Name:        "kind",
				Description: "Kind of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Kind")},
			{
				Name:        "location",
				Description: "Location of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Location")},
			{
				Name:        "max_size_bytes",
				Description: "The max size of the database expressed in bytes.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Database.Properties.MaxSizeBytes")},
			{
				Name:        "recovery_services_recovery_point_resource_id",
				Description: "Specifies the resource ID of the recovery point to restore from if createMode is RestoreLongTermRetentionBackup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.RecoveryServicesRecoveryPointID")},
			{
				Name:        "requested_service_objective_id",
				Description: "The configured service level objective ID of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.RequestedServiceObjectiveName")},
			{
				Name:        "restore_point_in_time",
				Description: "Specifies the point in time of the source database that will be restored to create the new database.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Database.Properties.RestorePointInTime").Transform(convertDateToTime),
			},
			{
				Name:        "requested_service_objective_name",
				Description: "The name of the configured service level objective of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.RequestedServiceObjectiveName")},
			{
				Name:        "retention_policy_id",
				Description: "Retention policy ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LongTermRetentionPolicy.ID")},
			{
				Name:        "retention_policy_name",
				Description: "Retention policy Name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LongTermRetentionPolicy.Name")},
			{
				Name:        "retention_policy_type",
				Description: "Long term Retention policy Type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LongTermRetentionPolicy.Type")},
			{
				Name:        "source_database_deletion_date",
				Description: "Specifies the time that the database was deleted when createMode is Restore and sourceDatabaseId is the deleted database's original resource id.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Database.Properties.SourceDatabaseDeletionDate").Transform(convertDateToTime),
			},
			{
				Name:        "source_database_id",
				Description: "Specifies the resource ID of the source database if createMode is Copy, NonReadableSecondary, OnlineSecondary, PointInTimeRestore, Recovery, or Restore.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.SourceDatabaseID")},
			{
				Name:        "zone_redundant",
				Description: "Indicates if the database is zone redundant or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Database.Properties.ZoneRedundant")},
			{
				Name:        "create_mode",
				Description: "Specifies the mode of database creation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.CreateMode")},
			{
				Name:        "read_scale",
				Description: "ReadScale indicates whether read-only connections are allowed to this database or not if the database is a geo-secondary.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.ReadScale")},
			//{
			//	Name:        "recommended_index",
			//	Description: "The recommended indices for this database.", // this could be available through service_tier_advisors
			//	Type:        proto.ColumnType_JSON,
			//	Transform:   transform.FromField("DatabaseProperties.RecommendedIndex"), // not correct
			//},
			{
				Name:        "retention_policy_property",
				Description: "Long term Retention policy Property.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LongTermRetentionPolicy.Properties")},
			{
				Name:        "sample_name",
				Description: "Indicates the name of the sample schema to apply when creating this database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.SampleName")},
			{
				Name:        "service_level_objective",
				Description: "The current service level objective of the database.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Properties.RequestedServiceObjectiveName")},
			{
				Name:        "service_tier_advisors",
				Description: "The list of service tier advisors for this database.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Advisors"), // is it ok to be an array?
			},
			{
				Name:        "transparent_data_encryption",
				Description: "The transparent data encryption info for this database.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TransparentDataEncryption"), // is it ok to be an array?
			},
			{
				Name:        "vulnerability_assessments",
				Description: "The vulnerability assessments for this database.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DatabaseVulnerabilityAssessments")}, // is it ok to be an array?
			{
				Name:        "vulnerability_assessment_scan_records",
				Description: "The vulnerability assessment scan records for this database.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.VulnerabilityAssessmentScanRecords")}, // is it ok to be an array?

			{
				Name:        "audit_policy",
				Description: "The database blob auditing policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AuditPolicies"),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Database.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Database.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Database.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Database.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

func idToServerName(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.SqlDatabase).Description.Database
	if data.ID == nil {
		return nil, nil
	}
	serverName := strings.Split(*data.ID, "/")[8]
	return serverName, nil
}
