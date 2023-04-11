package serviceService

import (
	"fmt"
	"serverMonitor/internal/webService/pkg/database"
	"serverMonitor/pkg/typed"
)

func DeleteService(service *typed.Service) error {
	result := database.DB.Delete(service)
	if result.Error != nil {
		fmt.Errorf("failed to delete service, err:%+v\n", result.Error)
	}
	if result.RowsAffected > 1 {
		fmt.Errorf("delete service{%s} more than 1", service.Name)
	}
	if result.RowsAffected == 0 {
		fmt.Errorf("service {%s} is not exist", service.Name)
	}
	return nil
}
