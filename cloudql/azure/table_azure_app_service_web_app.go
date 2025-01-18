package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureAppServiceWebApp(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_app_service_web_app",
		Description: "Azure App Service Web App",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetAppServiceWebApp,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAppServiceWebApp,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the app service web app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify an app service web app uniquely.",
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
				Description: "The resource type of the app service web app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Type")},
			{
				Name:        "client_affinity_enabled",
				Description: "Specify whether client affinity is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.ClientAffinityEnabled")},
			{
				Name:        "client_cert_enabled",
				Description: "Specify whether client certificate authentication is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.ClientCertEnabled")},
			{
				Name:        "default_site_hostname",
				Description: "Default hostname of the app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.DefaultHostName")},
			{
				Name:        "enabled",
				Description: "Specify whether the app is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.Enabled")},
			{
				Name:        "host_name_disabled",
				Description: "Specify whether the public hostnames of the app is disabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.HostNamesDisabled")},
			{
				Name:        "https_only",
				Description: "Specify whether configuring a web site to accept only https requests.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.HTTPSOnly")},
			{
				Name:        "identity",
				Description: "Managed service identity for the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(webAppIdentity),
			},
			{
				Name:        "outbound_ip_addresses",
				Description: "List of IP addresses that the app uses for outbound connections (e.g. database access).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.OutboundIPAddresses")},
			{
				Name:        "possible_outbound_ip_addresses",
				Description: "List of possible IP addresses that the app uses for outbound connections (e.g. database access).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.PossibleOutboundIPAddresses")},
			{
				Name:        "reserved",
				Description: "Specify whether the app is reserved.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Site.Properties.Reserved")},
			{
				Name:        "host_names",
				Description: "A list of hostnames associated with the app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Properties.HostNames")},
			{
				Name:        "auth_settings",
				Description: "Describes the Authentication/Authorization settings of an app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SiteAuthSettings")},
			{
				Name:        "configuration",
				Description: "Describes the configuration of an app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SiteConfigResource")},
			{
				Name:        "diagnostic_logs_configuration",
				Description: "Describes the logging configuration of an app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SiteLogConfig")},
			{
				Name:        "site_config",
				Description: "A map of all configuration for the app.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Properties.SiteConfig")},
			{
				Name:        "storage_info_value",
				Description: "AzureStorageInfoValue azure Files or Blob Storage access information value for dictionary storage.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "vnet_connection",
				Description: "Describes the virtual network connection for the app.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.VnetInfo")},

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

				Transform: transform.FromField("Description.Site.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION
				//// TRANSFORM FUNCTION

				Transform: transform.FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}

func webAppIdentity(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.AppServiceWebApp).Description.Site
	objectMap := make(map[string]interface{})
	if data.Identity != nil {
		if types.SafeString(data.Identity.Type) != "" {
			objectMap["Type"] = data.Identity.Type
		}
		if data.Identity.TenantID != nil {
			objectMap["TenantID"] = data.Identity.TenantID
		}
		if data.Identity.PrincipalID != nil {
			objectMap["PrincipalID"] = data.Identity.PrincipalID
		}
		if data.Identity.UserAssignedIdentities != nil {
			objectMap["UserAssignedIdentities"] = data.Identity.UserAssignedIdentities
		}
	}
	return objectMap, nil
}
