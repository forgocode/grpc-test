package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"serverMonitor/pkg/log"
	"serverMonitor/pkg/typed"
)

func ConfigInit(configPath string, logLevel zapcore.LevelEnabler, ch chan int) (*typed.ConfigYaml, *zap.SugaredLogger) {
	conf := ReadConfig(configPath, ch)
	logger := log.InitLogger(conf.Log.FileName, logLevel)
	return conf, logger
}
func ReadConfig(configPath string, ch chan int) *typed.ConfigYaml {

	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("failed to read config, err: %+v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config change, reload config\n")
		ch <- 1
	})
	return reloadConfig()
}

func ConfigReload() *typed.ConfigYaml {
	return reloadConfig()
}

func reloadConfig() *typed.ConfigYaml {
	config := &typed.ConfigYaml{
		Db: typed.DbConfig{
			DbName:   viper.GetString("db.dbName"),
			Database: viper.GetString("db.database"),
			URL:      viper.GetString("db.url"),
			Port:     viper.GetUint16("db.port"),
			User:     viper.GetString("db.user"),
			Passwd:   viper.GetString("db.passwd"),
		},
		Log: typed.LogConfig{
			FileName: viper.GetString("log.fileName"),
			Level:    viper.GetString("log.level"),
		},
		Etcd: typed.EtcdConfig{
			EndPoints: viper.GetStringSlice("etcd.endpoints"),
		},
	}
	return config
}
