package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureSQLServer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_sql_server",
		Description: "Azure SQL Server",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetSqlServer,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404", "InvalidApiVersionParameter"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSqlServer,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the SQL server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a SQL server uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.ID")},
			{
				Name:        "type",
				Description: "The resource type of the SQL server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Type")},
			{
				Name:        "state",
				Description: "The state of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.State")},
			{
				Name:        "kind",
				Description: "The Kind of sql server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Kind")},
			{
				Name:        "location",
				Description: "The resource location.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Location")},
			{
				Name:        "administrator_login",
				Description: "Specifies the username of the administrator for this server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.AdministratorLogin")},
			{
				Name:        "administrator_login_password",
				Description: "The administrator login password.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.AdministratorLoginPassword")},
			{
				Name:        "minimal_tls_version",
				Description: "Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.MinimalTLSVersion")},
			{
				Name:        "public_network_access",
				Description: "Whether or not public endpoint access is allowed for this server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.PublicNetworkAccess")},
			{
				Name:        "version",
				Description: "The version of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.Version")},
			{
				Name:        "fully_qualified_domain_name",
				Description: "The fully qualified domain name of the server.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Properties.FullyQualifiedDomainName")},
			{
				Name:        "server_audit_policy",
				Description: "Specifies the audit policy configuration for server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServerBlobAuditingPolicies")},
			{
				Name:        "server_security_alert_policy",
				Description: "Specifies the security alert policy configuration for server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServerSecurityAlertPolicies")},
			{
				Name:        "server_azure_ad_administrator",
				Description: "Specifies the active directory administrator.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServerAzureADAdministrators")},
			{
				Name:        "server_vulnerability_assessment",
				Description: "Specifies the server's vulnerability assessment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServerVulnerabilityAssessments")},
			{
				Name:        "firewall_rules",
				Description: "A list of firewall rules for this server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FirewallRules")},
			{
				Name:        "automatic_tuning",
				Description: "Automatic tuning setting for this server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AutomaticTuning")},
			{
				Name:        "encryption_protector",
				Description: "The server encryption protector.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EncryptionProtectors")},
			{
				Name:        "private_endpoint_connections",
				Description: "The private endpoint connections of the sql server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PrivateEndpointConnections")},
			{
				Name:        "tags_src",
				Description: "Specifies the set of tags attached to the server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Server.Tags"),
			},
			{
				Name:        "virtual_network_rules",
				Description: "A list of virtual network rules for this server.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.VirtualNetworkRules")},
			{
				Name:        "failover_groups",
				Description: "A list of failover groups for this server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FailoverGroups")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Server.Name"),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Server.Tags"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Server.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Server.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

type PrivateConnectionInfo struct {
	PrivateEndpointConnectionId                      string
	PrivateEndpointId                                string
	PrivateEndpointConnectionName                    string
	PrivateEndpointConnectionType                    string
	PrivateLinkServiceConnectionStateStatus          string
	PrivateLinkServiceConnectionStateDescription     string
	PrivateLinkServiceConnectionStateActionsRequired string
	ProvisioningState                                string
}
