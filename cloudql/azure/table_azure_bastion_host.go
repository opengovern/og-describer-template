package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureBastionHost(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_bastion_host",
		Description: "Azure Bastion Host",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetBastionHosts,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBastionHosts,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the bastion host.",
				Transform:   transform.FromField("Description.BastianHost.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a bastion host uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BastianHost.ID")},
			{
				Name:        "dns_name",
				Description: "FQDN for the endpoint on which bastion host is accessible.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BastianHost.Properties.DNSName")},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BastianHost.Etag")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the bastion host resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BastianHost.Properties.ProvisioningState")},
			{
				Name:        "type",
				Description: "The resource type of the bastion host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BastianHost.Type"),
			},
			{
				Name:        "ip_configurations",
				Description: "IP configuration of the bastion host resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BastianHost.Properties.IPConfigurations")},

			// Steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BastianHost.ID").Transform(idToAkas),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BastianHost.Tags")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BastianHost.Name")},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BastianHost.Location")},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup")},
		}),
	}
}
