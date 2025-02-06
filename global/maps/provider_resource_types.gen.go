package maps

import (
	"github.com/opengovern/og-describer-render/discovery/describers"
	model "github.com/opengovern/og-describer-render/discovery/pkg/models"
	"github.com/opengovern/og-describer-render/discovery/provider"
	"github.com/opengovern/og-describer-render/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

var ResourceTypes = map[string]model.ResourceType{

	"Render/Blueprint": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Blueprint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListBlueprints),
		GetDescriber:    provider.DescribeSingleByRender(describers.GetBlueprint),
	},

	"Render/Deploy": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Deploy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListDeploys),
		GetDescriber:    nil,
	},

	"Render/Disk": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Disk",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListDisks),
		GetDescriber:    provider.DescribeSingleByRender(describers.GetDisk),
	},

	"Render/EnvGroup": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/EnvGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListEnvGroups),
		GetDescriber:    provider.DescribeSingleByRender(describers.GetEnvGroup),
	},

	"Render/Environment": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Environment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListEnvironments),
		GetDescriber:    provider.DescribeSingleByRender(describers.GetEnvironment),
	},

	"Render/Header": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Header",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListHeaders),
		GetDescriber:    nil,
	},

	"Render/Job": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Job",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListJobs),
		GetDescriber:    nil,
	},

	"Render/PostgresInstance": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/PostgresInstance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListPostgresInstances),
		GetDescriber:    provider.DescribeSingleByRender(describers.GetPostgresInstance),
	},

	"Render/Project": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Project",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListProjects),
		GetDescriber:    provider.DescribeSingleByRender(describers.GetProject),
	},

	"Render/Route": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Route",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListRoutes),
		GetDescriber:    nil,
	},

	"Render/Service": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Service",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListServices),
		GetDescriber:    provider.DescribeSingleByRender(describers.GetService),
	},

	"Render/Postgresql/Backup": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Render/Postgresql/Backup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByRender(describers.ListPostgresqlBackups),
		GetDescriber:    nil,
	},
}

var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"Render/Blueprint": {
		Name:            "Render/Blueprint",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Deploy": {
		Name:            "Render/Deploy",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Disk": {
		Name:            "Render/Disk",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/EnvGroup": {
		Name:            "Render/EnvGroup",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Environment": {
		Name:            "Render/Environment",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Header": {
		Name:            "Render/Header",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Job": {
		Name:            "Render/Job",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/PostgresInstance": {
		Name:            "Render/PostgresInstance",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Project": {
		Name:            "Render/Project",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Route": {
		Name:            "Render/Route",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Service": {
		Name:            "Render/Service",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Render/Postgresql/Backup": {
		Name:            "Render/Postgresql/Backup",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},
}

var ResourceTypesList = []string{
	"Render/Blueprint",
	"Render/Deploy",
	"Render/Disk",
	"Render/EnvGroup",
	"Render/Environment",
	"Render/Header",
	"Render/Job",
	"Render/PostgresInstance",
	"Render/Project",
	"Render/Route",
	"Render/Service",
	"Render/Postgresql/Backup",
}
