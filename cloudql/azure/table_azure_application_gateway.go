package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureApplicationGateway(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_application_gateway",
		Description: "Azure Application Gateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetApplicationGateway,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListApplicationGateway,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationGateway.Name")},
			{
				Name:        "id",
				Description: "The resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationGateway.ID")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the application gateway. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationGateway.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationGateway.Type")},
			{
				Name:        "enable_fips",
				Description: "Whether FIPS is enabled on the application gateway.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.ApplicationGateway.Properties.EnableFips"), Default: false,
			},
			{
				Name:        "enable_http2",
				Description: "Whether HTTP2 is enabled on the application gateway.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ApplicationGateway.Properties.EnableHTTP2"),
				Default:     false,
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationGateway.Etag")},
			{
				Name:        "force_firewall_policy_association",
				Description: "If true, associates a firewall policy with an application gateway regardless whether the policy differs from the WAF configuration.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.ApplicationGateway.Properties.ForceFirewallPolicyAssociation"), Default: false,
			},
			{
				Name:        "operational_state",
				Description: "Operational state of the application gateway. Possible values include: 'Stopped', 'Starting', 'Running', 'Stopping'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationGateway.Properties.OperationalState")},
			{
				Name:        "resource_guid",
				Description: "The resource GUID property of the application gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationGateway.Properties.ResourceGUID")},
			{
				Name:        "authentication_certificates",
				Description: "Authentication certificates of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayAuthenticationCertificates),
			},
			{
				Name:        "autoscale_configuration",
				Description: "Autoscale Configuration of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationGateway.Properties.AutoscaleConfiguration")},
			{
				Name:        "backend_address_pools",
				Description: "Backend address pool of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayBackendAddressPools),
			},
			{
				Name:        "backend_http_settings_collection",
				Description: "Backend http settings of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayBackendHTTPSettingsCollection),
			},
			{
				Name:        "custom_error_configurations",
				Description: "Custom error configurations of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationGateway.Properties.CustomErrorConfigurations")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "firewall_policy",
				Description: "Reference to the FirewallPolicy resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationGateway.Properties.FirewallPolicy")},
			{
				Name:        "frontend_ip_configurations",
				Description: "Frontend IP addresses of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayFrontendIPConfigurations),
			},
			{
				Name:        "frontend_ports",
				Description: "Frontend ports of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayFrontendPorts),
			},
			{
				Name:        "gateway_ip_configurations",
				Description: "Subnets of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayIPConfigurations),
			},
			{
				Name:        "http_listeners",
				Description: "Http listeners of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayHTTPListeners),
			},
			{
				Name:        "identity",
				Description: "The identity of the application gateway, if configured.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationGateway.Identity")},
			{
				Name:        "private_endpoint_connections",
				Description: "Private endpoint connections on application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayPrivateEndpointConnections),
			},
			{
				Name:        "private_link_configurations",
				Description: "PrivateLink configurations on application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayPrivateLinkConfigurations),
			},
			{
				Name:        "probes",
				Description: "Probes of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayProbes),
			},
			{
				Name:        "redirect_configurations",
				Description: "Redirect configurations of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationGateway.Properties.RedirectConfigurations")},
			{
				Name:        "request_routing_rules",
				Description: "Request routing rules of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayRequestRoutingRules),
			},
			{
				Name:        "rewrite_rule_sets",
				Description: "Rewrite rules for the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayRewriteRuleSets),
			},
			{
				Name:        "sku",
				Description: "SKU of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationGateway.Properties.SKU"),
			},
			{
				Name:        "ssl_certificates",
				Description: "SSL certificates of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewaySslCertificates),
			},
			{
				Name:        "ssl_policy",
				Description: "SSL policy of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationGateway.Properties.SSLPolicy"),
			},
			{
				Name:        "ssl_profiles",
				Description: "SSL profiles of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewaySslProfiles),
			},
			{
				Name:        "trusted_client_certificates",
				Description: "Trusted client certificates of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayTrustedClientCertificates),
			},
			{
				Name:        "trusted_root_certificates",
				Description: "Trusted root certificates of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayTrustedRootCertificates),
			},
			{
				Name:        "url_path_maps",
				Description: "URL path map of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractGatewayURLPathMaps),
			},
			{
				Name:        "web_application_firewall_configuration",
				Description: "Web application firewall configuration of the application gateway.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationGateway.Properties.WebApplicationFirewallConfiguration")},
			{
				Name:        "zones",
				Description: "A list of availability zones denoting where the resource needs to come from.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.ApplicationGateway.Zones")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ApplicationGateway.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ApplicationGateway.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.ApplicationGateway.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ApplicationGateway.Location").Transform(toLower),
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
// the contents of DiagnosticSettings

