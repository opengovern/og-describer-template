package entraid

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-entraid/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableEntraIdApplication(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_application",
		Description: "Represents an Azure Active Directory (Azure AD) application.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdApplication,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name for the application.", Transform: transform.FromField("Description.DisplayName")},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the application.", Transform: transform.FromField("Description.Id")},
			{Name: "app_id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the application that is assigned to an application by Azure AD.", Transform: transform.FromField("Description.AppId")},
			{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time.", Transform: transform.FromField("Description.CreatedDateTime")},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Free text field to provide a description of the application object to end users.", Transform: transform.FromField("Description.Description")},
			{Name: "is_authorization_service_enabled", Type: proto.ColumnType_BOOL, Description: "Is authorization service enabled.", Default: false},
			{Name: "oauth2_require_post_response", Type: proto.ColumnType_BOOL, Description: "Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. The default is false, which specifies that only GET requests are allowed.", Transform: transform.FromField("Description.Oauth2RequirePostResponse"), Default: false},
			{Name: "publisher_domain", Type: proto.ColumnType_STRING, Description: "The verified publisher domain for the application.", Transform: transform.FromField("Description.PublisherDomain")},
			{Name: "sign_in_audience", Type: proto.ColumnType_STRING, Description: "Specifies the Microsoft accounts that are supported for the current application.", Transform: transform.FromField("Description.SignInAudience")},
			{Name: "api", Type: proto.ColumnType_JSON, Description: "Specifies settings for an application that implements a web API.", Transform: transform.FromField("Description.Api")},
			{Name: "identifier_uris", Type: proto.ColumnType_JSON, Description: "The URIs that identify the application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.", Transform: transform.FromField("Description.IdentifierUris")},
			{Name: "info", Type: proto.ColumnType_JSON, Description: "Basic profile information of the application such as app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience.", Transform: transform.FromField("Description.Info")},
			{Name: "key_credentials", Type: proto.ColumnType_JSON, Description: "The collection of key credentials associated with the application.", Transform: transform.FromField("Description.KeyCredentials")},
			{Name: "owner_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.OwnerIds"), Description: "Id of the owners of the application. The owners are a set of non-admin users who are allowed to modify this object."},
			{Name: "parental_control_settings", Type: proto.ColumnType_JSON, Description: "Specifies parental control settings for an application.", Transform: transform.FromField("Description.ParentalControlSettings")},
			{Name: "password_credentials", Type: proto.ColumnType_JSON, Description: "The collection of password credentials associated with the application.", Transform: transform.FromField("Description.PasswordCredentials")},
			{Name: "spa", Type: proto.ColumnType_JSON, Description: "Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.", Transform: transform.FromField("Description.Spa")},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Description: "Custom strings that can be used to categorize and identify the application.", Transform: transform.FromField("Description.TagsSrc")},
			{Name: "web", Type: proto.ColumnType_JSON, Description: "Specifies settings for a web application.", Transform: transform.FromField("Description.Web")},

			// Standard columns
			{Name: "tags", Type: proto.ColumnType_JSON, Description: ColumnDescriptionTags, Transform: transform.From(adApplicationTags)},
			{Name: "title", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTitle, Transform: transform.From(adApplicationTitle)},
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: ColumnDescriptionTenant, Transform: transform.FromField("Description.TenantID")},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func adApplicationTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	application := d.HydrateItem.(opengovernance.AdApplication).Description
	tags := application.TagsSrc
	return TagsToMap(tags)
}

func adApplicationTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.AdApplication).Description

	title := data.DisplayName
	if title == nil {
		title = data.Id
	}

	return title, nil
}
