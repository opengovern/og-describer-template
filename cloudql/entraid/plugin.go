package entraid

import (
	"context"
	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-entraid"

// Plugin creates this (entraid) plugin
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
			"entraid_admin_consent_request_policy":           tableEntraIdAdminConsentRequestPolicy(ctx),
			"entraid_application":                            tableEntraIdApplication(ctx),
			"entraid_app_registration":                       tableEntraIdAppRegistration(ctx),
			"entraid_authorization_policy":                   tableEntraIdAuthorizationPolicy(ctx),
			"entraid_conditional_access_policy":              tableEntraIdConditionalAccessPolicy(ctx),
			"entraid_directory_role":                         tableEntraIdDirectoryRole(ctx),
			"entraid_directory_setting":                      tableEntraIdDirectorySetting(ctx),
			"entraid_domain":                                 tableEntraIdDomain(ctx),
			"entraid_group":                                  tableEntraIdGroup(ctx),
			"entraid_tenant":                                 tableEntraIdTenant(ctx),
			"entraid_group_membership":                       tableEntraIdGroupMembership(ctx),
			"entraid_identity_provider":                      tableEntraIdIdentityProvider(ctx),
			"entraid_security_defaults_policy":               tableEntraIdSecurityDefaultsPolicy(ctx),
			"entraid_service_principal":                      tableEntraIdServicePrincipal(ctx),
			"entraid_enterprise_application":                 tableEntraIdEnterpriseApplication(ctx),
			"entraid_managed_identity":                       tableEntraIdManagedIdentity(ctx),
			"entraid_microsoft_application":                  tableEntraIdMicrosoftApplication(ctx),
			"entraid_user":                                   tableEntraIdUser(ctx),
			"entraid_device":                                 tableEntraIdDevice(ctx),
			"entraid_user_registration_details":              tableEntraIdUserRegistrationDetails(ctx),
			"entraid_application_app_role_assigned_to":       tableEntraIdApplicationAppRoleAssignment(ctx),
			"entraid_service_principal_app_role_assigned_to": tableEntraIdServicePrincipalAppRoleAssignedTo(ctx),
			"entraid_service_principal_app_role_assignment":  tableEntraIdServicePrincipalAppRoleAssignment(ctx),
			"entraid_user_app_role_assignment":               tableEntraIdUserAppRoleAssignment(ctx),
		},
	}

	return p
}
