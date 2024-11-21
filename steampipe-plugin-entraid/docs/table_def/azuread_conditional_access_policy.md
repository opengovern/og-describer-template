# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>Specifies the identifier of a conditionalAccessPolicy object.</td></tr>
	<tr><td>display_name</td><td>Specifies a display name for the conditionalAccessPolicy object.</td></tr>
	<tr><td>state</td><td>Specifies the state of the conditionalAccessPolicy object. Possible values are: enabled, disabled, enabledForReportingButNotEnforced.</td></tr>
	<tr><td>created_date_time</td><td>The create date of the conditional access policy.</td></tr>
	<tr><td>modified_date_time</td><td>The modification date of the conditional access policy.</td></tr>
	<tr><td>operator</td><td>Defines the relationship of the grant controls. Possible values: AND, OR.</td></tr>
	<tr><td>applications</td><td>Applications and user actions included in and excluded from the policy.</td></tr>
	<tr><td>application_enforced_restrictions</td><td>Session control to enforce application restrictions. Only Exchange Online and Sharepoint Online support this session control.</td></tr>
	<tr><td>built_in_controls</td><td>List of values of built-in controls required by the policy. Possible values: block, mfa, compliantDevice, domainJoinedDevice, approvedApplication, compliantApplication, passwordChange, unknownFutureValue.</td></tr>
	<tr><td>client_app_types</td><td>Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other.</td></tr>
	<tr><td>custom_authentication_factors</td><td>List of custom controls IDs required by the policy.</td></tr>
	<tr><td>cloud_app_security</td><td>Session control to apply cloud app security.</td></tr>
	<tr><td>locations</td><td>Locations included in and excluded from the policy.</td></tr>
	<tr><td>persistent_browser</td><td>Session control to define whether to persist cookies or not. All apps should be selected for this session control to work correctly.</td></tr>
	<tr><td>platforms</td><td>Platforms included in and excluded from the policy.</td></tr>
	<tr><td>sign_in_frequency</td><td>Session control to enforce signin frequency.</td></tr>
	<tr><td>sign_in_risk_levels</td><td>Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue.</td></tr>
	<tr><td>terms_of_use</td><td>List of terms of use IDs required by the policy.</td></tr>
	<tr><td>users</td><td>Users, groups, and roles included in and excluded from the policy.</td></tr>
	<tr><td>user_risk_levels</td><td>User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
</table>