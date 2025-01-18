package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureEventGridDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_eventgrid_domain",
		Description: "Azure Event Grid Domain",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetEventGridDomain,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "400", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEventGridDomain,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Name")},
			{
				Name:        "id",
				Description: "Fully qualified identifier of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.ID")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Type")},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the event grid domain resource. Possible values include: 'Creating', 'Updating', 'Deleting', 'Succeeded', 'Canceled', 'Failed'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Properties.ProvisioningState")},
			{
				Name:        "auto_create_topic_with_first_subscription",
				Description: "This Boolean is used to specify the creation mechanism for 'all' the event grid domain topics associated with this event grid domain resource.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.Properties.AutoCreateTopicWithFirstSubscription")},
			{
				Name:        "auto_delete_topic_with_last_subscription",
				Description: "This Boolean is used to specify the deletion mechanism for 'all' the Event Grid Domain Topics associated with this Event Grid Domain resource.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.Properties.AutoDeleteTopicWithLastSubscription")},
			{
				Name:        "created_at",
				Description: "The timestamp of resource creation (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Domain.SystemData.CreatedAt").Transform(convertDateToTime),
			},
			{
				Name:        "created_by",
				Description: "The identity that created the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.SystemData.CreatedBy")},
			{
				Name:        "created_by_type",
				Description: "The type of identity that created the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.SystemData.CreatedByType")},
			{
				Name:        "disable_local_auth",
				Description: "This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the domain.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.Properties.DisableLocalAuth")},
			{
				Name:        "endpoint",
				Description: "Endpoint for the Event Grid Domain Resource which is used for publishing the events.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Properties.Endpoint")},
			{
				Name:        "identity_type",
				Description: "The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identity. Possible values include: 'None', 'SystemAssigned', 'UserAssigned', 'SystemAssignedUserAssigned'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Domain.Identity.Type"),
			},
			{
				Name:        "input_schema",
				Description: "This determines the format that Event Grid should expect for incoming events published to the Event Grid Domain Resource. Possible values include: 'EventGridSchema', 'CustomEventSchema', 'CloudEventSchemaV10'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Properties.InputSchema")},
			{
				Name:        "last_modified_at",
				Description: "The timestamp of resource last modification (UTC).",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.Domain.SystemData.LastModifiedAt").Transform(convertDateToTime),
			},
			{
				Name:        "last_modified_by",
				Description: "The identity that last modified the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.SystemData.LastModifiedBy")},
			{
				Name:        "last_modified_by_type",
				Description: "The type of identity that last modified the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.SystemData.LastModifiedByType")},
			{
				Name:        "location",
				Description: "Location of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Location")},
			{
				Name:        "principal_id",
				Description: "The principal ID of resource identity.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Domain.Identity.PrincipalID"),
			},
			{
				Name:        "public_network_access",
				Description: "This determines if traffic is allowed over public network. By default it is enabled.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Properties.PublicNetworkAccess")},
			{
				Name:        "sku_name",
				Description: "Name of this SKU. Possible values include: 'Basic', 'Standard'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Name"),
			},
			{
				Name:        "user_assigned_identities",
				Description: "The list of user identities associated with the resource. The user identity dictionary key references will be ARM resource ids.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.Identity.UserAssignedIdentities")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the eventgrid domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "inbound_ip_rules",
				Description: "This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.Properties.InboundIPRules")},
			{
				Name:        "input_schema_mapping",
				Description: "Information about the InputSchemaMapping which specified the info about mapping event payload.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.Properties.InputSchemaMapping")},
			{
				Name:        "private_endpoint_connections",
				Description: "List of private endpoint connections.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(extractEventgridDomainPrivaterEndPointConnections),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Domain.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Domain.Location").Transform(formatRegion).Transform(toLower),
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

// If we return the API response directly, the output does not provide
// all the contents of DiagnosticSettings

//// TRANSFORM FUNCTIONS

// If we return the private endpoint connection directly from api response we will not receive all the properties of private endpoint connections.
func extractEventgridDomainPrivaterEndPointConnections(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("extractEventgridDomainPrivaterEndPointConnections")
	domain := d.HydrateItem.(opengovernance.EventGridDomain).Description.Domain
	var privateEndpointConnectionsInfo []map[string]interface{}
	if domain.Properties.PrivateEndpointConnections != nil {
		privateEndpointConnections := domain.Properties.PrivateEndpointConnections
		for _, endpoint := range privateEndpointConnections {
			objectMap := make(map[string]interface{})
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
				if endpoint.Properties.PrivateEndpoint != nil {
					if endpoint.Properties.PrivateEndpoint.ID != nil {
						objectMap["endpointId"] = endpoint.Properties.PrivateEndpoint.ID
					}
				}
				if endpoint.Properties.GroupIDs != nil {
					objectMap["groupIds"] = endpoint.Properties.GroupIDs
				}
				if endpoint.Properties.ProvisioningState != nil {
					if *endpoint.Properties.ProvisioningState != "" {
						objectMap["provisioningState"] = endpoint.Properties.ProvisioningState
					}
				}
				if endpoint.Properties.PrivateLinkServiceConnectionState != nil {
					if endpoint.Properties.PrivateLinkServiceConnectionState.Status != nil {
						if *endpoint.Properties.PrivateLinkServiceConnectionState.Status != "" {
							objectMap["privateLinkServiceConnectionStateStatus"] = endpoint.Properties.PrivateLinkServiceConnectionState.Status
						}
					}
					if endpoint.Properties.PrivateLinkServiceConnectionState.Description != nil {
						objectMap["privateLinkServiceConnectionStateDescription"] = endpoint.Properties.PrivateLinkServiceConnectionState.Description
					}
					if endpoint.Properties.PrivateLinkServiceConnectionState.ActionsRequired != nil {
						objectMap["privateLinkServiceConnectionStateActionsRequired"] = endpoint.Properties.PrivateLinkServiceConnectionState.ActionsRequired
					}
				}
			}
			privateEndpointConnectionsInfo = append(privateEndpointConnectionsInfo, objectMap)
		}
	}
	return privateEndpointConnectionsInfo, nil
}
