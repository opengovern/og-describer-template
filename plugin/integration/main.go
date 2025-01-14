package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/opengovern/og-describer-github/plugin/integration/configs"
	"github.com/opengovern/og-describer-github/plugin/integration/integration"
	"github.com/opengovern/opencomply/services/integration/integration-type/interfaces"
	"os"
)

func main() {
	i := integration.Integration{}
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Debug,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	var pluginMap = map[string]plugin.Plugin{
		configs.IntegrationTypeGithubAccount.String(): &interfaces.IntegrationTypePlugin{Impl: &i},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: interfaces.HandshakeConfig,
		Plugins:         pluginMap,
		Logger:          logger,
	})
}
