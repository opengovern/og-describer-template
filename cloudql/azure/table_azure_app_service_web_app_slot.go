package azure

import (
	"context"
	"strings"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-06-01/web"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureAppServiceWebAppSlot(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_app_service_web_app_slot",
		Description: "Azure App Service Web App Slot",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "app_name", "resource_group"}),
			Hydrate:    opengovernance.GetAppServiceWebAppSlot,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAppServiceWebAppSlot,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Resource Name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Name").Transform(extractName)},
			{
				Name:        "app_name",
				Description: "The name of the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AppName")},
			{
				Name:        "id",
				Description: "Resource ID of the app slot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.ID")},
			{
				Name:        "kind",
				Description: "Contains the kind of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Kind")},
			{
				Name:        "state",
				Description: "Current state of the app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.State")},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Type")},
			{
				Name:        "last_modified_time_utc",
				Description: "Last time the app was modified, in UTC.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Site.Properties.LastModifiedTimeUTC")},
			{
				Name:        "repository_site_name",
				Description: "Name of the repository site.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.RepositorySiteName")},
			{
				Name:        "usage_state",
				Description: "State indicating whether the app has exceeded its quota usage. Read-only. Possible values include: 'UsageStateNormal', 'UsageStateExceeded'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.UsageState")},
			{
				Name:        "enabled",
				Description: "Indicates wheather the app is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.Enabled")},
			{
				Name:        "availability_state",
				Description: "Management information availability state for the app. Possible values include: 'Normal', 'Limited', 'DisasterRecoveryMode'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.AvailabilityState")},
			{
				Name:        "server_farm_id",
				Description: "Resource ID of the associated App Service plan, formatted as: '/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.ServerFarmID")},
			{
				Name:        "reserved",
				Description: "True if reserved; otherwise, false.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.Reserved")},
			{
				Name:        "is_xenon",
				Description: "Obsolete: Hyper-V sandbox.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.IsXenon")},
			{
				Name:        "hyper_v",
				Description: "Hyper-V sandbox.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.HyperV")},
			{
				Name:        "scm_site_also_stopped",
				Description: "True to stop SCM (KUDU) site when the app is stopped; otherwise, false. The default is false.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.ScmSiteAlsoStopped")},
			{
				Name:        "target_swap_slot",
				Description: "Specifies which deployment slot this app will swap into.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.TargetSwapSlot")},
			{
				Name:        "client_affinity_enabled",
				Description: "True to enable client affinity; false to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is true.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.ClientAffinityEnabled")},
			{
				Name:        "client_cert_mode",
				Description: "This composes with ClientCertEnabled setting. ClientCertEnabled: false means ClientCert is ignored. ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required.ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted. Possible values include: 'Required', 'Optional'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.ClientCertMode")},
			{
				Name:        "client_cert_exclusion_paths",
				Description: "Client certificate authentication comma-separated exclusion paths.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.ClientCertExclusionPaths")},
			{
				Name:        "host_names_disabled",
				Description: "True to disable the public hostnames of the app; otherwise, false. If true, the app is only accessible via API management process.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.HostNamesDisabled")},
			{
				Name:        "custom_domain_verification_id",
				Description: "Unique identifier that verifies the custom domains assigned to the app. The customer will add this ID to a text record for verification.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.CustomDomainVerificationID")},
			{
				Name:        "outbound_ip_addresses",
				Description: "List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from tenants that site can be hosted with current settings.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.OutboundIPAddresses")},
			{
				Name:        "possible_outbound_ip_addresses",
				Description: "List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from all tenants except dataComponent.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.PossibleOutboundIPAddresses")},
			{
				Name:        "container_size",
				Description: "Size of the function container.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Site.Properties.ContainerSize")},
			{
				Name:        "suspended_till",
				Description: "App suspended till in case memory-time quota is exceeded.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Site.Properties.SuspendedTill")},
			{
				Name:        "is_default_container",
				Description: "True if the app is a default container; otherwise, false.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.IsDefaultContainer")},
			{
				Name:        "default_host_name",
				Description: "Default hostname of the app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.DefaultHostName")},
			{
				Name:        "https_only",
				Description: "Configures a web site to accept only https requests.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.HTTPSOnly")},
			{
				Name:        "redundancy_mode",
				Description: "Site redundancy mode. Possible values include: 'RedundancyModeNone', 'RedundancyModeManual', 'RedundancyModeFailover', 'RedundancyModeActiveActive', 'RedundancyModeGeoRedundant'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.RedundancyMode")},

			// JSON fields
			{
				Name:        "identity",
				Description: "Managed service identity.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Identity")},
			{
				Name:        "host_names",
				Description: "Hostnames associated with the app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Properties.HostNames")},
			{
				Name:        "enabled_host_names",
				Description: "Enabled hostnames for the app. Hostnames need to be assigned (see HostNames) AND enabled. Otherwise, the app is not served on those hostnames.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Properties.EnabledHostNames")},
			{
				Name:        "host_name_ssl_states",
				Description: "Hostname SSL states are used to manage the SSL bindings for app's hostnames.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Properties.HostNameSSLStates")},
			{
				Name:        "site_config",
				Description: "Configuration of the app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Properties.SiteConfig")},
			{
				Name:        "traffic_manager_host_names",
				Description: "Azure Traffic Manager hostnames associated with the app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Properties.TrafficManagerHostNames")},
			{
				Name:        "hosting_environment_profile",
				Description: "App Service Environment to use for the app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Properties.HostingEnvironmentProfile")},
			{
				Name:        "slot_swap_status",
				Description: "Status of the last deployment slot swap operation.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Properties.SlotSwapStatus")},

			{
				Name:        "site_config_resource",
				Description: "Configuration of an app, such as platform version and bitness, default documents, virtual applications, Always On, etc.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Properties.SiteConfig"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Location").Transform(toLower),
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

type SlotInfo struct {
	SiteProperties *web.SiteProperties
	Identity       *web.ManagedServiceIdentity
	ID             *string
	Name           *string
	AppName        *string
	Kind           *string
	Location       *string
	Type           *string
	Tags           map[string]*string
}

func extractName(_ context.Context, d *transform.TransformData) (interface{}, error) {
	name := d.HydrateItem.(opengovernance.AppServiceWebAppSlot).Description.Site.Name
	parts := strings.Split(*name, "/")
	if len(parts) > 1 {
		return parts[1], nil
	} else {
		return nil, nil
	}
}
