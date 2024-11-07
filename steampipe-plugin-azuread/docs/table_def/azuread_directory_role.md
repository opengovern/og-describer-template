# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The unique identifier for the directory role.</td></tr>
	<tr><td>description</td><td>The description for the directory role.</td></tr>
	<tr><td>display_name</td><td>The display name for the directory role.</td></tr>
	<tr><td>role_template_id</td><td>The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only.</td></tr>
	<tr><td>member_ids</td><td>Id of the owners of the application. The owners are a set of non-admin users who are allowed to modify this object.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
</table>