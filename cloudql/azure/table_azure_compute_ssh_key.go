package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION ////

func tableAzureComputeSshKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_ssh_key",
		Description: "Azure Compute SSH Key",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetComputeSSHPublicKey,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeSSHPublicKey,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The unique ID identifying the resource in subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSHPublicKey.ID")},
			{
				Name:        "name",
				Description: "Name of the SSH key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSHPublicKey.Name")},
			{
				Name:        "type",
				Description: "The type of the resource in Azure.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSHPublicKey.Type")},
			{
				Name:        "public_key",
				Description: "SSH public key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSHPublicKey.Properties.PublicKey")},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSHPublicKey.Location").Transform(toLower),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SSHPublicKey.Tags")},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSHPublicKey.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SSHPublicKey.ID").Transform(idToAkas),
			},
		}),
	}
}
