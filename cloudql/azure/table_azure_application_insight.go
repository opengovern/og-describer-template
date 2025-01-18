package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureApplicationInsight(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_application_insight",
		Description: "Azure Application Insight",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetApplicationInsightsComponent,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListApplicationInsightsComponent,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the application insight.",
				Transform:   transform.FromField("Description.Component.Name"),
			},
			{
				Name:        "id",
				Description: "Contains id to identify the application insight uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.ID"),
			},
			{
				Name:        "app_id",
				Description: "Application insights unique id for your Application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.Properties.AppID"),
			},
			{
				Name:        "connection_string",
				Description: "Application Insights component connection string.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.Properties.ConnectionString"),
			},
			{
				Name:        "creation_date",
				Description: "Creation date for the Application Insights component.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Component.Properties.CreationDate"),
			},
			{
				Name:        "disable_ip_masking",
				Description: "Disable IP masking.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Component.Properties.DisableIPMasking"),
			},
			{
				Name:        "disable_local_auth",
				Description: "Disable Non-AAD based Auth.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Component.Properties.DisableLocalAuth"),
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.Etag"),
			},
			{
				Name:        "force_customer_storage_for_profiler",
				Description: "Force users to create their own storage account for profiler and debugger.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Component.Properties.ForceCustomerStorageForProfiler"),
			},
			{
				Name:        "immediate_purge_data_on_30_days",
				Description: "Purge data immediately after 30 days.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Component.Properties.ImmediatePurgeDataOn30Days"),
			},
			{
				Name:        "instrumentation_key",
				Description: "Application Insights Instrumentation key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.Properties.InstrumentationKey"),
			},
			{
				Name:        "kind",
				Description: "The kind of application that this component refers to, used to customize UI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.Kind"),
			},
			{
				Name:        "provisioning_state",
				Description: "Current state of this component.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.Properties.ProvisioningState"),
			},
			{
				Name:        "request_source",
				Description: "Describes what tool created this Application Insights component.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.Properties.RequestSource"),
			},
			{
				Name:        "retention_in_days",
				Description: "Retention period in days.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Component.Properties.RetentionInDays"),
			},
			{
				Name:        "sampling_percentage",
				Description: "Percentage of the data produced by the application being monitored that is being sampled for Application Insights telemetry.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.Component.Properties.SamplingPercentage"),
			},
			{
				Name:        "tenant_id",
				Description: "Azure Tenant ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.Properties.TenantID"),
			},
			{
				Name:        "type",
				Description: "The resource type of the application insight.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.Type"),
			},
			{
				Name:        "workspace_resource_id",
				Description: "Resource Id of the log analytics workspace to which the data will be ingested.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.Properties.WorkspaceResourceID"),
			},
			{
				Name:        "application_type",
				Description: "Type of application being monitored.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Component.Properties.ApplicationType"),
			},
			{
				Name:        "flow_type",
				Description: "Determines what kind of flow this component was created by.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Component.Properties.FlowType"),
			},
			{
				Name:        "ingestion_mode",
				Description: "Indicates the flow of the ingestion.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Component.Properties.IngestionMode"),
			},
			{
				Name:        "private_link_scoped_resources",
				Description: "List of linked private link scope resources.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Component.Properties.PrivateLinkScopedResources"),
			},
			{
				Name:        "public_network_access_for_ingestion",
				Description: "The network access type for accessing Application Insights ingestion.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Component.Properties.PublicNetworkAccessForIngestion"),
			},
			{
				Name:        "public_network_access_for_query",
				Description: "The network access type for accessing Application Insights query.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Component.Properties.PublicNetworkAccessForQuery"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Component.Name"),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Component.Tags"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Transform:   transform.FromField("Description.Component.ID").Transform(idToAkas),
				Type:        proto.ColumnType_JSON,
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Transform:   transform.FromField("Description.Component.Location").Transform(toLower),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Transform:   transform.FromField("Description.ResourceGroup"),
				Type:        proto.ColumnType_STRING,
			},
		}),
	}
}
