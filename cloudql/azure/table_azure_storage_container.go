package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITIO

func tableAzureStorageContainer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_storage_container",
		Description: "Azure Storage Container",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group", "account_name"}),
			Hydrate:    opengovernance.GetStorageContainer,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "ContainerNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStorageContainer,
		},

		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the container.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ListContainerItem.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a container uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ListContainerItem.ID")},
			{
				Name:        "account_name",
				Description: "The friendly name that identifies the storage account.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AccountName")},
			{
				Name:        "deleted",
				Description: "Indicates whether the blob container was deleted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ListContainerItem.Properties.Deleted")},
			{
				Name:        "public_access",
				Description: "Specifies whether data in the container may be accessed publicly and the level of access.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ListContainerItem.Properties.PublicAccess"),
			},
			{
				Name:        "type",
				Description: "Specifies the type of the container.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ListContainerItem.Type")},
			{
				Name:        "default_encryption_scope",
				Description: "Default the container to use specified encryption scope for all writes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ListContainerItem.Properties.DefaultEncryptionScope")},
			{
				Name:        "deleted_time",
				Description: "Specifies the time when the container was deleted.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.ListContainerItem.Properties.DeletedTime").Transform(convertDateToTime),
			},
			{
				Name:        "deny_encryption_scope_override",
				Description: "Indicates whether block override of encryption scope from the container default, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ListContainerItem.Properties.DenyEncryptionScopeOverride")},
			{
				Name:        "has_immutability_policy",
				Description: "The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container. The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ListContainerItem.Properties.HasImmutabilityPolicy")},
			{
				Name:        "has_legal_hold",
				Description: "The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ListContainerItem.Properties.HasLegalHold")},
			{
				Name:        "last_modified_time",
				Description: "Specifies the date and time the container was last modified.",
				Type:        proto.ColumnType_TIMESTAMP,

				Transform: transform.FromField("Description.ListContainerItem.Properties.LastModifiedTime").Transform(convertDateToTime),
			},
			{
				Name:        "lease_status",
				Description: "Specifies the lease status of the container.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ListContainerItem.Properties.LeaseStatus"),
			},
			{
				Name:        "lease_state",
				Description: "Specifies the lease state of the container.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ListContainerItem.Properties.LeaseState"),
			},
			{
				Name:        "lease_duration",
				Description: "Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased. Possible values are: 'Infinite', 'Fixed'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ListContainerItem.Properties.LeaseDuration"),
			},
			{
				Name:        "remaining_retention_days",
				Description: "Remaining retention days for soft deleted blob container.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ListContainerItem.Properties.RemainingRetentionDays")},
			{
				Name:        "version",
				Description: "The version of the deleted blob container.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ListContainerItem.Properties.Version")},
			{
				Name:        "immutability_policy",
				Description: "The ImmutabilityPolicy property of the container.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ImmutabilityPolicy")},
			{
				Name:        "legal_hold",
				Description: "The LegalHold property of the container.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ListContainerItem.Properties.LegalHold")},
			{
				Name:        "metadata",
				Description: "A name-value pair to associate with the container as metadata.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.ListContainerItem.Properties.Metadata")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ListContainerItem.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.ListContainerItem.ID").Transform(idToAkas),
			},

			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.FromField("Description.ResourceGroup")},
		}),
	}
}
