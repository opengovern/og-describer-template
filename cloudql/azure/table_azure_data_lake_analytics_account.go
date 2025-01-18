package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureDataLakeAnalyticsAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_data_lake_analytics_account",
		Description: "Azure Data Lake Analytics account",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetDataLakeAnalyticsAccount,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "400"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDataLakeAnalyticsAccount,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.ID")},
			{
				Name:        "state",
				Description: "The state of the data lake analytics account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.State")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning status of the data lake analytics account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Type")},
			{
				Name:        "account_id",
				Description: "The unique identifier associated with this data lake analytics account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.AccountID")},
			{
				Name:        "creation_time",
				Description: "The data lake analytics account creation time.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.DataLakeAnalyticsAccount.Properties.CreationTime").Transform(convertDateToTime),
			},
			{
				Name:        "current_tier",
				Description: "The commitment tier in use for current month.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.CurrentTier")},
			{
				Name:        "default_data_lake_store_account",
				Description: "The default data lake store account associated with this data lake analytics account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.DefaultDataLakeStoreAccount")},
			{
				Name:        "endpoint",
				Description: "The full cname endpoint for this data lake analytics account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.Endpoint")},
			{
				Name:        "firewall_state",
				Description: "The current state of the IP address firewall for this data lake analytics account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.FirewallState")},
			{
				Name:        "last_modified_time",
				Description: "The data lake analytics account last modified time.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.DataLakeAnalyticsAccount.Properties.LastModifiedTime").Transform(convertDateToTime),
			},
			{
				Name:        "max_degree_of_parallelism",
				Description: "The maximum supported degree of parallelism for this data lake analytics account.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.MaxDegreeOfParallelism")},
			{
				Name:        "max_degree_of_parallelism_per_job",
				Description: "The maximum supported degree of parallelism per job for this data lake analytics account.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.MaxDegreeOfParallelismPerJob")},
			{
				Name:        "max_job_count",
				Description: "The maximum supported jobs running under the data lake analytics account at the same time.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.MaxJobCount")},
			{
				Name:        "min_priority_per_job",
				Description: "The minimum supported priority per job for this data lake analytics account.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.MinPriorityPerJob")},
			{
				Name:        "new_tier",
				Description: "The commitment tier to use for next month.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.NewTier")},
			{
				Name:        "query_store_retention",
				Description: "The number of days that job metadata is retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.QueryStoreRetention")},
			{
				Name:        "system_max_degree_of_parallelism",
				Description: "The system defined maximum supported degree of parallelism for this account, which restricts the maximum value of parallelism the user can set for the data lake analytics account.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.SystemMaxDegreeOfParallelism")},
			{
				Name:        "system_max_job_count",
				Description: "The system defined maximum supported jobs running under the account at the same time, which restricts the maximum number of running jobs the user can set for the data lake analytics account.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.SystemMaxJobCount")},
			{
				Name:        "compute_policies",
				Description: "The list of compute policies associated with this data lake analytics account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.ComputePolicies")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the data lake analytics account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResource")},
			{
				Name:        "data_lake_store_accounts",
				Description: "The list of data lake store accounts associated with this data lake analytics account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.DataLakeStoreAccounts")},
			{
				Name:        "firewall_allow_azure_ips",
				Description: "The current state of allowing or disallowing IPs originating within azure through the firewall. If the firewall is disabled, this is not enforced.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.FirewallAllowAzureIPs")},
			{
				Name:        "firewall_rules",
				Description: "The list of firewall rules associated with this data lake analytics account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Properties.FirewallRules")},
			{
				Name:        "storage_accounts",
				Description: "The list of azure blob storage accounts associated with this data lake analytics account.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.DataLakeAnalyticsAccount.Properties.StorageAccounts")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DataLakeAnalyticsAccount.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.DataLakeAnalyticsAccount.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DataLakeAnalyticsAccount.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
