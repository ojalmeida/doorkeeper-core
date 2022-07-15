package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var Config Configuration

func LoadConfig(path string) {

	configFile, err := os.Open(path)

	if err != nil {
		panic(fmt.Errorf("unable to open configuration file : %w", err))
	}

	data, err := io.ReadAll(configFile)

	if err != nil {
		panic(fmt.Errorf("unable to read configuration file : %w", err))
	}

	err = json.Unmarshal(data, &Config)

	if err != nil {
		panic(fmt.Errorf("unable to unmarshal configuration file : %w", err))
	}

}
