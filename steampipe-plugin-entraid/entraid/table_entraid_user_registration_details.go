package entraid

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEntraIdUserRegistrationDetails(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_user_registration_details",
		Description: "Represents an Azure AD user registration details.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdUserRegistrationDetails,
		},

		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier for the user. Should be treated as an opaque identifier.",
				Transform:   transform.FromField("Description.Id"),
			},
			{
				Name:        "user_display_name",
				Type:        proto.ColumnType_STRING,
				Description: "User Display name.",
				Transform:   transform.FromField("Description.UserDisplayName"),
			},
			{
				Name:        "user_principal_name",
				Type:        proto.ColumnType_STRING,
				Description: "User Principal name.",
				Transform:   transform.FromField("Description.UserPrincipalName"),
			},
			{
				Name:        "user_type",
				Type:        proto.ColumnType_STRING,
				Description: "User Type.",
				Transform:   transform.FromField("Description.UserType"),
			},
			{
				Name:        "system_preferred_authentication_methods",
				Type:        proto.ColumnType_JSON,
				Description: "SystemPreferredAuthenticationMethods",
				Transform:   transform.FromField("Description.SystemPreferredAuthenticationMethods"),
			},
			{
				Name:        "is_system_preferred_authentication_method_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "IsSystemPreferredAuthenticationMethodEnabled",
				Transform:   transform.FromField("Description.IsSystemPreferredAuthenticationMethodEnabled"),
			},
			{
				Name:        "is_admin",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the user is admin or not.",
				Transform:   transform.FromField("Description.IsAdmin"),
			},
			{
				Name:        "is_mfa_capable",
				Type:        proto.ColumnType_BOOL,
				Description: "IsMfaCapable",
				Transform:   transform.FromField("Description.IsMfaCapable"),
			},
			{
				Name:        "is_mfa_registered",
				Type:        proto.ColumnType_BOOL,
				Description: "IsMfaRegistered",
				Transform:   transform.FromField("Description.IsMfaRegistered"),
			},
			{
				Name:        "is_sspr_capable",
				Type:        proto.ColumnType_BOOL,
				Description: "IsSsprCapable",
				Transform:   transform.FromField("Description.IsSsprCapable"),
			},
			{
				Name:        "is_sspr_registered",
				Type:        proto.ColumnType_BOOL,
				Description: "IsSsprRegistered",
				Transform:   transform.FromField("Description.IsSsprRegistered"),
			},
			{
				Name:        "is_sspr_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "IsSsprEnabled",
				Transform:   transform.FromField("Description.IsSsprEnabled"),
			},
			{
				Name:        "is_passwordless_capable",
				Type:        proto.ColumnType_BOOL,
				Description: "IsPasswordlessCapable",
				Transform:   transform.FromField("Description.IsPasswordlessCapable"),
			},
			{
				Name:        "last_updated_date_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "LastUpdatedDateTime",
				Transform:   transform.FromField("Description.LastUpdatedDateTime"),
			},
			{
				Name:        "methods_registered",
				Type:        proto.ColumnType_JSON,
				Description: "MethodsRegistered",
				Transform:   transform.FromField("Description.MethodsRegistered"),
			},
			{
				Name:        "user_preferred_method_for_secondary_authentication",
				Type:        proto.ColumnType_STRING,
				Description: "UserPreferredMethodForSecondaryAuthentication",
				Transform:   transform.FromField("Description.UserPreferredMethodForSecondaryAuthentication"),
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,

				Transform: transform.FromField("Description.Id")},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTenant,
				Transform:   transform.FromField("Description.TenantID")},
		}),
	}
}
