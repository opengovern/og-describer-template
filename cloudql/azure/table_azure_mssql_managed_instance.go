package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureMSSQLManagedInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_mssql_managed_instance",
		Description: "Azure Microsoft SQL Managed Instance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetMssqlManagedInstance,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404", "InvalidApiVersionParameter"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListMssqlManagedInstance,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a managed instance uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.ID")},
			{
				Name:        "type",
				Description: "The resource type of the managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Type")},
			{
				Name:        "state",
				Description: "The state of the managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.State")},
			{
				Name:        "administrator_login",
				Description: "Administrator username for the managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.AdministratorLogin")},
			{
				Name:        "administrator_login_password",
				Description: "Administrator password for the managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.AdministratorLoginPassword")},
			{
				Name:        "collation",
				Description: "Collation of the managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.Collation")},
			{
				Name:        "dns_zone",
				Description: "The Dns zone that the managed instance is in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.DNSZone")},
			{
				Name:        "dns_zone_partner",
				Description: "The resource id of another managed instance whose DNS zone this managed instance will share after creation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.DNSZonePartner")},
			{
				Name:        "fully_qualified_domain_name",
				Description: "The fully qualified domain name of the managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.FullyQualifiedDomainName")},
			{
				Name:        "instance_pool_id",
				Description: "The Id of the instance pool this managed server belongs to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.InstancePoolID")},
			{
				Name:        "license_type",
				Description: "The license type of the managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.LicenseType")},
			{
				Name:        "maintenance_configuration_id",
				Description: "Specifies maintenance configuration id to apply to this managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.MaintenanceConfigurationID")},
			{
				Name:        "managed_instance_create_mode",
				Description: "Specifies the mode of database creation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.ManagedInstanceCreateMode")},
			{
				Name:        "minimal_tls_version",
				Description: "Minimal TLS version of the managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.MinimalTLSVersion")},
			{
				Name:        "proxy_override",
				Description: "Connection type used for connecting to the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.ProxyOverride")},
			{
				Name:        "public_data_endpoint_enabled",
				Description: "Whether or not the public data endpoint is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.PublicDataEndpointEnabled")},
			{
				Name:        "restore_point_in_time",
				Description: "Specifies the point in time of the source database that will be restored to create the new database.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.ManagedInstance.Properties.RestorePointInTime").Transform(convertDateToTime),
			},
			{
				Name:        "source_managed_instance_id",
				Description: "The resource identifier of the source managed instance associated with create operation of this instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.SourceManagedInstanceID")},
			{
				Name:        "storage_size_in_gb",
				Description: "The managed instance storage size in GB.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.StorageSizeInGB")},
			{
				Name:        "subnet_id",
				Description: "Subnet resource ID for the managed instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.SubnetID")},
			{
				Name:        "timezone_id",
				Description: "Id of the timezone.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.TimezoneID")},
			{
				Name:        "v_cores",
				Description: "The number of vcores of the managed instance.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ManagedInstance.Properties.VCores")},
			{
				Name:        "encryption_protectors",
				Description: "The managed instance encryption protectors.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedInstanceEncryptionProtectors")},
			{
				Name:        "identity",
				Description: "The azure active directory identity of the managed instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedInstance.Identity")},
			{
				Name:        "security_alert_policies",
				Description: "The security alert policies of the managed instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedDatabaseSecurityAlertPolicies")},
			{
				Name:        "sku",
				Description: "Managed instance SKU.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedInstance.SKU")},
			{
				Name:        "vulnerability_assessments",
				Description: "The managed instance vulnerability assessments.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.ManagedInstanceVulnerabilityAssessments")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ManagedInstance.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedInstance.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.ManagedInstance.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ManagedInstance.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
