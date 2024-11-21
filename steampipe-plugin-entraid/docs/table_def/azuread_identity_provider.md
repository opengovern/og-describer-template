# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>id</td><td>The ID of the identity provider.</td></tr>
	<tr><td>name</td><td>The display name of the identity provider.</td></tr>
	<tr><td>type</td><td>The identity provider type is a required field. For B2B scenario: Google, Facebook. For B2C scenario: Microsoft, Google, Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo, QQ, WeChat, OpenIDConnect.</td></tr>
	<tr><td>client_id</td><td>The client ID for the application. This is the client ID obtained when registering the application with the identity provider.</td></tr>
	<tr><td>client_secret</td><td>The client secret for the application. This is the client secret obtained when registering the application with the identity provider. This is write-only. A read operation will return ****.</td></tr>
	<tr><td>filter</td><td>Odata query to search for resources.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
</table>