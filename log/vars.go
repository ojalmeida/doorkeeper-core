package log

import (
	"fmt"
	"github.com/ojalmeida/doorkeeper-core/config"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func LoadLoggers() {

	infoFile, err := os.OpenFile(config.Config.Logging.Info.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)

	if err != nil {
		panic(fmt.Errorf("unable to open info logging file : %w", err))
	}

	warnFile, err := os.OpenFile(config.Config.Logging.Warn.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)

	if err != nil {
		panic(fmt.Errorf("unable to open warn logging file : %w", err))
	}

	errorFile, err := os.OpenFile(config.Config.Logging.Error.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)

	if err != nil {
		panic(fmt.Errorf("unable to open error logging file : %w", err))
	}

	Info = log.New(infoFile, config.Config.Logging.Info.Prefix, log.Ldate|log.Ltime)
	Warn = log.New(warnFile, config.Config.Logging.Warn.Prefix, log.Ldate|log.Ltime)
	Error = log.New(errorFile, config.Config.Logging.Error.Prefix, log.Ldate|log.Ltime)

}
