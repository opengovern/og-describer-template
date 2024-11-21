# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>Unique ID representing the sign-in activity.</td></tr>
	<tr><td>created_date_time</td><td>Date and time (UTC) the sign-in was initiated.</td></tr>
	<tr><td>user_display_name</td><td>Display name of the user that initiated the sign-in.</td></tr>
	<tr><td>user_principal_name</td><td>User principal name of the user that initiated the sign-in.</td></tr>
	<tr><td>user_id</td><td>ID of the user that initiated the sign-in.</td></tr>
	<tr><td>app_id</td><td>Unique GUID representing the app ID in the Azure Active Directory.</td></tr>
	<tr><td>app_display_name</td><td>App name displayed in the Azure Portal.</td></tr>
	<tr><td>ip_address</td><td>IP address of the client used to sign in.</td></tr>
	<tr><td>client_app_used</td><td>Identifies the legacy client used for sign-in activity.</td></tr>
	<tr><td>correlation_id</td><td>The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity.</td></tr>
	<tr><td>conditional_access_status</td><td>Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue.</td></tr>
	<tr><td>is_interactive</td><td>Indicates if a sign-in is interactive or not.</td></tr>
	<tr><td>risk_detail</td><td>Provides the &#39;reason&#39; behind a specific state of a risky user, sign-in or a risk event. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, unknownFutureValue.</td></tr>
	<tr><td>risk_level_aggregated</td><td>Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue.</td></tr>
	<tr><td>risk_level_during_sign_in</td><td>Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue.</td></tr>
	<tr><td>risk_state</td><td>Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.</td></tr>
	<tr><td>resource_display_name</td><td>Name of the resource the user signed into.</td></tr>
	<tr><td>resource_id</td><td>ID of the resource that the user signed into.</td></tr>
	<tr><td>risk_event_types</td><td>Risk event types associated with the sign-in. The possible values are: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence, generic, and unknownFutureValue.</td></tr>
	<tr><td>status</td><td>Sign-in status. Includes the error code and description of the error (in case of a sign-in failure).</td></tr>
	<tr><td>device_detail</td><td>Device information from where the sign-in occurred; includes device ID, operating system, and browser.</td></tr>
	<tr><td>location</td><td>Provides the city, state, and country code where the sign-in originated.</td></tr>
	<tr><td>applied_conditional_access_policies</td><td>Provides a list of conditional access policies that are triggered by the corresponding sign-in activity.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
</table>