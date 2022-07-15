package doorkeeper_core

import (
	"github.com/ojalmeida/doorkeeper-core/addons"
	"github.com/ojalmeida/doorkeeper-core/config"
	"github.com/ojalmeida/doorkeeper-core/log"
	"github.com/ojalmeida/doorkeeper-core/server"
	"os"
)

var configLoaded = false

func Start() {

	if !configLoaded {

		panic("file config not set")

	}

	server.Start()

}

func BindPlugin(path string, plugin addons.Plugin) {

	plugin.SetInfoLogger(log.Info)
	plugin.SetWarnLogger(log.Warn)
	plugin.SetErrorLogger(log.Error)

	server.Plugins[path] = plugin

}

func BindMiddleware(path string, middleware addons.Middleware) {

	middleware.SetInfoLogger(log.Info)
	middleware.SetWarnLogger(log.Warn)
	middleware.SetErrorLogger(log.Error)

	server.Middlewares[path] = middleware

}

func SetConfigFile(file *os.File) {

	config.LoadConfig(file.Name())
	log.LoadLoggers()
	configLoaded = true

}
