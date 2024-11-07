package azuread

import (
	"context"

	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureAdMicrosoftApplication(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuread_microsoft_application",
		Description: "Represents an Azure Microsoft Application.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdMicrosoftApplication,
		},

		Columns: azureKaytuColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the service principal.", Transform: transform.FromField("Description.Id")},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name for the service principal.", Transform: transform.FromField("Description.DisplayName")},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the associated application (its appId property).", Transform: transform.FromField("Description.AppId")},
			{Name: "account_enabled", Type: proto.ColumnType_BOOL, Description: "true if the service principal account is enabled; otherwise, false.", Transform: transform.FromField("Description.AccountEnabled")},
			{Name: "app_display_name", Type: proto.ColumnType_STRING, Description: "The display name exposed by the associated application.", Transform: transform.FromField("Description.AppDisplayName")},
			{Name: "app_owner_organization_id", Type: proto.ColumnType_STRING, Description: "Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications.", Transform: transform.FromField("Description.AppOwnerOrganizationId")},
			{Name: "app_role_assignment_required", Type: proto.ColumnType_BOOL, Description: "Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false.", Transform: transform.FromField("Description.AppRoleAssignmentRequired")},
			{Name: "service_principal_type", Type: proto.ColumnType_STRING, Description: "Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Azure AD internally.", Transform: transform.FromField("Description.ServicePrincipalType")},
			{Name: "sign_in_audience", Type: proto.ColumnType_STRING, Description: "Specifies the Microsoft accounts that are supported for the current application. Supported values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount.", Transform: transform.FromField("Description.SignInAudience")},
			{Name: "app_description", Type: proto.ColumnType_STRING, Description: "The description exposed by the associated application.", Transform: transform.FromField("Description.AppDescription")},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Free text field to provide an internal end-user facing description of the service principal.", Transform: transform.FromField("Description.Description")},
			{Name: "login_url", Type: proto.ColumnType_STRING, Description: "Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.", Transform: transform.FromField("Description.LoginUrl")},
			{Name: "logout_url", Type: proto.ColumnType_STRING, Description: "Specifies the URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols.", Transform: transform.FromField("Description.LogoutUrl")},
			{Name: "add_ins", Type: proto.ColumnType_JSON, Description: "Defines custom behavior that a consuming service can use to call an app in specific contexts.", Transform: transform.FromField("Description.AddIns")},
			{Name: "alternative_names", Type: proto.ColumnType_JSON, Description: "Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.", Transform: transform.FromField("Description.AlternativeNames")},
			{Name: "app_roles", Type: proto.ColumnType_JSON, Description: "The roles exposed by the application which this service principal represents.", Transform: transform.FromField("Description.AppRoles")},
			{Name: "info", Type: proto.ColumnType_JSON, Description: "Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy statement URLs.", Transform: transform.FromField("Description.Info")},
			{Name: "key_credentials", Type: proto.ColumnType_JSON, Description: "The collection of key credentials associated with the service principal.", Transform: transform.FromField("Description.KeyCredentials")},
			{Name: "notification_email_addresses", Type: proto.ColumnType_JSON, Description: "Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.", Transform: transform.FromField("Description.NotificationEmailAddresses")},
			{Name: "owner_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Owners"), Description: "Id of the owners of the application. The owners are a set of non-admin users who are allowed to modify this object."},
			{Name: "password_credentials", Type: proto.ColumnType_JSON, Description: "Represents a password credential associated with a service principal.", Transform: transform.FromField("Description.PasswordCredentials")},
			{Name: "oauth2_permission_scopes", Type: proto.ColumnType_JSON, Description: "The published permission scopes.", Transform: transform.FromField("Description.PublishedPermissionScopes")},
			{Name: "reply_urls", Type: proto.ColumnType_JSON, Description: "The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application.", Transform: transform.FromField("Description.ReplyUrls")},
			{Name: "service_principal_names", Type: proto.ColumnType_JSON, Description: "Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD.", Transform: transform.FromField("Description.ServicePrincipalNames")},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Description: "Custom strings that can be used to categorize and identify the service principal.", Transform: transform.FromField("Description.Tags")},

			{Name: "tags", Type: proto.ColumnType_JSON, Description: ColumnDescriptionTags, Transform: transform.From(adMicrosoftApplicationTags)},
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.From(adMicrosoftApplicationTitle)},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Transform: transform.FromField("Description.TenantID")},
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

func adMicrosoftApplicationTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	servicePrincipal := d.HydrateItem.(opengovernance.AdMicrosoftApplication).Description
	tags := servicePrincipal.TagsSrc
	if tags == nil {
		return nil, nil
	}
	return TagsToMap(tags)
}

func adMicrosoftApplicationTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.AdMicrosoftApplication).Description

	title := data.DisplayName
	if title == nil {
		title = data.Id
	}

	return title, nil
}
