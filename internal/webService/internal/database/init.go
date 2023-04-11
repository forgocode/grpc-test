package database

import (
	"fmt"
	conf "serverMonitor/internal/webService/pkg/config"
	"serverMonitor/internal/webService/pkg/database"
	"serverMonitor/pkg/typed"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//TODO(sync once)
func Init(user, passwd, url, database string, port uint16) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, url, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("open mysql err, err:  %+v", err))
	}
	return db
}

func PingDb() error {
	client, err := database.DB.DB()

	if err != nil {
		return err
	}
	if client.Ping() != nil {
		return fmt.Errorf("can not to connect to mysql")
	}
	return nil
}

func MigratorTable(tableName string) {
	switch tableName {
	case "service":
		database.DB.AutoMigrate(&typed.Service{})
	}
}

func DbInit() {
	database.DB = Init(conf.WebserviceConfig.Db.User, conf.WebserviceConfig.Db.Passwd, conf.WebserviceConfig.Db.URL, conf.WebserviceConfig.Db.DbName, conf.WebserviceConfig.Db.Port)
}
