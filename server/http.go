package server

import (
	"github.com/ojalmeida/doorkeeper-core/config"
	"net/http"
	"time"
)

var s *http.Server

func init() {

	s = &http.Server{}

	s.SetKeepAlivesEnabled(false)

	s.ReadTimeout = time.Duration(config.Config.Timeout.Read) * time.Second
	s.WriteTimeout = time.Duration(config.Config.Timeout.Write) * time.Second

	s.Handler = Handler{}
	s.Addr = "127.0.0.1:8080"

}

func Start() {

	err := s.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {

		panic(err.Error())

	}

}
