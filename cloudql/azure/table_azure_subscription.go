package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSubscription(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_subscription",
		Description: "Azure Subscription",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSubscription,
		},

		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The fully qualified ID for the subscription. For example, /subscriptions/00000000-0000-0000-0000-000000000000.",
				Transform:   transform.FromField("Description.Subscription.ID")},
			{
				Name:        "subscription_id",
				Description: "The subscription ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscription.SubscriptionID")},
			{
				Name:        "display_name",
				Description: "A friendly name that identifies a subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscription.DisplayName")},
			{
				Name:        "tenant_id",
				Description: "The subscription tenant ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TenantID"),
			},
			{
				Name:        "state",
				Description: "The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted. Possible values include: 'StateEnabled', 'StateWarned', 'StatePastDue', 'StateDisabled', 'StateDeleted'",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Subscription.State"),
			},
			{
				Name:        "authorization_source",
				Description: "The authorization source of the request. Valid values are one or more combinations of Legacy, RoleBased, Bypassed, Direct and Management. For example, 'Legacy, RoleBased'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscription.AuthorizationSource")},
			{
				Name:        "managed_by_tenants",
				Description: "An array containing the tenants managing the subscription.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "subscription_policies",
				Description: "The subscription policies.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subscription.SubscriptionPolicies")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Subscription.DisplayName")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Subscription.ID").Transform(idToAkas),
			},
		}),
	}
}
