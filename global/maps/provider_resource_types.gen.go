package maps

import (
	"github.com/opengovern/og-describer-fly/discovery/describers"
	model "github.com/opengovern/og-describer-fly/discovery/pkg/models"
	"github.com/opengovern/og-describer-fly/discovery/provider"
	"github.com/opengovern/og-describer-fly/global"
)

var ResourceTypes = map[string]model.ResourceType{

	"Fly/App": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Fly/App",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByFly(describers.ListApps),
		GetDescriber:    provider.DescribeSingleByFly(describers.GetApp),
	},

	"Fly/Machine": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Fly/Machine",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByFly(describers.ListMachines),
		GetDescriber:    provider.DescribeSingleByFly(describers.GetMachine),
	},

	"Fly/Volume": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Fly/Volume",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByFly(describers.ListVolumes),
		GetDescriber:    provider.DescribeSingleByFly(describers.GetVolume),
	},

	"Fly/Secret": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Fly/Secret",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByFly(describers.ListSecrets),
		GetDescriber:    nil,
	},
}
