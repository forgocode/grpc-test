package main

import (
	"net/http"
	_ "net/http/pprof"
	conf "serverMonitor/internal/webService/pkg/config"
	"serverMonitor/internal/webService/pkg/start"
	"serverMonitor/pkg/config"
)

func main() {
	isConfigUpdated := make(chan int)
	conf.WebserviceConfig, conf.Logger = config.ConfigInit(conf.ConfigPath, conf.LogLevel, isConfigUpdated)
	go func() {
		select {
		case <-isConfigUpdated:
			conf.WebserviceConfig = config.ConfigReload()
		}
	}()
	go func() {
		conf.Logger.Infof("start to http pprof\n")
		http.ListenAndServe("0.0.0.0:8081", nil)
	}()
	start.Start()
	select {}
}
