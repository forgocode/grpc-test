package config

import (
	"go.uber.org/zap"
	"serverMonitor/pkg/typed"
)

var WebserviceConfig *typed.ConfigYaml
var Logger *zap.SugaredLogger
var LogLevel = zap.NewAtomicLevel()

const ConfigPath = "/root/goWorkspace/serverMonitor/config/webService/config.yaml"
