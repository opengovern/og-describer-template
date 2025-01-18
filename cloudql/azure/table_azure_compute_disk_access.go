package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureComputeDiskAccess(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_disk_access",
		Description: "Azure Compute Disk Access",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetComputeDiskAccess,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeDiskAccess,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiskAccess.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiskAccess.ID")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiskAccess.Type")},
			{
				Name:        "provisioning_state",
				Description: "The disk access resource provisioning state.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiskAccess.Properties.ProvisioningState")},
			{
				Name:        "time_created",
				Description: "The time when the disk access was created.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.DiskAccess.Properties.TimeCreated").Transform(convertDateToTime),
			},
			{
				Name:        "private_endpoint_connections",
				Description: "The private endpoint connections details.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractPrivateEndpointConnections),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DiskAccess.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiskAccess.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.DiskAccess.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.DiskAccess.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}

type PrivateEndpointConnection struct {
	// ID - READ-ONLY; private endpoint connection Id
	ID string
	// Name - READ-ONLY; private endpoint connection name
	Name string
	// Type - READ-ONLY; private endpoint connection type
	Type string
	// PrivateEndpointID - The Id of private end point.
	PrivateEndpointID string
	// ProvisioningState - The provisioning state of the private endpoint connection resource. Possible values include: 'PrivateEndpointConnectionProvisioningStateSucceeded', 'PrivateEndpointConnectionProvisioningStateCreating', 'PrivateEndpointConnectionProvisioningStateDeleting', 'PrivateEndpointConnectionProvisioningStateFailed'
	ProvisioningState                                string
	PrivateLinkServiceConnectionStateStatus          string
	PrivateLinkServiceConnectionStateDescription     string
	PrivateLinkServiceConnectionStateActionsRequired string
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

// If we return the API response directly, the output will not provide
// all the properties of PrivateEndpointConnections
func extractPrivateEndpointConnections(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	diskAccess := d.HydrateItem.(opengovernance.ComputeDiskAccess).Description.DiskAccess
	var PrivateEndpointConnections []PrivateEndpointConnection
	if diskAccess.Properties.PrivateEndpointConnections != nil {
		for _, connection := range diskAccess.Properties.PrivateEndpointConnections {
			var PrivateConnection PrivateEndpointConnection
			if connection.ID != nil {
				PrivateConnection.ID = *connection.ID
			}
			if connection.Name != nil {
				PrivateConnection.Name = *connection.Name
			}
			if connection.Type != nil {
				PrivateConnection.Type = *connection.Type
			}
			if connection.Properties != nil {
				if connection.Properties.PrivateEndpoint != nil {
					PrivateConnection.PrivateEndpointID = *connection.Properties.PrivateEndpoint.ID
				}
				if connection.Properties.PrivateLinkServiceConnectionState != nil {
					if connection.Properties.PrivateLinkServiceConnectionState.ActionsRequired != nil {
						PrivateConnection.PrivateLinkServiceConnectionStateActionsRequired = *connection.Properties.PrivateLinkServiceConnectionState.ActionsRequired
					}
					if connection.Properties.PrivateLinkServiceConnectionState.Description != nil {
						PrivateConnection.PrivateLinkServiceConnectionStateDescription = *connection.Properties.PrivateLinkServiceConnectionState.Description
					}
					if connection.Properties.PrivateLinkServiceConnectionState.Status != nil {
						if *connection.Properties.PrivateLinkServiceConnectionState.Status != "" {
							PrivateConnection.PrivateLinkServiceConnectionStateStatus = string(*connection.Properties.PrivateLinkServiceConnectionState.Status)
						}
					}
				}
				if connection.Properties.ProvisioningState != nil {
					if *connection.Properties.ProvisioningState != "" {
						PrivateConnection.ProvisioningState = string(*connection.Properties.ProvisioningState)
					}
				}
			}

			PrivateEndpointConnections = append(PrivateEndpointConnections, PrivateConnection)
		}
	}
	return PrivateEndpointConnections, nil
}
