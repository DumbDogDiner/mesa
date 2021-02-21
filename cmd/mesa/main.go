package main

import (
	"os"

	"github.com/dumbdogdiner/mesa/internal/pkg/plugin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	loader := plugin.NewPluginLoader()
	loader.FetchPlugins()
}
