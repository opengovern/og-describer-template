package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureExpressRouteCircuit(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_express_route_circuit",
		Description: "Azure Express Route Circuit",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetExpressRouteCircuit,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListExpressRouteCircuit,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the circuit.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ExpressRouteCircuit.Name")},
			{
				Name:        "id",
				Description: "Resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ExpressRouteCircuit.ID")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ExpressRouteCircuit.Etag")},
			{
				Name:        "sku_name",
				Description: "The name of the SKU.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ExpressRouteCircuit.SKU.Name")},
			{
				Name:        "sku_tier",
				Description: "The tier of the SKU. Possible values include: 'Standard', 'Premium', 'Basic', 'Local'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ExpressRouteCircuit.SKU.Tier"),
			},
			{
				Name:        "sku_family",
				Description: "The family of the SKU. Possible values include: 'UnlimitedData', 'MeteredData'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ExpressRouteCircuit.SKU.Family"),
			},
			{
				Name:        "allow_classic_operations",
				Description: "Allow classic operations.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.AllowClassicOperations")},
			{
				Name:        "circuit_provisioning_state",
				Description: "The CircuitProvisioningState state of the resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.CircuitProvisioningState")},
			{
				Name:        "service_provider_provisioning_state",
				Description: "The ServiceProviderProvisioningState state of the resource. Possible values include: 'NotProvisioned', 'Provisioning', 'Provisioned', 'Deprovisioning'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.ServiceProviderProvisioningState"),
			},
			{
				Name:        "authorizations",
				Description: "The list of authorizations.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.Authorizations")},
			{
				Name:        "peerings",
				Description: "The list of peerings.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.Peerings")},
			{
				Name:        "service_key",
				Description: "The ServiceKey.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.ServiceKey")},
			{
				Name:        "service_provider_notes",
				Description: "The ServiceProviderNotes.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.ServiceProviderNotes")},
			{
				Name:        "service_provider_properties",
				Description: "The ServiceProviderProperties.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.ServiceProviderProperties")},
			{
				Name:        "express_route_port",
				Description: "The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.ExpressRoutePort")},
			{
				Name:        "bandwidth_in_gbps",
				Description: "The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.",
				Type:        proto.ColumnType_DOUBLE,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.BandwidthInGbps")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the express route circuit resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.ProvisioningState"),
			},
			{
				Name:        "global_reach_enabled",
				Description: "Flag denoting global reach status.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Properties.GlobalReachEnabled")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ExpressRouteCircuit.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ExpressRouteCircuit.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ExpressRouteCircuit.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				// Check if context has been cancelled or if the limit has been hit (if specified)
				// if there is a limit, it will return the number of rows required to reach this limit
				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.ResourceGroup")},
		}),
	}
}

//// HYDRATE FUNCTIONS

// Create session

// Handle empty name or resourceGroup
