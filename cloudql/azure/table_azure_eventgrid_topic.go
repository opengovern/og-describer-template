package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureEventGridTopic(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_eventgrid_topic",
		Description: "Azure Event Grid Topic",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetEventGridTopic,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "400", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEventGridTopic,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.Name")},
			{
				Name:        "id",
				Description: "Fully qualified identifier of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.ID")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.Type")},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the event grid topic resource. Possible values include: 'Creating', 'Updating', 'Deleting', 'Succeeded', 'Canceled', 'Failed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.Properties.ProvisioningState")},
			{
				Name:        "created_at",
				Description: "The timestamp of resource creation (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Topic.SystemData.CreatedAt").Transform(convertDateToTime),
			},
			{
				Name:        "created_by",
				Description: "The identity that created the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.SystemData.CreatedBy")},
			{
				Name:        "created_by_type",
				Description: "The type of identity that created the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.SystemData.CreatedByType")},
			{
				Name:        "disable_local_auth",
				Description: "This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the topic.",
				Type:        proto.ColumnType_BOOL,

				Transform: transform.FromField("Description.Topic.Properties.DisableLocalAuth"), Default: false,
			},
			{
				Name:        "endpoint",
				Description: "Endpoint for the event grid topic resource which is used for publishing the events.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.Properties.Endpoint")},
			{
				Name:        "input_schema",
				Description: "This determines the format that event grid should expect for incoming events published to the event grid topic resource. Possible values include: 'EventGridSchema', 'CustomEventSchema', 'CloudEventSchemaV10'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.Properties.InputSchema")},
			{
				Name:        "kind",
				Description: "Kind of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.Type")},
			{
				Name:        "last_modified_at",
				Description: "The timestamp of resource last modification (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Topic.SystemData.LastModifiedAt").Transform(convertDateToTime),
			},
			{
				Name:        "last_modified_by",
				Description: "The identity that last modified the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.SystemData.LastModifiedBy")},
			{
				Name:        "last_modified_by_type",
				Description: "The type of identity that last modified the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.SystemData.LastModifiedByType")},
			{
				Name:        "location",
				Description: "Location of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.Location")},
			{
				Name:        "public_network_access",
				Description: "This determines if traffic is allowed over public network. By default it is enabled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.Properties.PublicNetworkAccess")},
			{
				Name:        "sku_name",
				Description: "Name of this SKU. Possible values include: 'Basic', 'Standard'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Topic.Name"),
			},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the eventgrid topic.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "extended_location",
				Description: "Extended location of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.Location")},
			{
				Name:        "identity",
				Description: "Identity information for the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Topic.Identity")},
			{
				Name:        "inbound_ip_rules",
				Description: "This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Topic.Properties.InboundIPRules")},
			{
				Name:        "input_schema_mapping",
				Description: "Information about the InputSchemaMapping which specified the info about mapping event payload.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Topic.Properties.InputSchemaMapping")},
			{
				Name:        "private_endpoint_connections",
				Description: "List of private endpoint connections for the event grid topic.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractEventgridTopicPrivaterEndPointConnections),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Topic.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Topic.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Topic.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Topic.Location").Transform(formatRegion).Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				// Create session
				Transform: transform.

					//// HYDRATE FUNCTIONS
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Return nil, if no input provided

// Create session

// Create session

// Pagination is not supported

// If we return the API response directly, the output does not provide
// all the contents of DiagnosticSettings

//// TRANSFORM FUNCTIONS

// If we return the private endpoint connection directly from api response we will not receive all the properties of private endpoint connections.
func extractEventgridTopicPrivaterEndPointConnections(ctx context.Context, d *transform.TransformData) (any, error) {
	var privateEndpointConnectionsInfo []map[string]any
	if d == nil {
		return privateEndpointConnectionsInfo, nil
	}
	if d.HydrateItem == nil {
		return privateEndpointConnectionsInfo, nil
	}
	topic := d.HydrateItem.(opengovernance.EventGridTopic).Description.Topic
	if topic.Properties != nil && topic.Properties.PrivateEndpointConnections != nil {
		for _, endpoint := range topic.Properties.PrivateEndpointConnections {
			objectMap := make(map[string]any)
			if endpoint == nil {
				continue
			}

			if endpoint.ID != nil {
				objectMap["id"] = endpoint.ID
			}

			if endpoint.Name != nil {
				objectMap["name"] = endpoint.Name
			}

			if endpoint.Type != nil {
				objectMap["type"] = endpoint.Type
			}

			if endpoint.Properties != nil {
				properties := *endpoint.Properties
				if properties.PrivateEndpoint != nil {
					if properties.PrivateEndpoint.ID != nil {
						objectMap["endpointId"] = properties.PrivateEndpoint.ID
					}
				}
				if properties.GroupIDs != nil {
					objectMap["groupIds"] = properties.GroupIDs
				}
				if properties.ProvisioningState != nil {
					if *properties.ProvisioningState != "" {
						objectMap["provisioningState"] = properties.ProvisioningState
					}
				}
				if properties.PrivateLinkServiceConnectionState != nil {
					state := *properties.PrivateLinkServiceConnectionState
					if state.Status != nil {
						if *state.Status != "" {
							objectMap["privateLinkServiceConnectionStateStatus"] = state.Status
						}
					}
					if state.Description != nil {
						objectMap["privateLinkServiceConnectionStateDescription"] = state.Description
					}
					if state.ActionsRequired != nil {
						objectMap["privateLinkServiceConnectionStateActionsRequired"] = state.ActionsRequired
					}
				}
			}
			privateEndpointConnectionsInfo = append(privateEndpointConnectionsInfo, objectMap)
		}
	}
	return privateEndpointConnectionsInfo, nil
}