//// TRANSFORM FUNCTIONS

// If we return the API response directly, the output will not provide all the properties of GatewayIPConfigurations
func extractGatewayIPConfigurations(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.GatewayIPConfigurations != nil {
		for _, i := range gateway.Properties.GatewayIPConfigurations {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of AuthenticationCertificates
func extractGatewayAuthenticationCertificates(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.AuthenticationCertificates != nil {
		for _, i := range gateway.Properties.AuthenticationCertificates {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of TrustedRootCertificates
func extractGatewayTrustedRootCertificates(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.TrustedRootCertificates != nil {
		for _, i := range gateway.Properties.TrustedRootCertificates {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of TrustedClientCertificates
func extractGatewayTrustedClientCertificates(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.TrustedClientCertificates != nil {
		for _, i := range gateway.Properties.TrustedClientCertificates {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of SslCertificates
func extractGatewaySslCertificates(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.SSLCertificates != nil {
		for _, i := range gateway.Properties.SSLCertificates {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of FrontendIPConfigurations
func extractGatewayFrontendIPConfigurations(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.FrontendIPConfigurations != nil {
		for _, i := range gateway.Properties.FrontendIPConfigurations {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of FrontendPorts
func extractGatewayFrontendPorts(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.FrontendPorts != nil {
		for _, i := range gateway.Properties.FrontendPorts {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of Probes
func extractGatewayProbes(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.Probes != nil {
		for _, i := range gateway.Properties.Probes {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of BackendAddressPools
func extractGatewayBackendAddressPools(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.BackendAddressPools != nil {
		for _, i := range gateway.Properties.BackendAddressPools {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of BackendHTTPSettingsCollection
func extractGatewayBackendHTTPSettingsCollection(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.BackendHTTPSettingsCollection != nil {
		for _, i := range gateway.Properties.BackendHTTPSettingsCollection {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of HTTPListeners
func extractGatewayHTTPListeners(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.HTTPListeners != nil {
		for _, i := range gateway.Properties.HTTPListeners {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of SslProfiles
func extractGatewaySslProfiles(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.SSLProfiles != nil {
		for _, i := range gateway.Properties.SSLProfiles {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of URLPathMaps
func extractGatewayURLPathMaps(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.URLPathMaps != nil {
		for _, i := range gateway.Properties.URLPathMaps {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of RequestRoutingRules
func extractGatewayRequestRoutingRules(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.RequestRoutingRules != nil {
		for _, i := range gateway.Properties.RequestRoutingRules {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Type != nil {
				objectMap["type"] = i.Type
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of RewriteRuleSets
func extractGatewayRewriteRuleSets(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.RewriteRuleSets != nil {
		for _, i := range gateway.Properties.RewriteRuleSets {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of PrivateLinkConfigurations
func extractGatewayPrivateLinkConfigurations(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.PrivateLinkConfigurations != nil {
		for _, i := range gateway.Properties.PrivateLinkConfigurations {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}

// If we return the API response directly, the output will not provide all the properties of PrivateEndpointConnections
func extractGatewayPrivateEndpointConnections(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.ApplicationGateway).Description.ApplicationGateway
	var properties []map[string]interface{}

	if gateway.Properties.PrivateEndpointConnections != nil {
		for _, i := range gateway.Properties.PrivateEndpointConnections {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.Name != nil {
				objectMap["name"] = i.Name
			}
			if i.Etag != nil {
				objectMap["type"] = i.Etag
			}
			if i.Properties != nil {
				objectMap["properties"] = i.Properties
				objectMap["provisioning_state"] = i.Properties.ProvisioningState
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}
