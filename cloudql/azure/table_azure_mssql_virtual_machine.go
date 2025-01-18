package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureMSSQLVirtualMachine(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_mssql_virtual_machine",
		Description: "Azure MS SQL Virtual Machine",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetSqlServerVirtualMachine,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSqlServerVirtualMachine,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Name")},
			{
				Name:        "id",
				Description: "The resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.ID")},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state to track the async operation status.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Type")},
			{
				Name:        "sql_image_offer",
				Description: "SQL image offer for the SQL virtual machine.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.SQLImageOffer")},
			{
				Name:        "sql_image_sku",
				Description: "SQL Server edition type. Possible values include: 'Developer', 'Express', 'Standard', 'Enterprise', 'Web'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.SQLImageSKU")},
			{
				Name:        "sql_management",
				Description: "SQL Server Management type. Possible values include: 'Full', 'LightWeight', 'NoAgent'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.SQLManagement")},
			{
				Name:        "sql_server_license_type",
				Description: "SQL server license type for the SQL virtual machine. Possible values include: 'PAYG', 'AHUB', 'DR'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.SQLServerLicenseType")},
			{
				Name:        "sql_virtual_machine_group_resource_id",
				Description: "ARM resource id of the SQL virtual machine group this SQL virtual machine is or will be part of.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.SQLVirtualMachineGroupResourceID")},
			{
				Name:        "virtual_machine_resource_id",
				Description: "ARM resource id of underlying virtual machine created from SQL marketplace image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.VirtualMachineResourceID")},
			{
				Name:        "auto_backup_settings",
				Description: "Auto backup settings for SQL Server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.AutoBackupSettings")},
			{
				Name:        "auto_patching_settings",
				Description: "Auto patching settings for applying critical security updates to SQL virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.AutoPatchingSettings")},
			{
				Name:        "identity",
				Description: "Azure Active Directory identity for the SQL virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Identity")},
			{
				Name:        "key_vault_credential_settings",
				Description: "Key vault credential settings for the SQL virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.KeyVaultCredentialSettings")},
			{
				Name:        "server_configurations_management_settings",
				Description: "SQL server configuration management settings for the SQL virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.ServerConfigurationsManagementSettings")},
			{
				Name:        "storage_configuration_settings",
				Description: "Storage configuration settings for the SQL virtual machine.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.StorageConfigurationSettings")},
			{
				Name:        "wsfc_domain_credentials",
				Description: "Domain credentials for setting up Windows Server Failover Cluster for SQL availability group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Properties.WsfcDomainCredentials")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMachine.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.VirtualMachine.ID").Transform(idToAkas),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMachine.Tags")},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMachine.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
