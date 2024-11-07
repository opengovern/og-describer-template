package azuread

import (
	"context"
	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-azuread"

// Plugin creates this (azuread) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_ResourceNotFound"}),
			},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: essdk.ConfigInstance,
			Schema:      essdk.ConfigSchema(),
		},
		TableMap: map[string]*plugin.Table{
			"azuread_admin_consent_request_policy": tableAzureAdAdminConsentRequestPolicy(ctx),
			"azuread_application":                  tableAzureAdApplication(ctx),
			"azuread_app_registration":             tableAzureAdAppRegistration(ctx),
			"azuread_authorization_policy":         tableAzureAdAuthorizationPolicy(ctx),
			"azuread_conditional_access_policy":    tableAzureAdConditionalAccessPolicy(ctx),
			"azuread_directory_audit_report":       tableAzureAdDirectoryAuditReport(ctx),
			"azuread_directory_role":               tableAzureAdDirectoryRole(ctx),
			"azuread_directory_setting":            tableAzureAdDirectorySetting(ctx),
			"azuread_domain":                       tableAzureAdDomain(ctx),
			"azuread_group":                        tableAzureAdGroup(ctx),
			"azuread_tenant":                       tableAzureAdTenant(ctx),
			"azuread_group_membership":             tableAzureAdGroupMembership(ctx),
			"azuread_identity_provider":            tableAzureAdIdentityProvider(ctx),
			"azuread_security_defaults_policy":     tableAzureAdSecurityDefaultsPolicy(ctx),
			"azuread_service_principal":            tableAzureAdServicePrincipal(ctx),
			"azuread_enterprise_application":       tableAzureAdEnterpriseApplication(ctx),
			"azuread_managed_identity":             tableAzureAdManagedIdentity(ctx),
			"azuread_microsoft_application":        tableAzureAdMicrosoftApplication(ctx),
			"azuread_user":                         tableAzureAdUser(ctx),
			"azuread_device":                       tableAzureAdDevice(ctx),
			"azuread_sign_in_report":               tableAzureAdSignInReport(ctx),
			"azuread_user_registration_details":    tableAzureAdUserRegistrationDetails(ctx),
		},
	}

	return p
}
