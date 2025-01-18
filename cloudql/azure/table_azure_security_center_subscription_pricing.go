package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureSecurityCenterPricing(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_security_center_subscription_pricing",
		Description: "Azure Security Center Subscription Pricing",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetSecurityCenterSubscriptionPricing,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityCenterSubscriptionPricing,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The pricing id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pricing.ID")},
			{
				Name:        "name",
				Description: "Name of the pricing.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pricing.Name")},
			{
				Name:        "pricing_tier",
				Type:        proto.ColumnType_STRING,
				Description: "The pricing tier value. Azure Security Center is provided in two pricing tiers: free and standard, with the standard tier available with a trial period. The standard tier offers advanced security capabilities, while the free tier offers basic security features.",
				Transform:   transform.FromField("Description.Pricing.Properties.PricingTier")},
			{
				Name:        "free_trial_remaining_time",
				Description: "The duration left for the subscriptions free trial period.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pricing.Properties.FreeTrialRemainingTime")},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Type of the pricing.",

				// Steampipe standard columns
				Transform: transform.FromField("Description.Pricing.Type")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Pricing.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				//// LIST FUNCTION

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.Pricing.ID").Transform(idToAkas),
			},
		}),
	}
}

//// HYDRATE FUNCTIONS

// Handle empty input for get call
