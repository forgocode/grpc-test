package config

import (
	"go.uber.org/zap"
	"serverMonitor/pkg/typed"
)

var LogServiceConfig *typed.ConfigYaml
var Logger *zap.SugaredLogger
var LogLevel = zap.NewAtomicLevel()

const ConfigPath = "/root/goWorkspace/serverMonitor/config/logService/config.yaml"
