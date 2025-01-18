package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzurePolicyAssignment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_policy_assignment",
		Description: "Azure Policy Assignment",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetPolicyAssignment,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPolicyAssignment,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the policy assignment.",
				Transform:   transform.FromField("Description.Assignment.ID")},
			{
				Name:        "name",
				Description: "The name of the policy assignment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assignment.Name")},
			{
				Name:        "display_name",
				Description: "The display name of the policy assignment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assignment.Properties.DisplayName")},
			{
				Name:        "policy_definition_id",
				Description: "The ID of the policy definition or policy set definition being assigned.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assignment.Properties.PolicyDefinitionID")},
			{
				Name:        "description",
				Description: "This message will be part of response in case of policy violation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assignment.Properties.Description")},
			{
				Name:        "enforcement_mode",
				Description: "The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assignment.Properties.EnforcementMode")},
			{
				Name:        "scope",
				Description: "The scope for the policy assignment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assignment.Properties.Scope")},
			{
				Name:        "sku_name",
				Description: "The name of the policy sku.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Resource.SKU.Name")},
			{
				Name:        "sku_tier",
				Description: "The policy sku tier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Resource.SKU.Tier"),
			},
			{
				Name:        "type",
				Description: "The type of the policy assignment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assignment.Type")},
			{
				Name:        "identity",
				Description: "The managed identity associated with the policy assignment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Assignment.Identity")},
			{
				Name:        "metadata",
				Description: "The policy assignment metadata.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Assignment.Properties.Metadata")},
			{
				Name:        "not_scopes",
				Description: "The policy's excluded scopes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Assignment.Properties.NotScopes")},
			{
				Name:        "parameters",
				Description: "The parameter values for the assigned policy rule.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Assignment.Properties.Parameters")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Assignment.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				//// LIST FUNCTION

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.Assignment.ID").Transform(idToAkas),
			},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS
