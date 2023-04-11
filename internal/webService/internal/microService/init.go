package microservice

import (
	"serverMonitor/internal/webService/pkg/config"
	"serverMonitor/pkg/constant"
	"serverMonitor/pkg/serviceRegister/pkg/manager"
	"serverMonitor/pkg/serviceRegister/pkg/register"
	"serverMonitor/pkg/typed"
)

var serviceRoot *manager.MicroServiceRoot

func GetWebServiceServiceRoot() *manager.MicroServiceRoot {
	if serviceRoot != nil {
		return serviceRoot
	}
	logGrpc := &typed.MicroService{}
	logGrpc.ServiceName = constant.LogGrpcName
	logGrpc.Endpoints = append(logGrpc.Endpoints, typed.Endpoint{IP: "192.168.0.100", Port: 20008})
	serviceRoot = manager.InitServiceRoot(logGrpc)
	return serviceRoot
}

func handleMicroService(svcs chan *typed.MicroService) {
	for {
		select {
		case svc := <-svcs:
			switch svc.Action {
			case typed.MicroServiceAdd, typed.MicroServiceUpdate:
				GetWebServiceServiceRoot().UpsertService(*svc)
			case typed.MicroServiceDel:
				GetWebServiceServiceRoot().DeleteService(*svc)
			}
		}
	}
}

func StartHandleService() {
	outPut := make(chan *typed.MicroService, 100)
	go handleMicroService(outPut)
	for _, svc := range constant.MicroServiceList {
		go func(serviceName string) {
			register.DiscoverServices(serviceName, outPut, config.WebserviceConfig)
		}(svc)
	}

}
