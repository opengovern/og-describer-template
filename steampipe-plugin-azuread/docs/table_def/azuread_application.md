# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>display_name</td><td>The display name for the application.</td></tr>
	<tr><td>id</td><td>The unique identifier for the application.</td></tr>
	<tr><td>app_id</td><td>The unique identifier for the application that is assigned to an application by Azure AD.</td></tr>
	<tr><td>created_date_time</td><td>The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time.</td></tr>
	<tr><td>description</td><td>Free text field to provide a description of the application object to end users.</td></tr>
	<tr><td>is_authorization_service_enabled</td><td>Is authorization service enabled.</td></tr>
	<tr><td>oauth2_require_post_response</td><td>Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. The default is false, which specifies that only GET requests are allowed.</td></tr>
	<tr><td>publisher_domain</td><td>The verified publisher domain for the application.</td></tr>
	<tr><td>sign_in_audience</td><td>Specifies the Microsoft accounts that are supported for the current application.</td></tr>
	<tr><td>api</td><td>Specifies settings for an application that implements a web API.</td></tr>
	<tr><td>identifier_uris</td><td>The URIs that identify the application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.</td></tr>
	<tr><td>info</td><td>Basic profile information of the application such as app&#39;s marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience.</td></tr>
	<tr><td>key_credentials</td><td>The collection of key credentials associated with the application.</td></tr>
	<tr><td>owner_ids</td><td>Id of the owners of the application. The owners are a set of non-admin users who are allowed to modify this object.</td></tr>
	<tr><td>parental_control_settings</td><td>Specifies parental control settings for an application.</td></tr>
	<tr><td>password_credentials</td><td>The collection of password credentials associated with the application.</td></tr>
	<tr><td>spa</td><td>Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.</td></tr>
	<tr><td>tags_src</td><td>Custom strings that can be used to categorize and identify the application.</td></tr>
	<tr><td>web</td><td>Specifies settings for a web application.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
</table>