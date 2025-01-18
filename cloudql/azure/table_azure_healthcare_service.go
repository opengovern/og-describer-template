package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureHealthcareService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_healthcare_service",
		Description: "Azure Healthcare Service",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetHealthcareService,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListHealthcareService,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServicesDescription.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServicesDescription.ID")},
			{
				Name:        "etag",
				Description: "An etag associated with the resource, used for optimistic concurrency when editing it.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServicesDescription.Etag")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServicesDescription.Type")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the healthcare service resource.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ServicesDescription.Properties.ProvisioningState"),
			},
			{
				Name:        "allow_credentials",
				Description: "If credentials are allowed via CORS.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ServicesDescription.Properties.CorsConfiguration.AllowCredentials")},
			{
				Name:        "audience",
				Description: "The audience url for the service.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ServicesDescription.Properties.AuthenticationConfiguration.Audience"),
			},
			{
				Name:        "authority",
				Description: "The authority url for the service.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ServicesDescription.Properties.AuthenticationConfiguration.Authority"),
			},
			{
				Name:        "kind",
				Description: "The kind of the service. Possible values include: 'Fhir', 'FhirStu3', 'FhirR4'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServicesDescription.Kind")},
			{
				Name:        "location",
				Description: "The resource location.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServicesDescription.Location")},
			{
				Name:        "max_age",
				Description: "The max age to be allowed via CORS.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ServicesDescription.Properties.CorsConfiguration.MaxAge")},
			{
				Name:        "smart_proxy_enabled",
				Description: "If the SMART on FHIR proxy is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ServicesDescription.Properties.AuthenticationConfiguration.SmartProxyEnabled")},
			{
				Name:        "access_policies",
				Description: "The access policies of the healthcare service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServicesDescription.Properties.AccessPolicies")},
			{
				Name:        "cosmos_db_configuration",
				Description: "The settings for the Cosmos DB database backing the service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServicesDescription.Properties.CosmosDbConfiguration")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the healthcare serive.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "headers",
				Description: "The headers to be allowed via CORS.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServicesDescription.Properties.CorsConfiguration.Origins")},
			{
				Name:        "methods",
				Description: "The methods to be allowed via CORS.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServicesDescription.Properties.CorsConfiguration.Methods")},
			{
				Name:        "origins",
				Description: "The origins to be allowed via CORS.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServicesDescription.Properties.CorsConfiguration.Origins")},
			{
				Name:        "private_endpoint_connections",
				Description: "List of private endpoint connections for healthcare service.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.ServicesDescription.Properties.PrivateEndpointConnections")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServicesDescription.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServicesDescription.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.ServicesDescription.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ServicesDescription.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				//// HYDRATE FUNCTIONS
				Transform: transform.

					// Empty check for param
					FromField("Description.ResourceGroup")},
		}),
	}
}

// In some cases resource does not give any notFound error
// instead of notFound error, it returns empty data

// Empty check

// SDK does not support pagination yet

// If we return the API response directly, the output will not provide the properties of PrivateEndpointConnections

// Empty check

// If we return the API response directly, the output will not provide all
// the contents of DiagnosticSettings
