package main

import (
	"github.com/opengovern/og-describer-render/steampipe-plugin-render/render"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: render.Plugin})
}
