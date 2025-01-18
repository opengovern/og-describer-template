package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSignalRService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_signalr_service",
		Description: "Azure SignalR Service",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetSignalrService,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSignalrService,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Name")},
			{
				Name:        "id",
				Description: "Fully qualified resource ID for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.ID")},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the resource. Possible values include: 'Unknown', 'Succeeded', 'Failed', 'Canceled', 'Running', 'Creating', 'Updating', 'Deleting', 'Moving'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Type")},
			{
				Name:        "external_ip",
				Description: "The publicly accessible IP of the SignalR service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.ExternalIP")},
			{
				Name:        "host_name",
				Description: "FQDN of the SignalR service instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.HostName")},
			{
				Name:        "host_name_prefix",
				Description: "Prefix for the host name of the SignalR service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.HostNamePrefix")},
			{
				Name:        "kind",
				Description: "The kind of the service. Possible values include: 'SignalR', 'RawWebSockets'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Kind")},
			{
				Name:        "public_port",
				Description: "The publicly accessible port of the SignalR service which is designed for browser/client side usage.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.PublicPort")},
			{
				Name:        "server_port",
				Description: "The publicly accessible port of the SignalR service which is designed for customer server side usage.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.ServerPort")},
			{
				Name:        "version",
				Description: "Version of the SignalR resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.Version")},
			{
				Name:        "cors",
				Description: "Cross-Origin Resource Sharing (CORS) settings of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.Cors")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the SignalR service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "features",
				Description: "List of SignalR feature flags.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.Features")},
			{
				Name:        "network_acls",
				Description: "Network ACLs of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.NetworkACLs")},
			{
				Name:        "private_endpoint_connections",
				Description: "Private endpoint connections to the SignalR resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractSignalRServicePrivateEndpointConnections),
			},
			{
				Name:        "sku",
				Description: "The billing information of the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.SKU")},
			{
				Name:        "upstream",
				Description: "Upstream settings when the Azure SignalR is in server-less mode.",
				Type:        proto.ColumnType_JSON,
				Transform:

				// Steampipe standard columns
				transform.FromField("Description.ResourceInfo.Properties.Upstream")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.ResourceInfo.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceInfo.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

type SignalRServicePrivateEndpointConnections struct {
	PrivateEndpointPropertyID         interface{}
	PrivateLinkServiceConnectionState interface{}
	ProvisioningState                 interface{}
	ID                                *string
	Name                              *string
	Type                              *string
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

// Handle empty name or resourceGroup

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

// Create session

// If we return the API response directly, the output does not provide all
// the contents of DiagnosticSettings

//// TRANSFORM FUNCTION

// If we return the API response directly, the output will not provide all the properties of PrivateEndpointConnections
func extractSignalRServicePrivateEndpointConnections(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	service := d.HydrateItem.(opengovernance.SignalrService).Description.ResourceInfo
	info := []SignalRServicePrivateEndpointConnections{}

	if service.Properties != nil && service.Properties.PrivateEndpointConnections != nil {
		for _, connection := range service.Properties.PrivateEndpointConnections {
			properties := SignalRServicePrivateEndpointConnections{}
			properties.ID = connection.ID
			properties.Name = connection.Name
			properties.Type = connection.Type
			if connection.Properties != nil {
				if connection.Properties.PrivateEndpoint != nil {
					properties.PrivateEndpointPropertyID = connection.Properties.PrivateEndpoint.ID
				}
				properties.PrivateLinkServiceConnectionState = connection.Properties.PrivateLinkServiceConnectionState
				properties.ProvisioningState = connection.Properties.ProvisioningState
			}
			info = append(info, properties)
		}
	}

	return info, nil
}
