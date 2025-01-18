package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureLighthouseAssignment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_lighthouse_assignment",
		Description: "Azure Lighthouse Assignment",
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: isNotFoundError([]string{"SubscriptionNotFound"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "registration_assignment_id",
					Require: plugin.Required,
				},
				{
					Name:      "scope",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
			},
			Hydrate: opengovernance.GetLighthouseAssignment,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"RegistrationAssignmentNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate:    opengovernance.ListLighthouseAssignment,
			KeyColumns: plugin.OptionalColumns([]string{"scope"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the registration assignment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseAssignment.Name"),
			},
			{
				Name:        "id",
				Description: "Fully qualified path of the registration assignment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseAssignment.ID"),
			},
			{
				Name:        "registration_assignment_id",
				Description: "The ID of the registration assignment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseAssignment.ID").Transform(lastPathElement),
			},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseAssignment.Type"),
			},
			{
				Name:        "scope",
				Description: "The scope of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Scope"),
			},
			{
				Name:        "registration_definition_id",
				Description: "ID of the associated registration definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseAssignment.Properties.RegistrationDefinitionID"),
			},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the registration assignment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseAssignment.Properties.ProvisioningState"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseAssignment.Name"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LighthouseAssignment.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		},
	}
}
