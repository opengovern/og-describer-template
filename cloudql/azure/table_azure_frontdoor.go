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

func tableAzureFrontDoor(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_frontdoor",
		Description: "Azure Front Door",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetFrontdoor,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListFrontdoor,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FrontDoor.Name")},
			{
				Name:        "id",
				Description: "Fully qualified resource ID for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FrontDoor.ID")},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the front door.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FrontDoor.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FrontDoor.Type")},
			{
				Name:        "cname",
				Description: "The host that each frontendEndpoint must CNAME to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getCname)},
			{
				Name:        "enabled_state",
				Description: "Operational status of the front door load balancer. Possible values include: 'Enabled', 'Disabled'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FrontDoor.Properties.EnabledState")},
			{
				Name:        "friendly_name",
				Description: "A friendly name for the front door.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FrontDoor.Properties.FriendlyName")},
			{
				Name:        "front_door_id",
				Description: "The ID of the front door.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FrontDoor.Properties.FrontdoorID")},
			{
				Name:        "resource_state",
				Description: "Resource status of the front door. Possible values include: 'Creating', 'Enabling', 'Enabled', 'Disabling', 'Disabled', 'Deleting'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FrontDoor.Properties.ResourceState")},
			{
				Name:        "backend_pools",
				Description: "Backend pools available to routing rules.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FrontDoor.Properties.BackendPools")},
			{
				Name:        "backend_pools_settings",
				Description: "Settings for all backend pools",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FrontDoor.Properties.BackendPoolsSettings")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "frontend_endpoints",
				Description: "Frontend endpoints available to routing rules.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FrontDoor.Properties.FrontendEndpoints")},
			{
				Name:        "health_probe_settings",
				Description: "Health probe settings associated with this Front Door instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FrontDoor.Properties.HealthProbeSettings")},
			{
				Name:        "load_balancing_settings",
				Description: "Load balancing settings associated with this front door instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FrontDoor.Properties.LoadBalancingSettings")},
			{
				Name:        "rules_engines",
				Description: "Rules engine configurations available to routing rules.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FrontDoor.Properties.RulesEngines")},
			{
				Name:        "routing_rules",
				Description: "Routing rules associated with this front door.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.FrontDoor.Properties.RoutingRules")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FrontDoor.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FrontDoor.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.FrontDoor.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.FrontDoor.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				//// HYDRATE FUNCTIONS
				Transform: transform.

					// Handle empty name or resourceGroup
					FromField("Description.ResourceGroup")},
		}),
	}
}

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

// Create session

// If we return the API response directly, the output does not provide
// all the contents of DiagnosticSettings

func getCname(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	frontDoor := d.HydrateItem.(opengovernance.Frontdoor).Description.FrontDoor
	return strings.ToLower(*frontDoor.Properties.FrontendEndpoints[0].Properties.HostName), nil
}
