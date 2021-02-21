package plugin

import (
	"plugin"

	"github.com/dumbdogdiner/mesa/pkg/shared/api"
	"github.com/rs/zerolog/log"
)

// Plugin is a wrapper for the built-in Go plugin, providing quick
// access to plugin methods.
type Plugin struct {
	internal *plugin.Plugin
}

// NewPlugin creates a new plugin and returns it.
func NewPlugin(internal *plugin.Plugin) Plugin {
	return Plugin{internal}
}

// GetPluginInfo returns the info of this plugin.
func (plugin *Plugin) GetPluginInfo() *api.PluginInfo {
	symbol, err := plugin.internal.Lookup("GetPluginInfo")
	if err != nil {
		log.Err(err).Msg("Failed to get plugin information")
		return nil
	}
	return symbol.(func() *api.PluginInfo)()
}
