package steampipe

import (
	"github.com/opengovern/og-describer-fly/pkg/sdk/es"
)

var Map = map[string]string{
  "Fly/App": "fly_app",
  "Fly/Machine": "fly_machine",
  "Fly/Volume": "fly_volume",
  "Fly/Secret": "fly_secret",
}

var DescriptionMap = map[string]interface{}{
  "Fly/App": opengovernance.App{},
  "Fly/Machine": opengovernance.Machine{},
  "Fly/Volume": opengovernance.Volume{},
  "Fly/Secret": opengovernance.Secret{},
}

var ReverseMap = map[string]string{
  "fly_app": "Fly/App",
  "fly_machine": "Fly/Machine",
  "fly_volume": "Fly/Volume",
  "fly_secret": "Fly/Secret",
}
