//go:generate go run ./table_azure_service_template.go

package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var table = Table{
	Service: "TrafficManager",
	Name:    "Profile",
}

const azureTableServiceTemplate = `
package azure

import (
	opengovernance "github.com/opengovern/og-describer-azure/pkg/SDK/generated"
	"context"
	"fmt"
	"github.com/opengovern/og-azure-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzure{{.Service}}{{.Name}}(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_{{.ServiceLowerCase}}_{{.NameLowerCase}}",
		Description: "Azure {{.Service}} {{.Name}}",
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:           opengovernance.Get{{.Service}}{{.Name}},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.List{{.Service}}{{.Name}},
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the {{.NameLowerCase}}.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.{{.Name}}.Id")},
			{
				Name:        "name",
				Description: "The name of the {{.NameLowerCase}}.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.{{.Name}}.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.{{.Name}}.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.{{.Name}}.Tags"), // probably needs a transform function
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.{{.Name}}.ID").Transform(idToAkas), // or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
			},
		}),
	}
}
`

type Table struct {
	Service, Name, ServiceLowerCase, NameLowerCase string
}

func main() {
	tmpl := template.New("azure_table")
	tmpl, err := tmpl.Parse(azureTableServiceTemplate)
	if err != nil {
		panic(err)
	}

	table.ServiceLowerCase = strings.ToLower(table.Service)
	table.NameLowerCase = strings.ToLower(table.Name)

	fileName := fmt.Sprintf("../table_azure_%s_%s.go", table.ServiceLowerCase, table.NameLowerCase)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, table)
	if err != nil {
		panic(err)
	}
}
