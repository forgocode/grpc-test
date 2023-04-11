package serviceService

import (
	"serverMonitor/internal/webService/pkg/database"
	"serverMonitor/pkg/typed"
)

func ListService() []typed.Service {
	var services []typed.Service
	database.DB.Find(&services)
	return services
}
