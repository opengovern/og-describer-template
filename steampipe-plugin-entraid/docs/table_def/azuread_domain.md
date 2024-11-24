# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The fully qualified name of the domain.</td></tr>
	<tr><td>authentication_type</td><td>Indicates the configured authentication type for the domain. The value is either Managed or Federated. Managed indicates a cloud managed domain where Azure AD performs user authentication. Federated indicates authentication is federated with an identity provider such as the tenant&#39;s on-premises Active Directory via Active Directory Federation Services.</td></tr>
	<tr><td>is_default</td><td>true if this is the default domain that is used for user creation. There is only one default domain per company.</td></tr>
	<tr><td>is_admin_managed</td><td>The value of the property is false if the DNS record management of the domain has been delegated to Microsoft 365. Otherwise, the value is true.</td></tr>
	<tr><td>is_initial</td><td>true if this is the initial domain created by Microsoft Online Services (companyname.onmicrosoft.com). There is only one initial domain per company.</td></tr>
	<tr><td>is_root</td><td>true if the domain is a verified root domain. Otherwise, false if the domain is a subdomain or unverified.</td></tr>
	<tr><td>is_verified</td><td>true if the domain has completed domain ownership verification.</td></tr>
	<tr><td>supported_services</td><td>The capabilities assigned to the domain. Can include 0, 1 or more of following values: Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune. The values which you can add/remove using Graph API include: Email, OfficeCommunicationsOnline, Yammer.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
</table>