package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureLighthouseDefinition(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_lighthouse_definition",
		Description: "Azure Lighthouse Definition",
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: isNotFoundError([]string{"SubscriptionNotFound"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "registration_definition_id",
					Require: plugin.Required,
				},
				{
					Name:      "scope",
					Require:   plugin.Optional,
					Operators: []string{"="},
				},
			},
			Hydrate: opengovernance.GetLighthouseDefinition,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"RegistrationDefinitionNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate:    opengovernance.ListMaintenanceConfiguration,
			KeyColumns: plugin.OptionalColumns([]string{"scope"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the registration definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseDefinition.Name"),
			},
			{
				Name:        "id",
				Description: "Fully qualified path of the registration definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseDefinition.ID"),
			},
			{
				Name:        "registration_definition_id",
				Description: "The ID of the registration definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseDefinition.ID").Transform(lastPathElement),
			},
			{
				Name:        "type",
				Description: "Type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseDefinition.Type"),
			},
			{
				Name:        "scope",
				Description: "The scope of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Scope"),
			},
			{
				Name:        "description",
				Description: "Description of the registration definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseDefinition.Properties.Description"),
			},
			{
				Name:        "registration_definition_name",
				Description: "Name of the registration definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseDefinition.Properties.RegistrationDefinitionName"),
			},
			{
				Name:        "managed_by_tenant_id",
				Description: "ID of the managedBy tenant.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseDefinition.Properties.ManagedByTenantID"),
			},
			{
				Name:        "managed_by_tenant_name",
				Description: "The name of the managedBy tenant.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseDefinition.Properties.ManagedByTenantName"),
			},
			{
				Name:        "managed_tenant_name",
				Description: "The name of the managed tenant.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseDefinition.Properties.ManageeTenantName"),
			},
			{
				Name:        "authorizations",
				Description: "Authorization details containing principal ID and role ID.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LighthouseDefinition.Properties.Authorizations"),
			},
			{
				Name:        "eligible_authorizations",
				Description: "The collection of eligible authorization objects describing the just-in-time access Azure Active Directory principals in the managedBy tenant will receive on the delegated resource in the managed tenant.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LighthouseDefinition.Properties.EligibleAuthorizations"),
			},
			{
				Name:        "plan",
				Description: "Plan details for the managed services.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LighthouseDefinition.Name"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LighthouseDefinition.ID").Transform(idToAkas),
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
