package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzurePolicyDefinition(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_policy_definition",
		Description: "Azure Policy Definition",
		// Get API operation is not working as expected, skipping for now
		// Get: &plugin.GetConfig{
		// 	KeyColumns: plugin.SingleColumn("name"),
		// 	Hydrate:    getPolicyDefinition,
		// },
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPolicyDefinition,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the policy definition.",
				Transform:   transform.FromField("ResourceID"),
			},
			{
				Name:        "name",
				Description: "The name of the policy definition.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Definition.Name")},
			{
				Name:        "display_name",
				Description: "The user-friendly display name of the policy definition.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Definition.Properties.DisplayName")},
			{
				Name:        "description",
				Description: "The policy definition description.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Definition.Properties.Description")},
			{
				Name:        "mode",
				Description: "The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Definition.Properties.Mode")},
			{
				Name:        "policy_type",
				Description: "The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static. Possible values include: 'NotSpecified', 'BuiltIn', 'Custom', 'Static'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Definition.Properties.PolicyType")},
			{
				Name:        "type",
				Description: "The type of the resource (Microsoft.Authorization/policyDefinitions).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Definition.Type")},
			{
				Name:        "metadata",
				Description: "The policy definition metadata.  Metadata is an open ended object and is typically a collection of key value pairs.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Definition.Properties.Metadata")},
			{
				Name:        "parameters",
				Description: "The parameter definitions for parameters used in the policy rule. The keys are the parameter names.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Definition.Properties.Parameters")},
			{
				Name:        "policy_rule",
				Description: "The policy rule.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Definition.Properties.PolicyRule")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Definition.Properties.DisplayName")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.TurboData.Akas")},
		}),
	}
}
