package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureCognitiveAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_cognitive_account",
		Description: "Azure Cognitive Account",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetCognitiveAccount,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCognitiveAccount,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Name")},
			{
				Name:        "id",
				Description: "Fully qualified resource ID for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.ID")},
			{
				Name:        "kind",
				Description: "The kind of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Kind")},
			{
				Name:        "provisioning_state",
				Description: "The status of the cognitive services account at the time the operation was called. Possible values include: 'Accepted', 'Creating', 'Deleting', 'Moving', 'Failed', 'Succeeded', 'ResolvingDNS'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the resource. E.g. 'Microsoft.Compute/virtualMachines' or 'Microsoft.Storage/storageAccounts'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Type")},
			{
				Name:        "custom_sub_domain_name",
				Description: "The subdomain name used for token-based authentication.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.CustomSubDomainName")},
			{
				Name:        "date_created",
				Description: "The date of cognitive services account creation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.DateCreated")},
			{
				Name:        "disable_local_auth",
				Description: "Checks if local auth is disabled for the resource.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.Account.Properties.DisableLocalAuth")},
			{
				Name:        "endpoint",
				Description: "The endpoint of the created account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.Endpoint")},
			{
				Name:        "etag",
				Description: "The resource etag.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Etag")},
			{
				Name:        "is_migrated",
				Description: "Checks if the resource is migrated from an existing key.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.Account.Properties.IsMigrated")},
			{
				Name:        "migration_token",
				Description: "The resource migration token.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.MigrationToken")},
			{
				Name:        "public_network_access",
				Description: "Whether or not public endpoint access is allowed for this account. Value is optional but if passed in, must be 'Enabled' or 'Disabled'. Possible values include: 'Enabled', 'Disabled'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.PublicNetworkAccess")},
			{
				Name:        "restore",
				Description: "Checks if restore is enabled for the resource.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.Account.Properties.Restore")},
			{
				Name:        "restrict_outbound_network_access",
				Description: "Checks if outbound network access is restricted for the resource.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.Account.Properties.RestrictOutboundNetworkAccess")},
			{
				Name:        "allowed_fqdn_list",
				Description: "The allowed FQDN list for the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.AllowedFqdnList")},
			{
				Name:        "api_properties",
				Description: "The api properties for special APIs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.APIProperties")},
			{
				Name:        "call_rate_limit",
				Description: "The call rate limit of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.CallRateLimit")},
			{
				Name:        "capabilities",
				Description: "The capabilities of the cognitive services account. Each item indicates the capability of a specific feature. The values are read-only and for reference only.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.Capabilities")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the cognitive service account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "encryption",
				Description: "The encryption properties for the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.Encryption")},
			{
				Name:        "endpoints",
				Description: "All endpoints of the cognitive services account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.Endpoints")},
			{
				Name:        "identity",
				Description: "The identity for the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Identity")},
			{
				Name:        "network_acls",
				Description: "A collection of rules governing the accessibility from specific network locations.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.NetworkACLs")},
			{
				Name:        "private_endpoint_connections",
				Description: "The private endpoint connection associated with the cognitive services account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractAccountPrivateEndpointConnections),
			},
			{
				Name:        "quota_limit",
				Description: "The quota limit of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.QuotaLimit")},
			{
				Name:        "sku",
				Description: "The resource model definition representing SKU.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.SKU")},
			{
				Name:        "sku_change_info",
				Description: "Sku change info of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.SKUChangeInfo")},
			{
				Name:        "system_data",
				Description: "The metadata pertaining to creation and last modification of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.SystemData")},
			{
				Name:        "user_owned_storage",
				Description: "The storage accounts for the resource.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Account.Properties.UserOwnedStorage")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Account.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

type CognitiveAccountPrivateEndpointConnections struct {
	PrivateEndpointID                 interface{}
	PrivateLinkServiceConnectionState interface{}
	ProvisioningState                 interface{}
	GroupIds                          []*string
	SystemData                        interface{}
	Location                          *string
	Etag                              *string
	ID                                *string
	Name                              *string
	Type                              *string
}

// If we return the API response directly, the output will not provide all the properties of PrivateEndpointConnections
func extractAccountPrivateEndpointConnections(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	account := d.HydrateItem.(opengovernance.CognitiveAccount).Description.Account
	privateEndpointConnectionInfo := []CognitiveAccountPrivateEndpointConnections{}

	if account.Properties.PrivateEndpointConnections != nil {
		for _, connection := range account.Properties.PrivateEndpointConnections {
			properties := CognitiveAccountPrivateEndpointConnections{}
			properties.SystemData = connection.SystemData
			properties.Location = connection.Location
			properties.Etag = connection.Etag
			properties.ID = connection.ID
			properties.Name = connection.Name
			properties.Type = connection.Type
			if connection.Properties != nil {
				if connection.Properties.PrivateEndpoint != nil {
					properties.PrivateEndpointID = connection.Properties.PrivateEndpoint.ID
				}
				properties.PrivateLinkServiceConnectionState = connection.Properties.PrivateLinkServiceConnectionState
				properties.ProvisioningState = connection.Properties.ProvisioningState
				properties.GroupIds = connection.Properties.GroupIDs
			}
			privateEndpointConnectionInfo = append(privateEndpointConnectionInfo, properties)
		}
	}

	return privateEndpointConnectionInfo, nil
}
