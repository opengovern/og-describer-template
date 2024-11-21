package provider

import (
	model "github.com/opengovern/og-describer-entraid/pkg/sdk/models"
	"github.com/opengovern/og-describer-entraid/provider/configs"
	"github.com/opengovern/og-describer-entraid/provider/describer"
)
var ResourceTypes = map[string]model.ResourceType{

	"Microsoft.Entra/groups": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/groups",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdGroup),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/groupMemberships": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/groupMemberships",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdGroupMembership),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/devices": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/devices",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdDevice),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/signInReports": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/signInReports",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdSignInReport),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/applications": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/applications",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdApplication),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/appRegistrations": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/appRegistrations",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdAppRegistration),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/enterpriseApplication": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/enterpriseApplication",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdEnterpriseApplication),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/managedIdentity": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/managedIdentity",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdManagedIdentity),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/microsoftApplication": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/microsoftApplication",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdMicrosoftApplication),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/domains": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/domains",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdDomain),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/tenant": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/tenant",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdTenant),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/identityproviders": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/identityproviders",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdIdentityProvider),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/securitydefaultspolicy": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/securitydefaultspolicy",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdSecurityDefaultsPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/authorizationpolicy": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/authorizationpolicy",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdAuthorizationPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/conditionalaccesspolicy": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/conditionalaccesspolicy",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdConditionalAccessPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/adminconsentrequestpolicy": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/adminconsentrequestpolicy",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdAdminConsentRequestPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/userregistrationdetails": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/userregistrationdetails",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdUserRegistrationDetails),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/serviceprincipals": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/serviceprincipals",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdServicePrinciple),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/users": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/users",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20User.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdUsers),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/directoryroles": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/directoryroles",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdDirectoryRole),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/directorysettings": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/directorysettings",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdDirectorySetting),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/directoryauditreport": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Entra/directoryauditreport",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeADByTenantID(describer.AdDirectoryAuditReport),
		GetDescriber:         nil,
	},
}
