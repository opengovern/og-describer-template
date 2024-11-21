package entraid

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableEntraIdConditionalAccessPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_conditional_access_policy",
		Description: "Represents an Azure Active Directory (Azure AD) Conditional Access Policy.",
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetAdConditionalAccessPolicy,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_ResourceNotFound", "Invalid object identifier"}),
			},
			KeyColumns: plugin.SingleColumn("id"),
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdConditionalAccessPolicy,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_UnsupportedQuery"}),
			},
		},

		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Specifies the identifier of a conditionalAccessPolicy object.",
				Transform:   transform.FromField("Description.Id")},
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "Specifies a display name for the conditionalAccessPolicy object.",
				Transform:   transform.FromField("Description.DisplayName")},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "Specifies the state of the conditionalAccessPolicy object. Possible values are: enabled, disabled, enabledForReportingButNotEnforced.",
				Transform:   transform.FromField("Description.State")},
			{
				Name:        "created_date_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The create date of the conditional access policy.",
				Transform:   transform.FromField("Description.CreatedDateTime")},
			{
				Name:        "modified_date_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The modification date of the conditional access policy.",
				Transform:   transform.FromField("Description.ModifiedDateTime")},
			{
				Name:        "operator",
				Type:        proto.ColumnType_STRING,
				Description: "Defines the relationship of the grant controls. Possible values: AND, OR.",
				Transform:   transform.FromField("Description.Operator")},

			// Json fields
			{
				Name:        "applications",
				Type:        proto.ColumnType_JSON,
				Description: "Applications and user actions included in and excluded from the policy.",
				Transform:   transform.FromField("Description.Applications")},
			{
				Name:        "application_enforced_restrictions",
				Type:        proto.ColumnType_JSON,
				Description: "Session control to enforce application restrictions. Only Exchange Online and Sharepoint Online support this session control.",
				Transform:   transform.FromField("Description.ApplicationEnforcedRestrictions")},
			{
				Name:        "built_in_controls",
				Type:        proto.ColumnType_JSON,
				Description: "List of values of built-in controls required by the policy. Possible values: block, mfa, compliantDevice, domainJoinedDevice, approvedApplication, compliantApplication, passwordChange, unknownFutureValue.",
				Transform:   transform.FromField("Description.BuiltInControls")},
			{
				Name:        "client_app_types",
				Type:        proto.ColumnType_JSON,
				Description: "Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other.",
				Transform:   transform.FromField("Description.ClientAppTypes")},
			{
				Name:        "custom_authentication_factors",
				Type:        proto.ColumnType_JSON,
				Description: "List of custom controls IDs required by the policy.",
				Transform:   transform.FromField("Description.CustomAuthenticationFactors")},
			{
				Name:        "cloud_app_security",
				Type:        proto.ColumnType_JSON,
				Description: "Session control to apply cloud app security.",
				Transform:   transform.FromField("Description.CloudAppSecurity")},
			{
				Name:        "locations",
				Type:        proto.ColumnType_JSON,
				Description: "Locations included in and excluded from the policy.",
				Transform:   transform.FromField("Description.Locations")},
			{
				Name:        "persistent_browser",
				Type:        proto.ColumnType_JSON,
				Description: "Session control to define whether to persist cookies or not. All apps should be selected for this session control to work correctly.",
				Transform:   transform.FromField("Description.PersistentBrowser")},
			{
				Name:        "platforms",
				Type:        proto.ColumnType_JSON,
				Description: "Platforms included in and excluded from the policy.",
				Transform:   transform.FromField("Description.Platforms")},
			{
				Name:        "sign_in_frequency",
				Type:        proto.ColumnType_JSON,
				Description: "Session control to enforce signin frequency.",
				Transform:   transform.FromField("Description.SignInFrequency")},
			{
				Name:        "sign_in_risk_levels",
				Type:        proto.ColumnType_JSON,
				Description: "Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue.",
				Transform:   transform.FromField("Description.SignInRiskLevels")},
			{
				Name:        "terms_of_use",
				Type:        proto.ColumnType_JSON,
				Description: "List of terms of use IDs required by the policy.",
				Transform:   transform.FromField("Description.TermsOfUse")},
			{
				Name:        "users",
				Type:        proto.ColumnType_JSON,
				Description: "Users, groups, and roles included in and excluded from the policy.",
				Transform:   transform.FromField("Description.Users")},
			{
				Name:        "user_risk_levels",
				Type:        proto.ColumnType_JSON,
				Description: "User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue.",
				Transform:   transform.FromField("Description.UserRiskLevel")},

			// Standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("Description.DisplayName")},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTenant,
				Transform:   transform.FromField("Description.TenantID")},
		}),
	}
}
