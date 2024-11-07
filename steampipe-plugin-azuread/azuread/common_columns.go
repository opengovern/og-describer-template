package azuread

import (
	"context"
	"encoding/json"
	"github.com/opengovern/og-azure-describer/azure/describer"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// column definitions for the common columns
func commonKaytuColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:      "cloud_environment",
			Type:      proto.ColumnType_STRING,
			Transform: transform.FromField("Metadata.CloudEnvironment"),
		},
		{
			Name:      "subscription_id",
			Type:      proto.ColumnType_STRING,
			Transform: transform.FromField("Metadata.SubscriptionID"),
		},
		{
			Name:        "og_account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The Platform Account ID in which the resource is located.",
			Transform:   transform.FromField("Metadata.SourceID"),
		},
		{
			Name:        "og_resource_id",
			Type:        proto.ColumnType_STRING,
			Description: "The unique ID of the resource in opengovernance.",
			Transform:   transform.FromField("ID"),
		},
		{
			Name:      "kaytu_metadata",
			Type:      proto.ColumnType_STRING,
			Transform: transform.FromField("Metadata").Transform(marshalJSON),
		},
		{
			Name:        "kaytu_description",
			Type:        proto.ColumnType_JSON,
			Description: "The full model description of the resource",
			Transform:   transform.FromField("Description").Transform(marshalAzureJSON),
		},
	}
}

// append the common azure columns onto the column list
func azureKaytuColumns(columns []*plugin.Column) []*plugin.Column {
	for _, c := range commonKaytuColumns() {
		found := false
		for _, col := range columns {
			if col.Name == c.Name {
				found = true
				break
			}
		}
		if !found {
			columns = append(columns, c)
		}
	}
	return columns
}

func marshalAzureJSON(_ context.Context, d *transform.TransformData) (interface{}, error) {
	b, err := json.Marshal(describer.JSONAllFieldsMarshaller{Value: d.Value})
	if err != nil {
		return nil, err
	}
	return string(b), nil
}
