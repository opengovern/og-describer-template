package maps

import (
	"github.com/opengovern/og-describer-fly/discovery/describers"
	model "github.com/opengovern/og-describer-fly/discovery/pkg/models"
	"github.com/opengovern/og-describer-fly/discovery/provider"
	"github.com/opengovern/og-describer-fly/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

var ResourceTypes = map[string]model.ResourceType{

	"Fly/App": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Fly/App",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByFly(describers.ListApps),
		GetDescriber:    provider.DescribeSingleByFly(describers.GetApp),
	},

	"Fly/Machine": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Fly/Machine",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByFly(describers.ListMachines),
		GetDescriber:    provider.DescribeSingleByFly(describers.GetMachine),
	},

	"Fly/Volume": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Fly/Volume",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByFly(describers.ListVolumes),
		GetDescriber:    provider.DescribeSingleByFly(describers.GetVolume),
	},

	"Fly/Secret": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Fly/Secret",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByFly(describers.ListSecrets),
		GetDescriber:    nil,
	},
}

var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"Fly/App": {
		Name:            "Fly/App",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Fly/Machine": {
		Name:            "Fly/Machine",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Fly/Volume": {
		Name:            "Fly/Volume",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Fly/Secret": {
		Name:            "Fly/Secret",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},
}

var ResourceTypesList = []string{
	"Fly/App",
	"Fly/Machine",
	"Fly/Volume",
	"Fly/Secret",
}
