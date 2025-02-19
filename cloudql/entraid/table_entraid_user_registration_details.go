package entraid

import (
        "context"

        opengovernance "github.com/opengovern/og-describer-entraid/discovery/pkg/es"
        "github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
        "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
        "github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEntraIdUserRegistrationDetails(_ context.Context) *plugin.Table {
        return &plugin.Table{
                Name:        "entraid_user_registration_details",
                Description: "Stores detailed secondary registration information for users within Microsoft Entra ID.",
                List: &plugin.ListConfig{
                        Hydrate: opengovernance.ListAdUserRegistrationDetails,
                },

                Columns: azureOGColumns([]*plugin.Column{
                        {
                                Name:        "user_object_id",
                                Type:        proto.ColumnType_STRING,
                                Description: "The unique, opaque identifier for the Entra ID user. References id in the entraid_user table.",
                                Transform:   transform.FromField("Description.Id"),
                        },
                        {
                                Name:        "registration_system_preferred_authentication_methods",
                                Type:        proto.ColumnType_JSON,
                                Description: "System-preferred authentication methods registered for the user.",
                                Transform:   transform.FromField("Description.SystemPreferredAuthenticationMethods"),
                        },
                        {
                                Name:        "registration_is_system_preferred_authentication_method_enabled",
                                Type:        proto.ColumnType_BOOL,
                                Description: "Indicates if the system-preferred authentication method is enabled for the user.",
                                Transform:   transform.FromField("Description.IsSystemPreferredAuthenticationMethodEnabled"),
                        },
                        {
                                Name:        "registration_has_admin_privileges_assigned",
                                Type:        proto.ColumnType_BOOL,
                                Description: "Indicates whether the user has administrative privileges assigned during registration.",
                                Transform:   transform.FromField("Description.IsAdmin"),
                        },
                        {
                                Name:        "registration_is_mfa_capable",
                                Type:        proto.ColumnType_BOOL,
                                Description: "Indicates if the user is capable of multi-factor authentication (MFA) during registration.",
                                Transform:   transform.FromField("Description.IsMfaCapable"),
                        },
                        {
                                Name:        "registration_is_mfa_registered",
                                Type:        proto.ColumnType_BOOL,
                                Description: "Indicates if the user has registered for MFA during registration.",
                                Transform:   transform.FromField("Description.IsMfaRegistered"),
                        },
                        {
                                Name:        "registration_is_sspr_capable",
                                Type:        proto.ColumnType_BOOL,
                                Description: "Indicates if the user is capable of self-service password reset (SSPR) during registration.",
                                Transform:   transform.FromField("Description.IsSsprCapable"),
                        },
                        {
                                Name:        "registration_is_sspr_registered",
                                Type:        proto.ColumnType_BOOL,
                                Description: "Indicates if the user has registered for SSPR during registration.",
                                Transform:   transform.FromField("Description.IsSsprRegistered"),
                        },
                        {
                                Name:        "registration_is_sspr_enabled",
                                Type:        proto.ColumnType_BOOL,
                                Description: "Indicates if SSPR is enabled for the user during registration.",
                                Transform:   transform.FromField("Description.IsSsprEnabled"),
                        },
                        {
                                Name:        "registration_is_passwordless_capable",
                                Type:        proto.ColumnType_BOOL,
                                Description: "Indicates if the user is capable of passwordless authentication during registration.",
                                Transform:   transform.FromField("Description.IsPasswordlessCapable"),
                        },
                        {
                                Name:        "registration_last_updated_date_time",
                                Type:        proto.ColumnType_TIMESTAMP,
                                Description: "The timestamp of the last update to the user's registration details.",
                                Transform:   transform.FromField("Description.LastUpdatedDateTime"),
                        },
                        {
                                Name:        "registration_methods_registered",
                                Type:        proto.ColumnType_JSON,
                                Description: "The authentication methods registered by the user during registration.",
                                Transform:   transform.FromField("Description.MethodsRegistered"),
                        },
                        {
                                Name:        "registration_user_preferred_method_for_secondary_authentication",
                                Type:        proto.ColumnType_STRING,
                                Description: "The user's preferred method for secondary authentication during registration.",
                                Transform:   transform.FromField("Description.UserPreferredMethodForSecondaryAuthentication"),
                        },
                        {
                                Name:        "tenant_id",
                                Type:        proto.ColumnType_STRING,
                                Description: ColumnDescriptionTenant,
                                Transform:   transform.FromField("Description.TenantID"),
                        },
                }),
        }
}
