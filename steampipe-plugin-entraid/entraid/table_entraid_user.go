package entraid

import (
	"context"
	"encoding/json"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEntraIdUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_user",
		Description: "Represents an Azure AD user account.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdUsers,
			KeyColumns: plugin.KeyColumnSlice{
				// Key fields
				{
					Name: "id", Require: plugin.Optional},
				{
					Name: "user_principal_name", Require: plugin.Optional},
				{
					Name: "user_type", Require: plugin.Optional},
				{
					Name: "account_enabled", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{
					Name: "display_name", Require: plugin.Optional},
			},
		},

		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name.",

				Transform: transform.FromField("Description.DisplayName")},
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier for the user. Should be treated as an opaque identifier.",
				Transform:   transform.FromField("Description.Id"),
			},
			{
				Name:        "user_principal_name",
				Type:        proto.ColumnType_STRING,
				Description: "Principal email of the active directory user.",

				Transform: transform.FromField("Description.UserPrincipalName")},
			{
				Name:        "account_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "True if the account is enabled; otherwise, false.",

				Transform: transform.FromField("Description.AccountEnabled")},
			{
				Name:        "created_date_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the user was created.",

				Transform: transform.FromField("Description.CreatedDateTime")},
			{
				Name:        "last_sign_in_date_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the user was last signed in.",

				Transform: transform.FromField("Description.LastSignInDateTime")},
			{
				Name:        "user_type",
				Type:        proto.ColumnType_STRING,
				Description: "A string value that can be used to classify user types in your directory.",

				Transform: transform.FromField("Description.UserType")},
			{
				Name:        "mail",
				Type:        proto.ColumnType_STRING,
				Description: "The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com.",

				Transform: transform.FromField("Description.Mail")},
			{
				Name:        "job_title",
				Type:        proto.ColumnType_STRING,
				Description: "The user job title.",

				Transform: transform.FromField("Description.JobTitle")},
			{
				Name:        "identities",
				Type:        proto.ColumnType_JSON,
				Description: "User identities",

				Transform: transform.FromField("Description.Identities")},
			{
				Name:        "password_policies",
				Type:        proto.ColumnType_STRING,
				Description: "Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword.",

				Transform: transform.FromField("Description.PasswordPolicies")},
			{
				Name:        "sign_in_sessions_valid_from_date_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).",

				Transform: transform.FromField("Description.SignInSessionsValidFromDateTime")},
			{
				Name:        "usage_location",
				Type:        proto.ColumnType_STRING,
				Description: "A two letter country code (ISO standard 3166), required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.",
				Transform:   transform.FromField("Description.UsageLocation")},
			{
				Name:        "im_addresses",
				Type:        proto.ColumnType_JSON,
				Description: "The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.",

				Transform: transform.FromField("Description.ImAddresses")},
			{
				Name:        "other_mails",
				Type:        proto.ColumnType_JSON,
				Description: "A list of additional email addresses for the user.",

				Transform: transform.FromField("Description.OtherMails")},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTenant,
				Transform:   transform.FromField("Description.TenantID")},
			{
				Name:        "metadata",
				Description: "Metadata of the Azure resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata").Transform(marshalJSON),
			},
			{
				Name:        "og_account_id",
				Type:        proto.ColumnType_STRING,
				Description: "The Platform Account ID in which the resource is located.",
				Transform:   transform.FromField("Metadata.SourceID")},
			{
				Name:        "og_resource_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique ID of the resource in opengovernance.",
				Transform:   transform.FromField("ID")},
		}),
	}
}

func marshalJSON(_ context.Context, d *transform.TransformData) (interface{}, error) {
	b, err := json.Marshal(d.Value)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}
