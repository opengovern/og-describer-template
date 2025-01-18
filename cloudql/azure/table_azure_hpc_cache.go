package azure

import (
	"context"
	"time"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureHPCCache(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_hpc_cache",
		Description: "Azure HPC Cache",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetHpcCache,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListHpcCache,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the cache.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cache.Name")},
			{
				Name:        "id",
				Description: "The resource ID of the cache.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cache.ID")},
			{
				Name:        "provisioning_state",
				Description: "ARM provisioning state. Possible values include: 'Succeeded', 'Failed', 'Cancelled', 'Creating', 'Deleting', 'Updating'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cache.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the cache.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cache.Type")},
			{
				Name:        "cache_size_gb",
				Description: "The size of the cache, in GB.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Cache.Properties.CacheSizeGB")},
			{
				Name:        "sku_name",
				Description: "The SKU for the cache.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cache.SKU.Name")},
			{
				Name:        "subnet",
				Description: "Subnet used for the cache.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cache.Properties.Subnet")},
			{
				Name:        "directory_services_settings",
				Description: "Specifies directory services settings of the cache.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cache.Properties.DirectoryServicesSettings")},
			{
				Name:        "encryption_settings",
				Description: "Specifies encryption settings of the cache.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cache.Properties.EncryptionSettings")},
			{
				Name:        "health",
				Description: "The health of the cache.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cache.Properties.Health")},
			{
				Name:        "identity",
				Description: "The identity of the cache, if configured.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cache.Identity")},
			{
				Name:        "mount_addresses",
				Description: "Array of IP addresses that can be used by clients mounting the cache.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cache.Properties.MountAddresses")},
			{
				Name:        "network_settings",
				Description: "Specifies network settings of the cache.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractHPCCacheNetworkSettings),
			},
			{
				Name:        "security_settings",
				Description: "Specifies security settings of the cache.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cache.Properties.SecuritySettings")},
			{
				Name:        "system_data",
				Description: "The system meta data relating to the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cache.SystemData")},
			{
				Name:        "upgrade_status",
				Description: "Upgrade status of the cache.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractHPCCacheUpgradeStatus),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cache.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Cache.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Cache.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Cache.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

type CacheNetworkSettingsInfo struct {
	Mtu              *int32
	UtilityAddresses []*string
	DNSServers       []*string
	DNSSearchDomain  *string
	NtpServer        *string
}

type CacheUpgradeStatusInfo struct {
	CurrentFirmwareVersion *string
	FirmwareUpdateStatus   interface{}
	FirmwareUpdateDeadline *time.Time
	LastFirmwareUpdate     *time.Time
	PendingFirmwareVersion *string
}

//// LIST FUNCTION

//// HYDRATE FUNCTION

// Handle empty name or resourceGroup

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

//// TRANSFORM FUNCTIONS

// If we return the API response directly, the output does not provide
// all the properties of NetworkSettings
func extractHPCCacheNetworkSettings(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	cache := d.HydrateItem.(opengovernance.HpcCache).Description.Cache
	var properties CacheNetworkSettingsInfo

	if cache.Properties.NetworkSettings != nil {
		if cache.Properties.NetworkSettings.Mtu != nil {
			properties.Mtu = cache.Properties.NetworkSettings.Mtu
		}
		if cache.Properties.NetworkSettings.UtilityAddresses != nil {
			properties.UtilityAddresses = cache.Properties.NetworkSettings.UtilityAddresses
		}
		if cache.Properties.NetworkSettings.DNSServers != nil {
			properties.DNSServers = cache.Properties.NetworkSettings.DNSServers
		}
		if cache.Properties.NetworkSettings.DNSSearchDomain != nil {
			properties.DNSSearchDomain = cache.Properties.NetworkSettings.DNSSearchDomain
		}
		if cache.Properties.NetworkSettings.NtpServer != nil {
			properties.NtpServer = cache.Properties.NetworkSettings.NtpServer
		}
	}

	return properties, nil
}

// If we return the API response directly, the output does not provide
// all the properties of UpgradeStatus
func extractHPCCacheUpgradeStatus(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	cache := d.HydrateItem.(opengovernance.HpcCache).Description.Cache
	var properties CacheUpgradeStatusInfo

	if cache.Properties.UpgradeStatus != nil {
		if cache.Properties.UpgradeStatus.CurrentFirmwareVersion != nil {
			properties.CurrentFirmwareVersion = cache.Properties.UpgradeStatus.CurrentFirmwareVersion
		}
		if cache.Properties.UpgradeStatus.FirmwareUpdateStatus != nil {
			if len(*cache.Properties.UpgradeStatus.FirmwareUpdateStatus) > 0 {
				properties.FirmwareUpdateStatus = cache.Properties.UpgradeStatus.FirmwareUpdateStatus
			}
		}
		if cache.Properties.UpgradeStatus.FirmwareUpdateDeadline != nil {
			properties.FirmwareUpdateDeadline = cache.Properties.UpgradeStatus.FirmwareUpdateDeadline
		}
		if cache.Properties.UpgradeStatus.LastFirmwareUpdate != nil {
			properties.LastFirmwareUpdate = cache.Properties.UpgradeStatus.LastFirmwareUpdate
		}
		if cache.Properties.UpgradeStatus.PendingFirmwareVersion != nil {
			properties.PendingFirmwareVersion = cache.Properties.UpgradeStatus.PendingFirmwareVersion
		}
	}

	return properties, nil
}
