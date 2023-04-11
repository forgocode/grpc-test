package config

import (
	"go.uber.org/zap"
	"serverMonitor/pkg/typed"
)

var UserServiceConfig *typed.ConfigYaml
var Logger *zap.SugaredLogger
var LogLevel = zap.NewAtomicLevel()

const ConfigPath = "/root/goWorkspace/serverMonitor/config/userService/config.yaml"
