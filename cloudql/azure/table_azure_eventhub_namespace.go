package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureEventHubNamespace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_eventhub_namespace",
		Description: "Azure Event Hub Namespace",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetEventhubNamespace,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "400", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEventhubNamespace,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EHNamespace.Name")},
			{
				Name:        "id",
				Description: "The ID of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EHNamespace.ID")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EHNamespace.Type")},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EHNamespace.Properties.ProvisioningState")},
			{
				Name:        "created_at",
				Description: "The time the namespace was created.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.EHNamespace.Properties.CreatedAt").Transform(convertDateToTime),
			},
			{
				Name:        "cluster_arm_id",
				Description: "Cluster ARM ID of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EHNamespace.Properties.ClusterArmID")},
			{
				Name:        "is_auto_inflate_enabled",
				Description: "Indicates whether auto-inflate is enabled for eventhub namespace.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EHNamespace.Properties.IsAutoInflateEnabled")},
			{
				Name:        "kafka_enabled",
				Description: "Indicates whether kafka is enabled for eventhub namespace, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EHNamespace.Properties.KafkaEnabled")},
			{
				Name:        "maximum_throughput_units",
				Description: "Upper limit of throughput units when auto-inflate is enabled, value should be within 0 to 20 throughput units.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EHNamespace.Properties.MaximumThroughputUnits")},
			{
				Name:        "metric_id",
				Description: "Identifier for azure insights metrics.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EHNamespace.Properties.MetricID")},
			{
				Name:        "service_bus_endpoint",
				Description: "Endpoint you can use to perform service bus operations.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EHNamespace.Properties.ServiceBusEndpoint")},
			{
				Name:        "sku_capacity",
				Description: "The Event Hubs throughput units, value should be 0 to 20 throughput units.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.EHNamespace.SKU.Capacity")},
			{
				Name:        "sku_name",
				Description: "Name of this SKU. Possible values include: 'Basic', 'Standard'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.EHNamespace.SKU.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The billing tier of this particular SKU. Valid values are: 'Basic', 'Standard', 'Premium'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EHNamespace.SKU.Tier")},
			{
				Name:        "updated_at",
				Description: "The time the namespace was updated.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.EHNamespace.Properties.UpdatedAt").Transform(convertDateToTime),
			},
			{
				Name:        "zone_redundant",
				Description: "Enabling this property creates a standard event hubs namespace in regions supported availability zones.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EHNamespace.Properties.ZoneRedundant")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the eventhub namespace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "encryption",
				Description: "Properties of BYOK encryption description.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EHNamespace.Properties.Encryption")},
			{
				Name:        "identity",
				Description: "Describes the properties of BYOK encryption description.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EHNamespace.Properties.Encryption")},
			{
				Name:        "network_rule_set",
				Description: "Describes the network rule set for specified namespace. The EventHub Namespace must be Premium in order to attach a EventHub Namespace Network Rule Set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NetworkRuleSet")},
			{
				Name:        "private_endpoint_connections",
				Description: "The private endpoint connections of the namespace.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.EHNamespace.Properties.PrivateEndpointConnections")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EHNamespace.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EHNamespace.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.EHNamespace.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.EHNamespace.Location").Transform(formatRegion).Transform(toLower),
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
