package azuread

import (
	"context"
	"github.com/opengovern/og-describer-entraid/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureAdTenant(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azuread_tenant",
		Description: "Represents an Azure AD Tenant.",
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetAdTenant,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Request_ResourceNotFound", "Invalid object identifier"}),
			},
			KeyColumns: plugin.SingleColumn("tenant_id"),
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAdTenant,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isIgnorableErrorPredicate([]string{"Invalid filter clause"}),
			},
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{Name: "tenant_id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the Tenant.", Transform: transform.FromField("Description.TenantID")},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The name displayed in the address book for the Tenant.", Transform: transform.FromField("Description.DisplayName")},
			{Name: "tenant_type", Type: proto.ColumnType_STRING, Description: "The type of the Tenant.", Transform: transform.FromField("Description.TenantType")},
			{Name: "created_date_time", Type: proto.ColumnType_TIMESTAMP, Description: "The creation time for the Tenant.", Transform: transform.FromField("Description.CreatedDateTime")},
			{Name: "verified_domains", Type: proto.ColumnType_JSON, Description: "Tenant verified domains.", Transform: transform.FromField("Description.VerifiedDomains")},
			{Name: "on_premises_sync_enabled", Type: proto.ColumnType_BOOL, Description: "Tenant verified domains.", Transform: transform.FromField("Description.OnPremisesSyncEnabled")},
			{
				Name:        "metadata",
				Description: "Metadata of the Azure resource",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata").Transform(marshalJSON),
			},
			{
				Name:        "og_account_id",
				Type:        proto.ColumnType_STRING,
				Description: "The Platform Account ID in which the resource is located.",
				Transform:   transform.FromField("Metadata.SourceID")},
			{
				Name:        "og_resource_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique ID of the resource in opengovernance.",
				Transform:   transform.FromField("ID")},
		}),
	}
}
