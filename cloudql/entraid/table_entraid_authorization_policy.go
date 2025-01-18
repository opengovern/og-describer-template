package entraid

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEntraIdAuthorizationPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_authorization_policy",
		Description: "Represents a policy that can control Azure Active Directory authorization settings.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdAuthorizationPolicy,
		},

		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "Display name for this policy.",
				Transform:   transform.FromField("Description.DisplayName")},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "ID of the authorization policy.",
				Transform:   transform.FromField("Description.Id")},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Description of this policy.",
				Transform:   transform.FromField("Description.Description")},

			// Other fields
			{
				Name:        "allowed_to_sign_up_email_based_subscriptions",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether users can sign up for email based subscriptions.",
				Transform:   transform.FromField("Description.AllowedToSignIpEmailBasedSubscriptions")},
			{
				Name:        "allowed_to_use_sspr",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant.",
				Transform:   transform.FromField("Description.AllowedToUseSspr")},
			{
				Name:        "allowed_email_verified_users_to_join_organization",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether a user can join the tenant by email validation.",
				Transform:   transform.FromField("Description.AllowedEmailVerifiedUsersToJoinOrganization")},
			{
				Name:        "allow_invites_from",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone.",
				Transform:   transform.FromField("Description.AllowInvitesFrom")},
			{
				Name:        "block_msol_powershell",
				Type:        proto.ColumnType_BOOL,
				Description: "To disable the use of MSOL PowerShell set this property to true. This will also disable user-based access to the legacy service endpoint used by MSOL PowerShell. This does not affect Azure AD Connect or Microsoft Graph.",
				Transform:   transform.FromField("Description.BlockMsolPowershell")},
			{
				Name:        "guest_user_role_id",
				Type:        proto.ColumnType_STRING,
				Description: "Represents role templateId for the role that should be granted to guest user.",
				Transform:   transform.FromField("Description.GuestUserRoleId")},

			// JSON fields
			{
				Name:        "default_user_role_permissions",
				Type:        proto.ColumnType_JSON,
				Description: "Specifies certain customizable permissions for default user role.",
				Transform:   transform.FromField("Description.DefaultUserRolePermissions")},

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
