package monitor

import (
	"fmt"
	"os/exec"
	"serverMonitor/internal/webService/internal/service/serviceService"
	"serverMonitor/pkg/typed"
	"sync"
	"time"
)

type ServiceRoot struct {
	Services map[string]*typed.Service
	sync.Mutex
}

var serviceRoot *ServiceRoot

func checkServiceStatus() {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	for name, service := range serviceRoot.Services {
		if isServiceReady(service) {
			serviceRoot.Services[name].Status = typed.ServiceReady
			continue
		}
		serviceRoot.Services[name].Status = typed.ServiceNotReady
	}
}

func NewTimer() {
	timer := time.NewTicker(time.Minute * 5)
	for {
		select {
		case <-timer.C:
			checkServiceStatus()
		}
	}
}

func isServiceReady(service *typed.Service) bool {
	cmd := exec.Command("curl", service.Url)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("service: %s is not reachable, err: %+v\n", service.Name, err)
		return false
	}
	return true

}

func AddService(service *typed.Service) error {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	if _, ok := serviceRoot.Services[service.Name]; ok {
		return fmt.Errorf("service :%+v is existed\n", service.Name)
	}
	serviceRoot.Services[service.Name] = service
	err := serviceService.InsertService(service)
	return err
}

func DeleteService(service *typed.Service) error {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	delete(serviceRoot.Services, service.Name)
	err := serviceService.DeleteService(service)
	return err
}

func UpdateService(newService *typed.Service) error {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	var oldService *typed.Service
	if _, ok := serviceRoot.Services[newService.Name]; !ok {
		return fmt.Errorf("new service is not exist, update failed")
	}
	oldService = serviceRoot.Services[newService.Name]
	delete(serviceRoot.Services, oldService.Name)
	serviceRoot.Services[newService.Name] = newService
	err := serviceService.UpdateService(oldService, newService)
	return err
}

func ListService() []typed.Service {
	serviceRoot.Lock()
	defer serviceRoot.Unlock()
	var services []typed.Service
	for _, s := range serviceRoot.Services {
		services = append(services, *s)
	}
	return services
}

func InitServiceRoot() {
	serviceRoot = &ServiceRoot{
		Services: make(map[string]*typed.Service),
	}
	services := serviceService.ListService()
	for _, s := range services {
		serviceRoot.Services[s.Name] = &s
	}
	fmt.Printf("service root: %+v\n", serviceRoot)
}

func getServiceRoot() *ServiceRoot {
	if serviceRoot != nil {
		return serviceRoot
	}
	return nil
}
