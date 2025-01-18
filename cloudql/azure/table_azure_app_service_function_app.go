package azure

import (
	"context"
	"strconv"
	"strings"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureAppServiceFunctionApp(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_app_service_function_app",
		Description: "Azure App Service Function App",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetAppServiceFunctionApp,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAppServiceFunctionApp,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the app service function app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify an app service function app uniquely.",
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
				Description: "The resource type of the app service function app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Site.Properties.Type")},
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
				Name:        "site_config",
				Description: "A map of all configuration for the app",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Site.Properties.SiteConfig")},
			{
				Name:        "language_runtime_version",
				Description: "The language runtime version of the app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getLanguageRuntimeVersion)},
			{
				Name:        "language_runtime_type",
				Description: "The language runtime type of the app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getLanguageRuntimeType)},
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

				Transform: transform.FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}

func getLanguageAndVersionFromFxVersion(fxVersion string) (string, float64, error) {
	splitFx := strings.Split(fxVersion, "|")
	if len(splitFx) == 2 {
		language, versionStr := strings.Split(fxVersion, "|")[0], strings.Split(fxVersion, "|")[1]
		version, err := strconv.ParseFloat(versionStr, 64)
		if err != nil {
			return "", 0, err
		}
		return strings.ToLower(language), version, nil
	}
	return "", 0, nil
}

func getLanguageRuntimeVersion(ctx context.Context, d *transform.TransformData) (any, error) {
	functionApp := d.HydrateItem.(opengovernance.AppServiceFunctionApp).Description
	if functionApp.Site.Properties.SiteConfig != nil {
		fxVersion := functionApp.Site.Properties.SiteConfig.LinuxFxVersion
		if fxVersion != nil && *fxVersion != "" {
			_, version, err := getLanguageAndVersionFromFxVersion(*fxVersion)
			if err == nil && version != 0 {
				return version, nil
			}
		}
		fxVersion = functionApp.Site.Properties.SiteConfig.WindowsFxVersion
		if fxVersion != nil && *fxVersion != "" {
			_, version, err := getLanguageAndVersionFromFxVersion(*fxVersion)
			if err == nil && version != 0 {
				return version, nil
			}
		}
	}
	if functionApp.SiteConfigResource.Properties != nil {
		switch {
		case functionApp.SiteConfigResource.Properties.PythonVersion != nil && *functionApp.SiteConfigResource.Properties.PythonVersion != "":
			return *functionApp.SiteConfigResource.Properties.PythonVersion, nil
		case functionApp.SiteConfigResource.Properties.PowerShellVersion != nil && *functionApp.SiteConfigResource.Properties.PowerShellVersion != "":
			return *functionApp.SiteConfigResource.Properties.PowerShellVersion, nil
		case functionApp.SiteConfigResource.Properties.JavaVersion != nil && *functionApp.SiteConfigResource.Properties.JavaVersion != "":
			return *functionApp.SiteConfigResource.Properties.JavaVersion, nil
		case functionApp.SiteConfigResource.Properties.NodeVersion != nil && *functionApp.SiteConfigResource.Properties.NodeVersion != "":
			return *functionApp.SiteConfigResource.Properties.NodeVersion, nil
		case functionApp.SiteConfigResource.Properties.NetFrameworkVersion != nil && *functionApp.SiteConfigResource.Properties.NetFrameworkVersion != "":
			return *functionApp.SiteConfigResource.Properties.NetFrameworkVersion, nil
		}
	}
	return nil, nil
}

func getLanguageRuntimeType(ctx context.Context, d *transform.TransformData) (any, error) {
	functionApp := d.HydrateItem.(opengovernance.AppServiceFunctionApp).Description
	if functionApp.Site.Properties.SiteConfig != nil {
		fxVersion := functionApp.Site.Properties.SiteConfig.LinuxFxVersion
		if fxVersion != nil && *fxVersion != "" {
			language, _, err := getLanguageAndVersionFromFxVersion(*fxVersion)
			if err == nil && language != "" {
				return language, nil
			}
		}
		fxVersion = functionApp.Site.Properties.SiteConfig.WindowsFxVersion
		if fxVersion != nil && *fxVersion != "" {
			language, _, err := getLanguageAndVersionFromFxVersion(*fxVersion)
			if err == nil && language != "" {
				return language, nil
			}
		}
	}
	if functionApp.SiteConfigResource.Properties != nil {
		switch {
		case functionApp.SiteConfigResource.Properties.PythonVersion != nil && *functionApp.SiteConfigResource.Properties.PythonVersion != "":
			return "python", nil
		case functionApp.SiteConfigResource.Properties.PowerShellVersion != nil && *functionApp.SiteConfigResource.Properties.PowerShellVersion != "":
			return "powershell", nil
		case functionApp.SiteConfigResource.Properties.JavaVersion != nil && *functionApp.SiteConfigResource.Properties.JavaVersion != "":
			return "java", nil
		case functionApp.SiteConfigResource.Properties.NodeVersion != nil && *functionApp.SiteConfigResource.Properties.NodeVersion != "":
			return "node", nil
		case functionApp.SiteConfigResource.Properties.NetFrameworkVersion != nil && *functionApp.SiteConfigResource.Properties.NetFrameworkVersion != "":
			return "dotnetcore", nil
		}
	}
	return nil, nil
}
