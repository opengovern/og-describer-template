package render

import (
	"context"
	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-render",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: essdk.ConfigInstance,
			Schema:      essdk.ConfigSchema(),
		},
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"render_blueprint":         tableRenderBlueprint(ctx),
			"render_deploy":            tableRenderDeploy(ctx),
			"render_disk":              tableRenderDisk(ctx),
			"render_env_group":         tableRenderEnvGroup(ctx),
			"render_environment":       tableRenderEnvironment(ctx),
			"render_header":            tableRenderHeader(ctx),
			"render_job":               tableRenderJob(ctx),
			"render_postgres_instance": tableRenderPostgres(ctx),
			"render_project":           tableRenderProject(ctx),
			"render_route":             tableRenderRoute(ctx),
			"render_service":           tableRenderService(ctx),
			"render_postgresql_backup": tableRenderPostgresqlBackup(ctx),
		},
	}
	for key, table := range p.TableMap {
		if table == nil {
			continue
		}
		if table.Get != nil && table.Get.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}
		if table.List != nil && table.List.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}

		opengovernanceTable := false
		for _, col := range table.Columns {
			if col != nil && col.Name == "platform_integration_id" {
				opengovernanceTable = true
			}
		}

		if opengovernanceTable {
			if table.Get != nil {
				table.Get.KeyColumns = append(table.Get.KeyColumns, plugin.OptionalColumns([]string{"platform_integration_id", "platform_resource_id"})...)
			}

			if table.List != nil {
				table.List.KeyColumns = append(table.List.KeyColumns, plugin.OptionalColumns([]string{"platform_integration_id", "platform_resource_id"})...)
			}
		}
	}
	return p
}
