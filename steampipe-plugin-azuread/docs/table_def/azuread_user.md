# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>display_name</td><td>The name displayed in the address book for the user. This is usually the combination of the user&#39;s first name, middle initial and last name.</td></tr>
	<tr><td>id</td><td>The unique identifier for the user. Should be treated as an opaque identifier.</td></tr>
	<tr><td>user_principal_name</td><td>Principal email of the active directory user.</td></tr>
	<tr><td>account_enabled</td><td>True if the account is enabled; otherwise, false.</td></tr>
	<tr><td>user_type</td><td>A string value that can be used to classify user types in your directory.</td></tr>
	<tr><td>given_name</td><td>The given name (first name) of the user.</td></tr>
	<tr><td>surname</td><td>Family name or last name of the active directory user.</td></tr>
	<tr><td>filter</td><td>Odata query to search for resources.</td></tr>
	<tr><td>on_premises_immutable_id</td><td>Used to associate an on-premises Active Directory user account with their Azure AD user object.</td></tr>
	<tr><td>created_date_time</td><td>The time at which the user was created.</td></tr>
	<tr><td>mail</td><td>The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com.</td></tr>
	<tr><td>mail_nickname</td><td>The mail alias for the user.</td></tr>
	<tr><td>password_policies</td><td>Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two may be specified together; for example: DisablePasswordExpiration, DisableStrongPassword.</td></tr>
	<tr><td>refresh_tokens_valid_from_date_time</td><td>Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).</td></tr>
	<tr><td>sign_in_sessions_valid_from_date_time</td><td>Any refresh tokens or sessions tokens (session cookies) issued before this time are invalid, and applications will get an error when using an invalid refresh or sessions token to acquire a delegated access token (to access APIs such as Microsoft Graph).</td></tr>
	<tr><td>usage_location</td><td>A two letter country code (ISO standard 3166), required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.</td></tr>
	<tr><td>member_of</td><td>A list the groups and directory roles that the user is a direct member of.</td></tr>
	<tr><td>additional_properties</td><td>A list of unmatched properties from the message are deserialized this collection.</td></tr>
	<tr><td>im_addresses</td><td>The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.</td></tr>
	<tr><td>other_mails</td><td>A list of additional email addresses for the user.</td></tr>
	<tr><td>password_profile</td><td>Specifies the password profile for the user. The profile contains the user’s password. This property is required when a user is created.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
	<tr><td>metadata</td><td>Metadata of the Azure resource</td></tr>
	<tr><td>og_account_id</td><td>The Platform Account ID in which the resource is located.</td></tr>
	<tr><td>og_resource_id</td><td>The unique ID of the resource in Kaytu.</td></tr>
</table>