package azure

import (
	"context"
	"fmt"
	"strings"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureAppConfiguration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_app_configuration",
		Description: "Azure App Configuration",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetAppConfiguration,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAppConfiguration,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationStore.Name")},
			{
				Name:        "id",
				Description: "The resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationStore.ID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the configuration store. Possible values include: 'Creating', 'Updating', 'Deleting', 'Succeeded', 'Failed', 'Canceled'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationStore.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationStore.Type")},
			{
				Name:        "creation_date",
				Description: "The creation date of configuration store.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.ConfigurationStore.Properties.CreationDate").Transform(convertDateToTime),
			},
			{
				Name:        "endpoint",
				Description: "The DNS endpoint where the configuration store API will be available.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationStore.Properties.Endpoint")},
			{
				Name:        "public_network_access",
				Description: "Control permission for data plane traffic coming from public networks while private endpoint is enabled. Possible values include: 'Enabled', 'Disabled'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(publicNetworkAccess)},
			{
				Name:        "sku_name",
				Description: "The SKU name of the configuration store.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationStore.SKU.Name")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the configuration store.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "encryption",
				Description: "The encryption settings of the configuration store.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ConfigurationStore.Properties.Encryption")},
			{
				Name:        "identity",
				Description: "The managed identity information, if configured.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ConfigurationStore.Identity")},
			{
				Name:        "private_endpoint_connections",
				Description: "The list of private endpoint connections that are set up for this resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractAppConfigurationPrivateEndpointConnections),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationStore.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ConfigurationStore.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.ConfigurationStore.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ConfigurationStore.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				//// HYDRATE FUNCTIONS
				Transform: transform.

					//// TRANSFORM FUNCTION
					FromField("Description.ResourceGroup")},
		}),
	}
}

func publicNetworkAccess(_ context.Context, d *transform.TransformData) (interface{}, error) {
	server := d.HydrateItem.(opengovernance.AppConfiguration).Description.ConfigurationStore
	if server.Properties.PublicNetworkAccess != nil {
		return strings.ToLower(fmt.Sprintf("%v", *server.Properties.PublicNetworkAccess)), nil
	} else {
		return nil, nil
	}
}

// If we return the API response directly, the output will not provide all the properties of PrivateEndpointConnections
func extractAppConfigurationPrivateEndpointConnections(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	server := d.HydrateItem.(opengovernance.AppConfiguration).Description.ConfigurationStore
	var properties []map[string]interface{}

	if server.Properties.PrivateEndpointConnections != nil {
		for _, i := range server.Properties.PrivateEndpointConnections {
			objectMap := make(map[string]interface{})
			if i.ID != nil {
				objectMap["id"] = i.ID
			}
			if i.ID != nil {
				objectMap["name"] = i.Name
			}
			if i.ID != nil {
				objectMap["type"] = i.Type
			}
			if i.Properties != nil {
				if i.Properties.PrivateEndpoint != nil {
					objectMap["privateEndpointPropertyId"] = i.Properties.PrivateEndpoint.ID
				}
				if i.Properties.PrivateLinkServiceConnectionState != nil {
					if i.Properties.PrivateLinkServiceConnectionState.ActionsRequired != nil {
						if len(*i.Properties.PrivateLinkServiceConnectionState.ActionsRequired) > 0 {
							objectMap["privateLinkServiceConnectionStateActionsRequired"] = i.Properties.PrivateLinkServiceConnectionState.ActionsRequired
						}
					}
					if i.Properties.PrivateLinkServiceConnectionState.Status != nil {
						if len(*i.Properties.PrivateLinkServiceConnectionState.Status) > 0 {
							objectMap["privateLinkServiceConnectionStateStatus"] = i.Properties.PrivateLinkServiceConnectionState.Status
						}
					}
					if i.Properties.PrivateLinkServiceConnectionState.Description != nil {
						objectMap["privateLinkServiceConnectionStateDescription"] = i.Properties.PrivateLinkServiceConnectionState.Description
					}
				}
				if i.Properties.ProvisioningState != nil {
					if len(*i.Properties.ProvisioningState) > 0 {
						objectMap["provisioningState"] = i.Properties.ProvisioningState
					}
				}
			}
			properties = append(properties, objectMap)
		}
	}

	return properties, nil
}
