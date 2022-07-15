package addons

import (
	"log"
	"net/http"
)

type Plugin interface {
	Handle(r *http.Request, w http.ResponseWriter) error
	SetInfoLogger(logger *log.Logger)
	SetWarnLogger(logger *log.Logger)
	SetErrorLogger(logger *log.Logger)
	Name() string
}

type Middleware interface {
	Handle(r *http.Request, w http.ResponseWriter) (bool, error)
	SetInfoLogger(logger *log.Logger)
	SetWarnLogger(logger *log.Logger)
	SetErrorLogger(logger *log.Logger)
	Name() string
}
