package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureOperationalInsightsWorkspaces(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_operationalinsights_workspaces",
		Description: "Azure OperationalInsights Workspaces",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetOperationalInsightsWorkspaces,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListOperationalInsightsWorkspaces,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the workspaces.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspaces.ID")},
			{
				Name:        "name",
				Description: "The name of the workspaces.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Name")},
			{
				Name:        "location",
				Description: "The location of the Log Analytics workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Location"),
			},
			{
				Name:        "type",
				Description: "The type of the Log Analytics workspace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Type"),
			},
			{
				Name:        "sku",
				Description: "The SKU (pricing level) of the Log Analytics workspace.",
				Transform:   transform.FromField("Description.Workspace.Properties.SKU"),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "retention_in_days",
				Description: "The retention period for the Log Analytics workspace data in days.",
				Transform:   transform.FromField("Description.Workspace.Properties.RetentionInDays"),
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the Log Analytics workspace.",
				Transform:   transform.FromField("Description.Workspace.Properties.ProvisioningState"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "workspace_capping",
				Description: "The workspace capping properties.",
				Transform:   transform.FromField("Description.Workspace.Properties.WorkspaceCapping"),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "created_date",
				Description: "Workspace creation date.",
				Transform:   transform.FromField("Description.Workspace.Properties.CreatedDate").Transform(transform.NullIfZeroValue),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "modified_date",
				Description: "Workspace modification date.",
				Transform:   transform.FromField("Description.Workspace.Properties.ModifiedDate").Transform(transform.NullIfZeroValue),
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "customer_id",
				Description: "Represents the ID associated with the workspace.",
				Transform:   transform.FromField("WDescription.Workspace.Properties.CustomerID"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "public_network_access_for_ingestion",
				Description: "The network access type for accessing Log Analytics ingestion.",
				Transform:   transform.FromField("Description.Workspace.Properties.PublicNetworkAccessForIngestion"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "public_network_access_for_query",
				Description: "The network access type for accessing Log Analytics query.",
				Transform:   transform.FromField("Description.Workspace.Properties.PublicNetworkAccessForQuery"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "force_cmk_for_query",
				Description: "Indicates whether customer managed storage is mandatory for query management.",
				Transform:   transform.FromField("Description.Workspace.Properties.ForceCmkForQuery"),
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "private_link_scoped_resources",
				Description: "List of linked private link scope resources.",
				Transform:   transform.FromField("Description.Workspace.Properties.PrivateLinkScopedResources"),
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "enable_data_export",
				Description: "Flag that indicates if data should be exported.",
				Transform:   transform.FromField("WDescription.Workspace.Properties.Features.EnableDataExport"),
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "immediate_purge_data_on_30_days",
				Description: "Flag that describes if we want to remove the data after 30 days.",
				Transform:   transform.FromField("Description.Workspace.Properties.Features.ImmediatePurgeDataOn30Days"),
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "enable_log_access_using_only_resource_permissions",
				Description: "Flag that indicates which permission to use - resource or workspace or both.",
				Transform:   transform.FromField("WDescription.Workspace.Properties.Features.EnableLogAccessUsingOnlyResourcePermissions"),
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "cluster_resource_id",
				Description: "Dedicated LA cluster resourceId that is linked to the workspaces.",
				Transform:   transform.FromField("Description.Workspace.Properties.Features.ClusterResourceID"),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "disable_local_auth",
				Description: "Disable Non-AAD based Auth.",
				Transform:   transform.FromField("Description.Workspace..Properties.Features.DisableLocalAuth"),
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "tags",
				Description: "The tags assigned to the Log Analytics workspace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workspace.Tags"),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workspace.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.Workspace.Properties.WorkspaceCapping").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: "The region of the Log Analytics workspace.",
				Transform:   transform.FromField("Description.Workspace.Location").Transform(toLower),
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource_group",
				Description: "The resource group of the Log Analytics workspace.",
				Transform:   transform.FromField("Description.ResourceGroup"),
				Type:        proto.ColumnType_STRING,
			},
		}),
	}
}
