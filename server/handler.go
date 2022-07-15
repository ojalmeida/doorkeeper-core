package server

import (
	"context"
	"github.com/google/uuid"
	"github.com/ojalmeida/doorkeeper-core/config"
	"github.com/ojalmeida/doorkeeper-core/log"
	"net/http"
	"regexp"
	"time"
)

type Handler struct {
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	id := uuid.NewString()

	ctxWithID := context.WithValue(r.Context(), "request-id", id)
	ctxWithTimeout, cancel := context.WithTimeout(ctxWithID, time.Duration(config.Config.Timeout.Processing)*time.Second)

	defer cancel()
	r = r.Clone(ctxWithTimeout)

	for path, middleware := range Middlewares {

		if regexp.MustCompile(path).MatchString(r.URL.RequestURI()) {

			init := time.Now()

			proceed, err := middleware.Handle(r, w)

			if err != nil {
				log.Error.Printf("middleware %s returned an error: %s\n", middleware.Name(), err.Error())
			}

			// $uuid $method $uri $processing-time $addon
			log.Info.Printf("%s %s %s %d %s\n", id, r.Method, r.URL.RequestURI(), time.Now().Sub(init).Milliseconds(), middleware.Name())

			if !proceed {
				return
			}

		}

	}

	for path, plugin := range Plugins {

		if regexp.MustCompile(path).MatchString(r.URL.RequestURI()) {

			init := time.Now()

			err := plugin.Handle(r, w)

			if err != nil {
				log.Error.Println(err.Error())
			}

			// $uuid $method $uri $processing-time $addon
			log.Info.Printf("%s %s %s %d %s\n", id, r.Method, r.URL.RequestURI(), time.Now().Sub(init).Milliseconds(), plugin.Name())

			return

		}

	}

}
