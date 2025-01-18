package fly

import (
	"context"
	"github.com/opengovern/og-describer-fly/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableFlySecret(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "fly_secret",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecret,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("label"),
			Hydrate:    opengovernance.GetSecret,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "label", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Label"), Description: "The label of the secret."},
			{Name: "public_key", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.PublicKey"), Description: "The public key of the secret as an array of integers."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Type"), Description: "The type of the secret."},
		}),
	}
}
