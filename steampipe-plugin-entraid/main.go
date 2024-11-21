package main

import (
	"github.com/opengovern/og-describer-entraid/steampipe-plugin-entraid/entraid"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: entraid.Plugin})
}
