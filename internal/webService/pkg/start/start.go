package start

import (
	"serverMonitor/internal/webService/internal/database"
	"serverMonitor/internal/webService/internal/microService"
	"serverMonitor/internal/webService/internal/monitor"
	"serverMonitor/internal/webService/internal/router"
)

func Start() {
	microservice.StartHandleService()
	database.DbInit()
	monitor.Start()
	router.Start()
}
