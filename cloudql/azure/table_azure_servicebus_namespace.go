package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureServiceBusNamespace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_servicebus_namespace",
		Description: "Azure ServiceBus Namespace",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetServicebusNamespace,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404", "400"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListServicebusNamespace,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SBNamespace.Name")},
			{
				Name:        "id",
				Description: "The unique id identifying the resource in subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SBNamespace.ID")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SBNamespace.Type")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SBNamespace.Properties.ProvisioningState")},
			{
				Name:        "zone_redundant",
				Description: "Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.SBNamespace.Properties.ZoneRedundant")},
			{
				Name:        "created_at",
				Description: "The time the namespace was created.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.SBNamespace.Properties.CreatedAt").Transform(convertDateToTime),
			},
			{
				Name:        "disable_local_auth",
				Description: "This property disables SAS authentication for the Service Bus namespace.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.SBNamespace.Properties.DisableLocalAuth"),
			},
			{
				Name:        "metric_id",
				Description: "The identifier for Azure insights metrics.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SBNamespace.Properties.MetricID")},
			{
				Name:        "servicebus_endpoint",
				Description: "Specifies the endpoint used to perform Service Bus operations.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SBNamespace.Properties.ServiceBusEndpoint")},
			{
				Name:        "sku_capacity",
				Description: "The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.SBNamespace.SKU.Capacity")},
			{
				Name:        "sku_name",
				Description: "Name of this SKU. Valid valuer are: 'Basic', 'Standard', 'Premium'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.SBNamespace.SKU.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The billing tier of this particular SKU. Valid values are: 'Basic', 'Standard', 'Premium'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SBNamespace.SKU.Tier")},
			{
				Name:        "status",
				Description: "Status of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SBNamespace.Properties.Status"),
			},
			{
				Name:        "updated_at",
				Description: "The time the namespace was updated.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.SBNamespace.Properties.UpdatedAt").Transform(convertDateToTime),
			},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the servicebus namespace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "encryption",
				Description: "Specifies the properties of BYOK encryption configuration. Customer-managed key encryption at rest (Bring Your Own Key) is only available on Premium namespaces.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SBNamespace.Properties.Encryption")},
			{
				Name:        "network_rule_set",
				Description: "Describes the network rule set for specified namespace. The ServiceBus Namespace must be Premium in order to attach a ServiceBus Namespace Network Rule Set.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NetworkRuleSet")},
			{
				Name:        "private_endpoint_connections",
				Description: "The private endpoint connections of the namespace.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.SBNamespace.Properties.PrivateEndpointConnections")},
			{
				Name:        "authorization_rules",
				Description: "The authorization rules for a namespace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AuthorizationRules"),
			},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SBNamespace.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SBNamespace.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.SBNamespace.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.SBNamespace.Location").Transform(formatRegion).Transform(toLower),
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
