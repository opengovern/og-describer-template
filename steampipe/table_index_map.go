package steampipe

import (
	"github.com/opengovern/og-describer-render/pkg/sdk/es"
)

var Map = map[string]string{
  "Render/Blueprint": "render_blueprint",
  "Render/Deploy": "render_deploy",
  "Render/Disk": "render_disk",
  "Render/EnvGroup": "render_env_group",
  "Render/Environment": "render_environment",
  "Render/Header": "render_header",
  "Render/Job": "render_job",
  "Render/PostgresInstance": "render_postgres_instance",
  "Render/Project": "render_project",
  "Render/Route": "render_route",
  "Render/Service": "render_service",
}

var DescriptionMap = map[string]interface{}{
  "Render/Blueprint": opengovernance.Blueprint{},
  "Render/Deploy": opengovernance.Deploy{},
  "Render/Disk": opengovernance.Disk{},
  "Render/EnvGroup": opengovernance.EnvGroup{},
  "Render/Environment": opengovernance.Environment{},
  "Render/Header": opengovernance.Header{},
  "Render/Job": opengovernance.Job{},
  "Render/PostgresInstance": opengovernance.Postgres{},
  "Render/Project": opengovernance.Project{},
  "Render/Route": opengovernance.Route{},
  "Render/Service": opengovernance.Service{},
}

var ReverseMap = map[string]string{
  "render_blueprint": "Render/Blueprint",
  "render_deploy": "Render/Deploy",
  "render_disk": "Render/Disk",
  "render_env_group": "Render/EnvGroup",
  "render_environment": "Render/Environment",
  "render_header": "Render/Header",
  "render_job": "Render/Job",
  "render_postgres_instance": "Render/PostgresInstance",
  "render_project": "Render/Project",
  "render_route": "Render/Route",
  "render_service": "Render/Service",
}
