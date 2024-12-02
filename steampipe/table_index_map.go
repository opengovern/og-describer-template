package steampipe

import (
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"
)

var Map = map[string]string{
	"Microsoft.Entra/groups":                    "entraid_group",
	"Microsoft.Entra/groupMemberships":          "entraid_group_membership",
	"Microsoft.Entra/devices":                   "entraid_device",
	"Microsoft.Entra/applications":              "entraid_application",
	"Microsoft.Entra/appRegistrations":          "entraid_app_registration",
	"Microsoft.Entra/enterpriseApplication":     "entraid_enterprise_application",
	"Microsoft.Entra/managedIdentity":           "entraid_managed_identity",
	"Microsoft.Entra/microsoftApplication":      "entraid_microsoft_application",
	"Microsoft.Entra/domains":                   "entraid_domain",
	"Microsoft.Entra/tenant":                    "entraid_tenant",
	"Microsoft.Entra/identityproviders":         "entraid_identity_provider",
	"Microsoft.Entra/securitydefaultspolicy":    "entraid_security_defaults_policy",
	"Microsoft.Entra/authorizationpolicy":       "entraid_authorization_policy",
	"Microsoft.Entra/conditionalaccesspolicy":   "entraid_conditional_access_policy",
	"Microsoft.Entra/adminconsentrequestpolicy": "entraid_admin_consent_request_policy",
	"Microsoft.Entra/userregistrationdetails":   "entraid_user_registration_details",
	"Microsoft.Entra/serviceprincipals":         "entraid_service_principal",
	"Microsoft.Entra/users":                     "entraid_user",
	"Microsoft.Entra/directoryroles":            "entraid_directory_role",
	"Microsoft.Entra/directorysettings":         "entraid_directory_setting",
}

var DescriptionMap = map[string]interface{}{
	"Microsoft.Entra/groups":                    opengovernance.AdGroup{},
	"Microsoft.Entra/groupMemberships":          opengovernance.AdGroupMembership{},
	"Microsoft.Entra/devices":                   opengovernance.AdDevice{},
	"Microsoft.Entra/applications":              opengovernance.AdApplication{},
	"Microsoft.Entra/appRegistrations":          opengovernance.AdAppRegistration{},
	"Microsoft.Entra/enterpriseApplication":     opengovernance.AdEnterpriseApplication{},
	"Microsoft.Entra/managedIdentity":           opengovernance.AdManagedIdentity{},
	"Microsoft.Entra/microsoftApplication":      opengovernance.AdMicrosoftApplication{},
	"Microsoft.Entra/domains":                   opengovernance.AdDomain{},
	"Microsoft.Entra/tenant":                    opengovernance.AdTenant{},
	"Microsoft.Entra/identityproviders":         opengovernance.AdIdentityProvider{},
	"Microsoft.Entra/securitydefaultspolicy":    opengovernance.AdSecurityDefaultsPolicy{},
	"Microsoft.Entra/authorizationpolicy":       opengovernance.AdAuthorizationPolicy{},
	"Microsoft.Entra/conditionalaccesspolicy":   opengovernance.AdConditionalAccessPolicy{},
	"Microsoft.Entra/adminconsentrequestpolicy": opengovernance.AdAdminConsentRequestPolicy{},
	"Microsoft.Entra/userregistrationdetails":   opengovernance.AdUserRegistrationDetails{},
	"Microsoft.Entra/serviceprincipals":         opengovernance.AdServicePrincipal{},
	"Microsoft.Entra/users":                     opengovernance.AdUsers{},
	"Microsoft.Entra/directoryroles":            opengovernance.AdDirectoryRole{},
	"Microsoft.Entra/directorysettings":         opengovernance.AdDirectorySetting{},
}

var ReverseMap = map[string]string{
	"entraid_group":                        "Microsoft.Entra/groups",
	"entraid_group_membership":             "Microsoft.Entra/groupMemberships",
	"entraid_device":                       "Microsoft.Entra/devices",
	"entraid_application":                  "Microsoft.Entra/applications",
	"entraid_app_registration":             "Microsoft.Entra/appRegistrations",
	"entraid_enterprise_application":       "Microsoft.Entra/enterpriseApplication",
	"entraid_managed_identity":             "Microsoft.Entra/managedIdentity",
	"entraid_microsoft_application":        "Microsoft.Entra/microsoftApplication",
	"entraid_domain":                       "Microsoft.Entra/domains",
	"entraid_tenant":                       "Microsoft.Entra/tenant",
	"entraid_identity_provider":            "Microsoft.Entra/identityproviders",
	"entraid_security_defaults_policy":     "Microsoft.Entra/securitydefaultspolicy",
	"entraid_authorization_policy":         "Microsoft.Entra/authorizationpolicy",
	"entraid_conditional_access_policy":    "Microsoft.Entra/conditionalaccesspolicy",
	"entraid_admin_consent_request_policy": "Microsoft.Entra/adminconsentrequestpolicy",
	"entraid_user_registration_details":    "Microsoft.Entra/userregistrationdetails",
	"entraid_service_principal":            "Microsoft.Entra/serviceprincipals",
	"entraid_user":                         "Microsoft.Entra/users",
	"entraid_directory_role":               "Microsoft.Entra/directoryroles",
	"entraid_directory_setting":            "Microsoft.Entra/directorysettings",
}
