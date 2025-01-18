package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureSearchService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_search_service",
		Description: "Azure Search Service",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetSearchService,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSearchService,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the resource.",
				Transform:   transform.FromField("Description.Service.Name")},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Fully qualified resource ID for the resource.",
				Transform:   transform.FromField("Description.Service.ID")},
			{
				Name:        "provisioning_state",
				Type:        proto.ColumnType_STRING,
				Description: "The state of the last provisioning operation performed on the search service.",
				Transform:   transform.FromField("Description.Service.Properties.ProvisioningState")},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the search service. Possible values include: 'running', deleting', 'provisioning', 'degraded', 'disabled', 'error' etc.",
				Transform:   transform.FromField("Description.Service.Properties.Status")},
			{
				Name:        "status_details",
				Type:        proto.ColumnType_STRING,
				Description: "The details of the search service status.",
				Transform:   transform.FromField("Description.Service.Properties.StatusDetails")},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the resource.",
				Transform:   transform.FromField("Description.Service.Type")},
			{
				Name:        "hosting_mode",
				Type:        proto.ColumnType_STRING,
				Description: "Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'. Possible values include: 'Default', 'HighDensity'.",
				Transform:   transform.FromField("Description.Service.Properties.HostingMode")},
			{
				Name:        "partition_count",
				Type:        proto.ColumnType_INT,
				Description: "The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.",
				Transform:   transform.FromField("Description.Service.Properties.PartitionCount")},
			{
				Name:        "public_network_access",
				Type:        proto.ColumnType_STRING,
				Description: "This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method. Possible values include: 'Enabled', 'Disabled'.",
				Transform:   transform.FromField("Description.Service.Properties.PublicNetworkAccess")},
			{
				Name:        "replica_count",
				Type:        proto.ColumnType_INT,
				Description: "The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.",
				Transform:   transform.FromField("Description.Service.Properties.ReplicaCount")},
			{
				Name:        "sku_name",
				Type:        proto.ColumnType_STRING,
				Description: "The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new search service.",
				Transform:   transform.FromField("Description.Service.SKU.Name")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the search service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "identity",
				Type:        proto.ColumnType_JSON,
				Description: "The identity of the resource.",
				Transform:   transform.FromField("Description.Service.Identity")},
			{
				Name:        "network_rule_set",
				Type:        proto.ColumnType_JSON,
				Description: "Network specific rules that determine how the azure cognitive search service may be reached.",
				Transform:   transform.FromField("Description.Service.Properties.NetworkRuleSet")},
			{
				Name:        "private_endpoint_connections",
				Type:        proto.ColumnType_JSON,
				Description: "The list of private endpoint connections to the azure cognitive search service.",
				Transform:   transform.FromField("Description.Service.Properties.PrivateEndpointConnections")},
			{
				Name:        "shared_private_link_resources",
				Type:        proto.ColumnType_JSON,
				Description: "The list of shared private link resources managed by the azure cognitive search service.",
				Transform:   transform.FromField("Description.Service.Properties.SharedPrivateLinkResources")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.ID").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}
