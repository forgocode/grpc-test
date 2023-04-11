package config

import (
	"testing"
	"time"
)

func TestReadConfig(t *testing.T) {
	configPath := "./../config.yaml"
	ReadConfig(configPath)
	for {
		// fmt.Printf("%+v\n", Config)
		Logger.Infof("%+v\n", Config)
		Logger.Errorf("%+v\n", Config)
		time.Sleep(10 * time.Second)
	}
}
