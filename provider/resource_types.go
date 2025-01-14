package provider

import (
	model "github.com/opengovern/og-describer-fly/pkg/sdk/models"
	"github.com/opengovern/og-describer-fly/provider/configs"
	"github.com/opengovern/og-describer-fly/provider/describer"
)

var ResourceTypes = map[string]model.ResourceType{

	"Fly/App": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Fly/App",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByFly(describer.ListApps),
		GetDescriber:    DescribeSingleByFly(describer.GetApp),
	},

	"Fly/Machine": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Fly/Machine",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByFly(describer.ListMachines),
		GetDescriber:    nil,
	},

	"Fly/Volume": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Fly/Volume",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByFly(describer.ListVolumes),
		GetDescriber:    nil,
	},

	"Fly/Secret": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Fly/Secret",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByFly(describer.ListSecrets),
		GetDescriber:    nil,
	},
}
