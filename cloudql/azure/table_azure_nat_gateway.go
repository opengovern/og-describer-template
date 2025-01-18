package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureNatGateway(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_nat_gateway",
		Description: "Azure NAT Gateway",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetNatGateway,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "404"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListNatGateway,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name that identifies the nat gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.Name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a nat gateway uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.ID")},
			{
				Name:        "etag",
				Description: "An unique read-only string that changes whenever the resource is updated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.Etag")},
			{
				Name:        "idle_timeout_in_minutes",
				Description: "The idle timeout of the nat gateway.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromP(extractNatGatewayProperties, "idleTimeoutInMinutes"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the nat gateway resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(extractNatGatewayProperties, "provisioningState"),
			},
			{
				Name:        "resource_guid",
				Description: "The provisioning state of the nat gateway resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(extractNatGatewayProperties, "resourceGUID"),
			},
			{
				Name:        "sku_name",
				Description: "The nat gateway SKU.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.SKU.Name")},
			{
				Name:        "type",
				Description: "The resource type of the nat gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.Type")},
			{
				Name:        "public_ip_addresses",
				Description: "An array of public ip addresses associated with the nat gateway resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(extractNatGatewayProperties, "publicIpAddresses"),
			},
			{
				Name:        "public_ip_prefixes",
				Description: "An array of public ip prefixes associated with the nat gateway resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(extractNatGatewayProperties, "publicIpPrefixes"),
			},
			{
				Name:        "subnets",
				Description: "An array of references to the subnets using this nat gateway resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(extractNatGatewayProperties, "subnets"),
			},
			{
				Name:        "zones",
				Description: "A list of availability zones denoting the zone in which Nat Gateway should be deployed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NatGateway.Zones")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NatGateway.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NatGateway.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.NatGateway.ID").Transform(idToAkas),
			},
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.NatGateway.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				Transform: transform.

					//// TRANSFORM FUNCTIONS
					FromField("Description.ResourceGroup")},
		}),
	}
}

func extractNatGatewayProperties(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	gateway := d.HydrateItem.(opengovernance.NatGateway).Description.NatGateway
	param := d.Param.(string)

	objectMap := make(map[string]interface{})

	if gateway.Properties.IdleTimeoutInMinutes != nil {
		objectMap["idleTimeoutInMinutes"] = *gateway.Properties.IdleTimeoutInMinutes
	}
	if gateway.Properties.ResourceGUID != nil {
		if *gateway.Properties.ResourceGUID != "" {
			objectMap["resourceGUID"] = gateway.Properties.ResourceGUID
		}
	}
	if gateway.Properties.ProvisioningState != nil {
		objectMap["provisioningState"] = gateway.Properties.ProvisioningState
	}
	if gateway.Properties.PublicIPAddresses != nil {
		objectMap["publicIpAddresses"] = gateway.Properties.PublicIPAddresses
	}
	if gateway.Properties.PublicIPPrefixes != nil {
		objectMap["publicIpPrefixes"] = gateway.Properties.PublicIPPrefixes
	}
	if gateway.Properties.Subnets != nil {
		objectMap["subnets"] = gateway.Properties.Subnets
	}

	if val, ok := objectMap[param]; ok {
		return val, nil
	}
	return nil, nil
}
