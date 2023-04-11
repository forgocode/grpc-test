package serviceService

import (
	"fmt"
	"serverMonitor/internal/webService/pkg/database"
	"serverMonitor/pkg/typed"
)

func UpdateService(oldService, newService *typed.Service) error {
	database.DB.First(oldService)
	oldService.Copy(newService)
	result := database.DB.Save(oldService)
	if result.Error != nil {
		fmt.Errorf("failed to update service{%s}, err:%+v\n", newService.Name, result.Error)
	}
	return nil
}
