package server

import "github.com/ojalmeida/doorkeeper-core/addons"

var (
	Plugins     = map[string]addons.Plugin{}
	Middlewares = map[string]addons.Middleware{}
)
