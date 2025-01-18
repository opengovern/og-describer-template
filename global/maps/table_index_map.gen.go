package maps

import (
	"github.com/opengovern/og-describer-fly/discovery/pkg/es"
)

var ResourceTypesToTables = map[string]string{
  "Fly/App": "fly_app",
  "Fly/Machine": "fly_machine",
  "Fly/Volume": "fly_volume",
  "Fly/Secret": "fly_secret",
}

var ResourceTypeToDescription = map[string]interface{}{
  "Fly/App": opengovernance.App{},
  "Fly/Machine": opengovernance.Machine{},
  "Fly/Volume": opengovernance.Volume{},
  "Fly/Secret": opengovernance.Secret{},
}

var TablesToResourceTypes = map[string]string{
  "fly_app": "Fly/App",
  "fly_machine": "Fly/Machine",
  "fly_volume": "Fly/Volume",
  "fly_secret": "Fly/Secret",
}
