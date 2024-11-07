# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>Indicates the unique ID for the activity.</td></tr>
	<tr><td>activity_date_time</td><td>Indicates the date and time the activity was performed.</td></tr>
	<tr><td>activity_display_name</td><td>Indicates the activity name or the operation name.</td></tr>
	<tr><td>category</td><td>Indicates which resource category that&#39;s targeted by the activity.</td></tr>
	<tr><td>correlation_id</td><td>Indicates a unique ID that helps correlate activities that span across various services. Can be used to trace logs across services.</td></tr>
	<tr><td>logged_by_service</td><td>Indicates information on which service initiated the activity (For example: Self-service Password Management, Core Directory, B2C, Invited Users, Microsoft Identity Manager, Privileged Identity Management.</td></tr>
	<tr><td>operation_type</td><td>Indicates the type of operation that was performed. The possible values include but are not limited to the following: Add, Assign, Update, Unassign, and Delete.</td></tr>
	<tr><td>result</td><td>Indicates the result of the activity. Possible values are: success, failure, timeout, unknownFutureValue.</td></tr>
	<tr><td>result_reason</td><td>Indicates the reason for failure if the result is failure or timeout.</td></tr>
	<tr><td>additional_details</td><td>Indicates additional details on the activity.</td></tr>
	<tr><td>initiated_by</td><td>Indicates information about the user or app initiated the activity.</td></tr>
	<tr><td>target_resources</td><td>Indicates information on which resource was changed due to the activity. Target Resource Type can be User, Device, Directory, App, Role, Group, Policy or Other.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
	<tr><td>filter</td><td>Odata query to search for directory audit reports.</td></tr>
</table>