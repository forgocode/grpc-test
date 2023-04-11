package serviceService

import (
	"fmt"
	"serverMonitor/internal/webService/pkg/database"
	"serverMonitor/pkg/typed"
)

func InsertService(service *typed.Service) error {
	result := database.DB.Create(service)
	if result.Error != nil {
		return fmt.Errorf("failed to insert service, err:%+v\n", result.Error)
	}
	return nil
}
