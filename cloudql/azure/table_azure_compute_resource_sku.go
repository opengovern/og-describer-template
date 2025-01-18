package azure

import (
	"context"
	"strings"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureResourceSku(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_resource_sku",
		Description: "Azure Compute Resource SKU",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeResourceSKU,
		},

		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of SKU",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceSKU.Name")},
			{
				Name:        "resource_type",
				Description: "The type of resource the SKU applies to",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceSKU.ResourceType")},
			{
				Name:        "tier",
				Description: "Specifies the tier of virtual machines in a scale set",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceSKU.Tier")},
			{
				Name:        "size",
				Description: "The Size of the SKU",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceSKU.Size")},
			{
				Name:        "family",
				Description: "The Family of this particular SKU",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceSKU.Family")},
			{
				Name:        "kind",
				Description: "The Kind of resources that are supported in this SKU",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceSKU.Kind")},
			{
				Name:        "default_capacity",
				Description: "Contains the default capacity",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ResourceSKU.Capacity.Default")},
			{
				Name:        "maximum_capacity",
				Description: "The maximum capacity that can be set",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ResourceSKU.Capacity.Maximum")},
			{
				Name:        "minimum_capacity",
				Description: "The minimum capacity that can be set",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ResourceSKU.Capacity.Minimum")},
			{
				Name:        "scale_type",
				Description: "The scale type applicable to the sku",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceSKU.Capacity.ScaleType"),
			},
			{
				Name:        "api_versions",
				Description: "The api versions that support this SKU",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceSKU.APIVersions")},
			{
				Name:        "capabilities",
				Description: "A name value pair to describe the capability",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(computeResourceSkuCapabilities),
			},
			{
				Name:        "costs",
				Description: "A list of metadata for retrieving price info",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(computeResourceSkuCosts),
			},
			{
				Name:        "location_info",
				Description: "A list of locations and availability zones in those locations where the SKU is available",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(computeResourceSkuLocationInfo),
			},
			{
				Name:        "locations",
				Description: "The set of locations that the SKU is available",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceSKU.Locations")},
			{
				Name:        "restrictions",
				Description: "The restrictions because of which SKU cannot be used",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(computeResourceSkuRestrictions),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceSKU.Name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(skuDataToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTION ////

func skuDataToAkas(ctx context.Context, d *transform.TransformData) (any, error) {
	sku := d.HydrateItem.(opengovernance.ComputeResourceSKU)
	locations := sku.Description.ResourceSKU.Locations
	id := "azure:///subscriptions/" + sku.Metadata.SubscriptionID + "/locations/" + *locations[0] + "/resourcetypes" + *sku.Description.ResourceSKU.ResourceType + "name/" + *sku.Description.ResourceSKU.Name
	akas := []string{strings.ToLower(id)}
	return akas, nil
}

//// HELPER TRANSFORM FUNCTIONS to populate columns always returning [{}]

func computeResourceSkuCapabilities(ctx context.Context, d *transform.TransformData) (any, error) {
	skuData := d.HydrateItem.(opengovernance.ComputeResourceSKU).Description.ResourceSKU
	if skuData.Capabilities == nil {
		return nil, nil
	}
	capabilities := []map[string]interface{}{}

	for _, a := range skuData.Capabilities {
		data := map[string]interface{}{}
		if a.Name != nil {
			data["name"] = *a.Name
		}
		if a.Value != nil {
			data["value"] = *a.Value
		}
		capabilities = append(capabilities, data)
	}

	return capabilities, nil
}

func computeResourceSkuRestrictions(ctx context.Context, d *transform.TransformData) (any, error) {
	skuData := d.HydrateItem.(opengovernance.ComputeResourceSKU).Description.ResourceSKU
	restrictions := []map[string]interface{}{}

	for _, a := range skuData.Restrictions {
		data := map[string]interface{}{}
		data["type"] = &a.Type
		data["reasonCode"] = &a.ReasonCode
		if a.Values != nil {
			data["Values"] = a.Values
		}
		if a.RestrictionInfo != nil {
			data["restrictionInfo"] = *a.RestrictionInfo
		}
		restrictions = append(restrictions, data)
	}

	return restrictions, nil
}

func computeResourceSkuLocationInfo(ctx context.Context, d *transform.TransformData) (any, error) {
	skuData := d.HydrateItem.(opengovernance.ComputeResourceSKU).Description.ResourceSKU
	locationInfo := []map[string]interface{}{}

	for _, a := range skuData.LocationInfo {
		data := map[string]interface{}{}
		if a.Location != nil {
			data["location"] = *a.Location
		}
		if a.Zones != nil {
			data["zones"] = a.Zones
		}
		locationInfo = append(locationInfo, data)
	}

	return locationInfo, nil
}

func computeResourceSkuCosts(ctx context.Context, d *transform.TransformData) (any, error) {
	skuData := d.HydrateItem.(opengovernance.ComputeResourceSKU).Description.ResourceSKU
	costs := []map[string]interface{}{}

	for _, a := range skuData.Costs {
		data := map[string]interface{}{}
		if a.MeterID != nil {
			data["meterID"] = *a.MeterID
		}
		if a.Quantity != nil {
			data["quantity"] = *a.Quantity
		}
		if a.ExtendedUnit != nil {
			data["extendedUnit"] = *a.ExtendedUnit
		}
		costs = append(costs, data)
	}

	return costs, nil
}
