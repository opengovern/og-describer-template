# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>display_name</td><td>Display name for this policy.</td></tr>
	<tr><td>id</td><td>ID of the authorization policy.</td></tr>
	<tr><td>description</td><td>Description of this policy.</td></tr>
	<tr><td>allowed_to_sign_up_email_based_subscriptions</td><td>Indicates whether users can sign up for email based subscriptions.</td></tr>
	<tr><td>allowed_to_use_sspr</td><td>Indicates whether the Self-Serve Password Reset feature can be used by users on the tenant.</td></tr>
	<tr><td>allowed_email_verified_users_to_join_organization</td><td>Indicates whether a user can join the tenant by email validation.</td></tr>
	<tr><td>allow_invites_from</td><td>Indicates who can invite external users to the organization. Possible values are: none, adminsAndGuestInviters, adminsGuestInvitersAndAllMembers, everyone.</td></tr>
	<tr><td>block_msol_powershell</td><td>To disable the use of MSOL PowerShell set this property to true. This will also disable user-based access to the legacy service endpoint used by MSOL PowerShell. This does not affect Azure AD Connect or Microsoft Graph.</td></tr>
	<tr><td>guest_user_role_id</td><td>Represents role templateId for the role that should be granted to guest user.</td></tr>
	<tr><td>default_user_role_permissions</td><td>Specifies certain customizable permissions for default user role.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
</table>