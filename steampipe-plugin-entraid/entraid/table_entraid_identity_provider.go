package entraid

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableEntraIdIdentityProvider(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "entraid_identity_provider",
		Description: "Represents an Azure Active Directory (Azure AD) identity provider.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdIdentityProvider,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_UnsupportedQuery", "Invalid filter clause"}),
			},
		},

		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the identity provider.",
				Transform:   transform.FromField("Description.Id")},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The display name of the identity provider.",
				Transform:   transform.FromField("Description.DisplayName")},

			// Other fields
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "The identity provider type is a required field. For B2B scenario: Google, Facebook. For B2C scenario: Microsoft, Google, Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo, QQ, WeChat, OpenIDConnect.",
				Transform:   transform.FromField("Description.Type")},
			{
				Name:        "client_id",
				Type:        proto.ColumnType_STRING,
				Description: "The client ID for the application. This is the client ID obtained when registering the application with the identity provider.",
				Transform:   transform.FromField("Description.ClientId")},

			{
				Name:        "client_secret",
				Type:        proto.ColumnType_STRING,
				Description: "The client secret for the application. This is the client secret obtained when registering the application with the identity provider. This is write-only. A read operation will return ****.",
				Transform:   transform.FromField("Description.ClientSecret")},

			// Standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("Description.DisplayName")},
			{
				Name:        "tenant_id",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTenant,
				Transform:   transform.FromField("Description.TenantID")},
		}),
	}
}
