# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The unique identifier for the service principal.</td></tr>
	<tr><td>display_name</td><td>The display name for the service principal.</td></tr>
	<tr><td>app_id</td><td>The unique identifier for the associated application (its appId property).</td></tr>
	<tr><td>account_enabled</td><td>true if the service principal account is enabled; otherwise, false.</td></tr>
	<tr><td>app_display_name</td><td>The display name exposed by the associated application.</td></tr>
	<tr><td>app_owner_organization_id</td><td>Contains the tenant id where the application is registered. This is applicable only to service principals backed by applications.</td></tr>
	<tr><td>app_role_assignment_required</td><td>Specifies whether users or other service principals need to be granted an app role assignment for this service principal before users can sign in or apps can get tokens. The default value is false.</td></tr>
	<tr><td>service_principal_type</td><td>Identifies whether the service principal represents an application, a managed identity, or a legacy application. This is set by Azure AD internally.</td></tr>
	<tr><td>sign_in_audience</td><td>Specifies the Microsoft accounts that are supported for the current application. Supported values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount.</td></tr>
	<tr><td>app_description</td><td>The description exposed by the associated application.</td></tr>
	<tr><td>description</td><td>Free text field to provide an internal end-user facing description of the service principal.</td></tr>
	<tr><td>login_url</td><td>Specifies the URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.</td></tr>
	<tr><td>logout_url</td><td>Specifies the URL that will be used by Microsoft&#39;s authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols.</td></tr>
	<tr><td>add_ins</td><td>Defines custom behavior that a consuming service can use to call an app in specific contexts.</td></tr>
	<tr><td>alternative_names</td><td>Used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.</td></tr>
	<tr><td>app_roles</td><td>The roles exposed by the application which this service principal represents.</td></tr>
	<tr><td>info</td><td>Basic profile information of the acquired application such as app&#39;s marketing, support, terms of service and privacy statement URLs.</td></tr>
	<tr><td>key_credentials</td><td>The collection of key credentials associated with the service principal.</td></tr>
	<tr><td>notification_email_addresses</td><td>Specifies the list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.</td></tr>
	<tr><td>owner_ids</td><td>Id of the owners of the application. The owners are a set of non-admin users who are allowed to modify this object.</td></tr>
	<tr><td>password_credentials</td><td>Represents a password credential associated with a service principal.</td></tr>
	<tr><td>oauth2_permission_scopes</td><td>The published permission scopes.</td></tr>
	<tr><td>reply_urls</td><td>The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to for the associated application.</td></tr>
	<tr><td>service_principal_names</td><td>Contains the list of identifiersUris, copied over from the associated application. Additional values can be added to hybrid applications. These values can be used to identify the permissions exposed by this app within Azure AD.</td></tr>
	<tr><td>tags_src</td><td>Custom strings that can be used to categorize and identify the service principal.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
</table>