# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
	<tr><td>display_name</td><td>The name displayed in the address book for the user. This is usually the combination of the user&#39;s first name, middle initial and last name.</td></tr>
	<tr><td>id</td><td>The unique identifier for the group.</td></tr>
	<tr><td>description</td><td>An optional description for the group.</td></tr>
	<tr><td>classification</td><td>Describes a classification for the group (such as low, medium or high business impact).</td></tr>
	<tr><td>created_date_time</td><td>The time at which the group was created.</td></tr>
	<tr><td>expiration_date_time</td><td>Timestamp of when the group is set to expire.</td></tr>
	<tr><td>is_assignable_to_role</td><td>Indicates whether this group can be assigned to an Azure Active Directory role or not.</td></tr>
	<tr><td>is_subscribed_by_mail</td><td>Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true.</td></tr>
	<tr><td>mail</td><td>The SMTP address for the group, for example, &#34;serviceadmins@contoso.onmicrosoft.com&#34;.</td></tr>
	<tr><td>mail_enabled</td><td>Specifies whether the group is mail-enabled.</td></tr>
	<tr><td>mail_nickname</td><td>The mail alias for the user.</td></tr>
	<tr><td>membership_rule</td><td>The mail alias for the group, unique in the organization.</td></tr>
	<tr><td>membership_rule_processing_state</td><td>Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused.</td></tr>
	<tr><td>on_premises_domain_name</td><td>Contains the on-premises Domain name synchronized from the on-premises directory.</td></tr>
	<tr><td>on_premises_last_sync_date_time</td><td>Indicates the last time at which the group was synced with the on-premises directory.</td></tr>
	<tr><td>on_premises_net_bios_name</td><td>Contains the on-premises NetBiosName synchronized from the on-premises directory.</td></tr>
	<tr><td>on_premises_sam_account_name</td><td>Contains the on-premises SAM account name synchronized from the on-premises directory.</td></tr>
	<tr><td>on_premises_security_identifier</td><td>Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud.</td></tr>
	<tr><td>on_premises_sync_enabled</td><td>True if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default).</td></tr>
	<tr><td>renewed_date_time</td><td>Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action.</td></tr>
	<tr><td>security_enabled</td><td>Specifies whether the group is a security group.</td></tr>
	<tr><td>security_identifier</td><td>Security identifier of the group, used in Windows scenarios.</td></tr>
	<tr><td>visibility</td><td>Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership.</td></tr>
	<tr><td>assigned_labels</td><td>The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group.</td></tr>
	<tr><td>group_types</td><td>Specifies the group type and its membership. If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it&#39;s either a security group or distribution group. For details, see [groups overview](https://docs.microsoft.com/en-us/graph/api/resources/groups-overview?view=graph-rest-1.0).</td></tr>
	<tr><td>member_ids</td><td>Id of Users and groups that are members of this group.</td></tr>
	<tr><td>owner_ids</td><td>Id od the owners of the group. The owners are a set of non-admin users who are allowed to modify this object.</td></tr>
	<tr><td>proxy_addresses</td><td>Email addresses for the group that direct to the same group mailbox. For example: [&#34;SMTP: bob@contoso.com&#34;, &#34;smtp: bob@sales.contoso.com&#34;]. The any operator is required to filter expressions on multi-valued properties.</td></tr>
	<tr><td>resource_behavior_options</td><td>Specifies the group behaviors that can be set for a Microsoft 365 group during creation. Possible values are AllowOnlyMembersToPost, HideGroupInOutlook, SubscribeNewGroupMembers, WelcomeEmailDisabled.</td></tr>
	<tr><td>resource_provisioning_options</td><td>Specifies the group resources that are provisioned as part of Microsoft 365 group creation, that are not normally part of default group creation. Possible value is Team.</td></tr>
	<tr><td>tags</td><td>A map of tags for the resource.</td></tr>
	<tr><td>title</td><td>Title of the resource.</td></tr>
	<tr><td>tenant_id</td><td>The Azure Tenant ID where the resource is located.</td></tr>
</table>