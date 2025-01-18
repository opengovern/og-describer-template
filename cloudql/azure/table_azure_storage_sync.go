package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureStorageSync(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_storage_sync",
		Description: "Azure Storage Sync",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetStorageSync,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStorageSync,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Name")},
			{
				Name:        "id",
				Description: "Fully qualified resource id for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.ID")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the storage sync service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Type")},
			{
				Name:        "incoming_traffic_policy",
				Description: "The incoming traffic policy of the storage sync service. Possible values include: 'AllowAllTraffic', 'AllowVirtualNetworksOnly'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Properties.IncomingTrafficPolicy")},
			{
				Name:        "last_operation_name",
				Description: "The last operation name of the storage sync service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Properties.LastOperationName")},
			{
				Name:        "last_workflow_id",
				Description: "The last workflow id of the storage sync service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Properties.LastWorkflowID")},
			{
				Name:        "storage_sync_service_status",
				Description: "The status of the storage sync service.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Service.Properties.StorageSyncServiceStatus")},
			{
				Name:        "storage_sync_service_uid",
				Description: "The uid of the storage sync service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Properties.StorageSyncServiceUID")},
			{
				Name:        "private_endpoint_connections",
				Description: "List of private endpoint connection associated with the specified storage sync service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractStorageSyncPrivateEndpointConnections),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Service.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Service.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Service.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Service.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

type StorageSyncPrivateEndpointConnections struct {
	PrivateEndpointPropertyID         interface{}
	PrivateLinkServiceConnectionState interface{}
	ProvisioningState                 interface{}
	ID                                *string
	Name                              *string
	Type                              *string
}

//// LIST FUNCTION

//// TRANSFORM FUNCTION

// If we return the API response directly, the output will not provide all the properties of PrivateEndpointConnections
func extractStorageSyncPrivateEndpointConnections(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	service := d.HydrateItem.(opengovernance.StorageSync).Description.Service
	info := []StorageSyncPrivateEndpointConnections{}

	if service.Properties != nil && service.Properties.PrivateEndpointConnections != nil {
		for _, connection := range service.Properties.PrivateEndpointConnections {
			properties := StorageSyncPrivateEndpointConnections{}
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
