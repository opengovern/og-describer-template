package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/automation/mgmt/2019-06-01/automation"
	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureApAutomationVariable(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_automation_variable",
		Description: "Azure Automation Variable",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"account_name", "name", "resource_group"}),
			Hydrate:    opengovernance.GetAutomationVariables,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			ParentHydrate: opengovernance.ListAutomationAccounts,
			Hydrate:       opengovernance.ListAutomationVariables,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the resource.",
				Transform:   transform.FromField("Description.Automation.Name"),
			},
			{
				Name:        "account_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the account.",
				Transform:   transform.FromField("Description.AccountName"),
			},
			{
				Name:        "id",
				Description: "Fully qualified resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.ID"),
			},
			{
				Name:        "description",
				Description: "The description for the variable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Properties.Description"),
			},
			{
				Name:        "creation_time",
				Description: "The creation time of the variable.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Automation.Properties.CreationTime"),
			},
			{
				Name:        "last_modified_time",
				Description: "The last modified time of the variable.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Automation.Properties.LastModifiedTime"),
			},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Type"),
			},
			{
				Name:        "is_encrypted",
				Description: "The encrypted flag of the variable.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Automation.Properties.IsEncrypted"),
			},
			{
				Name:        "value",
				Description: "The value of the variable.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Properties.Value"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Name"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Automation.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup"),
			},
		}),
	}
}

type VariableDetails struct {
	AccountName string
	automation.Variable
}
