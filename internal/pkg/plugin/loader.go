package plugin

import (
	"fmt"
	"io/ioutil"
	"os"
	"plugin"
	"strings"

	"github.com/rs/zerolog/log"
)

// Loader handles the loading and managing of server plugins.
type Loader struct {
	plugins []Plugin
}

// NewPluginLoader initializes a new PluginManager with no loaded plugins.
func NewPluginLoader() Loader {
	return Loader{plugins: make([]Plugin, 0)}
}

// FetchPlugins reads the ./plugins folder and looks for valid shared object files.
func (loader *Loader) FetchPlugins() {
	if _, err := os.Stat("./plugins"); os.IsNotExist(err) {
		os.Mkdir("plugins", 0755)
	}

	files, err := ioutil.ReadDir("plugins")
	if err != nil {
		log.Err(err).Msg("Failed to read plugin directory")
		return
	}

	for _, file := range files {
		// skip anything that isn't a .so file
		if !strings.HasSuffix(file.Name(), ".so") {
			continue
		}

		pl, err := plugin.Open(fmt.Sprintf("./plugins/%s", file.Name()))
		if err != nil {
			log.Err(err).Msgf("Error while loading '%s'", file.Name())
			continue
		}

		// instantiate plugin and append
		plugin := NewPlugin(pl)
		loader.plugins = append(loader.plugins, plugin)
		log.Info().Msgf("Found plugin '%s'", plugin.GetPluginInfo().Name)
	}

	log.Info().Msgf("Loaded %d plugins", len(loader.plugins))
}
