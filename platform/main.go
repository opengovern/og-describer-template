package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"os"
)


func main() {
	i := Integration{}
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Debug,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	var pluginMap = map[string]plugin.Plugin{
		IntegrationTypeGithubAccount.String(): &IntegrationTypePlugin{Impl: &i},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         pluginMap,
		Logger:          logger,
	})
}
