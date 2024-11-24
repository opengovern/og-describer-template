# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The unique identifier for the device. Inherited from directoryObject.</td></tr>
	<tr><td>display_name</td><td>The name displayed for the device.</td></tr>
	<tr><td>account_enabled</td><td>True if the account is enabled; otherwise, false.</td></tr>
	<tr><td>device_id</td><td>Unique identifier set by Azure Device Registration Service at the time of registration.</td></tr>
	<tr><td>approximate_last_sign_in_date_time</td><td>The timestamp type represents date and time information using ISO 8601 format and is always in UTC time.</td></tr>
	<tr><td>filter</td><td>Odata query to search for resources.</td></tr>
	<tr><td>is_compliant</td><td>True if the device is compliant; otherwise, false.</td></tr>
	<tr><td>is_managed</td><td>True if the device is managed; otherwise, false.</td></tr>
	<tr><td>mdm_app_id</td><td>Application identifier used to register device into MDM.</td></tr>
	<tr><td>operating_system</td><td>The type of operating system on the device.</td></tr>
	<tr><td>operating_system_version</td><td>The version of the operating system on the device.</td></tr>
	<tr><td>profile_type</td><td>A string value that can be used to classify device types.</td></tr>
	<tr><td>trust_type</td><td>Type of trust for the joined device. Possible values: Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD).</td></tr>
	<tr><td>extension_attributes</td><td>Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD.</td></tr>
	<tr><td>member_of</td><td>A list the groups and directory roles that the device is a direct member of.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
</table>