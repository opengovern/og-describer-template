package maps
import (
	"github.com/opengovern/og-describer-entraid/discovery/describers"
	"github.com/opengovern/og-describer-entraid/discovery/provider"
	"github.com/opengovern/og-describer-entraid/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
	model "github.com/opengovern/og-describer-entraid/discovery/pkg/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"Microsoft.Entra/groups": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/groups",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdGroup),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/groupMemberships": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/groupMemberships",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdGroupMembership),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/devices": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/devices",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdDevice),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/applications": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/applications",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdApplication),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/appRegistrations": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/appRegistrations",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdAppRegistration),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/enterpriseApplication": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/enterpriseApplication",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdEnterpriseApplication),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/managedIdentity": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/managedIdentity",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdManagedIdentity),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/microsoftApplication": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/microsoftApplication",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdMicrosoftApplication),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/domains": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/domains",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdDomain),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/tenant": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/tenant",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdTenant),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/identityproviders": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/identityproviders",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdIdentityProvider),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/securitydefaultspolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/securitydefaultspolicy",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdSecurityDefaultsPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/authorizationpolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/authorizationpolicy",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdAuthorizationPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/conditionalaccesspolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/conditionalaccesspolicy",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdConditionalAccessPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/adminconsentrequestpolicy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/adminconsentrequestpolicy",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdAdminConsentRequestPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/userregistrationdetails": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/userregistrationdetails",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdUserRegistrationDetails),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/serviceprincipals": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/serviceprincipals",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdServicePrinciple),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/users": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/users",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20AD%20User.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdUsers),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/directoryroles": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/directoryroles",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdDirectoryRole),
		GetDescriber:         nil,
	},

	"Microsoft.Entra/directorysettings": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "Microsoft.Entra/directorysettings",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeADByTenantID(describers.AdDirectorySetting),
		GetDescriber:         nil,
	},
}


var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"Microsoft.Entra/groups": {
		Name:         "Microsoft.Entra/groups",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/groupMemberships": {
		Name:         "Microsoft.Entra/groupMemberships",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/devices": {
		Name:         "Microsoft.Entra/devices",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/applications": {
		Name:         "Microsoft.Entra/applications",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/appRegistrations": {
		Name:         "Microsoft.Entra/appRegistrations",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/enterpriseApplication": {
		Name:         "Microsoft.Entra/enterpriseApplication",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/managedIdentity": {
		Name:         "Microsoft.Entra/managedIdentity",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/microsoftApplication": {
		Name:         "Microsoft.Entra/microsoftApplication",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/domains": {
		Name:         "Microsoft.Entra/domains",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/tenant": {
		Name:         "Microsoft.Entra/tenant",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/identityproviders": {
		Name:         "Microsoft.Entra/identityproviders",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/securitydefaultspolicy": {
		Name:         "Microsoft.Entra/securitydefaultspolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/authorizationpolicy": {
		Name:         "Microsoft.Entra/authorizationpolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/conditionalaccesspolicy": {
		Name:         "Microsoft.Entra/conditionalaccesspolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/adminconsentrequestpolicy": {
		Name:         "Microsoft.Entra/adminconsentrequestpolicy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/userregistrationdetails": {
		Name:         "Microsoft.Entra/userregistrationdetails",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/serviceprincipals": {
		Name:         "Microsoft.Entra/serviceprincipals",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/users": {
		Name:         "Microsoft.Entra/users",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/directoryroles": {
		Name:         "Microsoft.Entra/directoryroles",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"Microsoft.Entra/directorysettings": {
		Name:         "Microsoft.Entra/directorysettings",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},
}


var ResourceTypesList = []string{
  "Microsoft.Entra/groups",
  "Microsoft.Entra/groupMemberships",
  "Microsoft.Entra/devices",
  "Microsoft.Entra/applications",
  "Microsoft.Entra/appRegistrations",
  "Microsoft.Entra/enterpriseApplication",
  "Microsoft.Entra/managedIdentity",
  "Microsoft.Entra/microsoftApplication",
  "Microsoft.Entra/domains",
  "Microsoft.Entra/tenant",
  "Microsoft.Entra/identityproviders",
  "Microsoft.Entra/securitydefaultspolicy",
  "Microsoft.Entra/authorizationpolicy",
  "Microsoft.Entra/conditionalaccesspolicy",
  "Microsoft.Entra/adminconsentrequestpolicy",
  "Microsoft.Entra/userregistrationdetails",
  "Microsoft.Entra/serviceprincipals",
  "Microsoft.Entra/users",
  "Microsoft.Entra/directoryroles",
  "Microsoft.Entra/directorysettings",
}