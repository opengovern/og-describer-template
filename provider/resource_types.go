package provider

import (
	model "github.com/opengovern/og-describer-render/pkg/sdk/models"
	"github.com/opengovern/og-describer-render/provider/configs"
	"github.com/opengovern/og-describer-render/provider/describer"
)

var ResourceTypes = map[string]model.ResourceType{

	"Render/Blueprint": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Render/Blueprint",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByRender(describer.ListBlueprints),
		GetDescriber:    DescribeSingleByRender(describer.GetBlueprint),
	},

	"Render/Deploy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Render/Deploy",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByRender(describer.ListDeploys),
		GetDescriber:    nil,
	},

	"Render/Disk": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Render/Disk",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByRender(describer.ListDisks),
		GetDescriber:    DescribeSingleByRender(describer.GetDisk),
	},

	"Render/EnvGroup": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Render/EnvGroup",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByRender(describer.ListEnvGroups),
		GetDescriber:    DescribeSingleByRender(describer.GetEnvGroup),
	},

	"Render/Environment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Render/Environment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByRender(describer.ListEnvironments),
		GetDescriber:    DescribeSingleByRender(describer.GetEnvironment),
	},

	"Render/Header": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Render/Header",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByRender(describer.ListHeaders),
		GetDescriber:    nil,
	},

	"Render/Job": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Render/Job",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByRender(describer.ListJobs),
		GetDescriber:    nil,
	},

	"Render/PostgresInstance": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Render/PostgresInstance",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByRender(describer.ListPostgresInstances),
		GetDescriber:    DescribeSingleByRender(describer.GetPostgresInstance),
	},

	"Render/Project": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Render/Project",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByRender(describer.ListProjects),
		GetDescriber:    DescribeSingleByRender(describer.GetProject),
	},

	"Render/Route": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Render/Route",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByRender(describer.ListRoutes),
		GetDescriber:    nil,
	},

	"Render/Service": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Render/Service",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByRender(describer.ListServices),
		GetDescriber:    DescribeSingleByRender(describer.GetService),
	},
}
