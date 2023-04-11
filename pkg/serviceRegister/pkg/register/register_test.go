package register

import (
	"serverMonitor/pkg/constant"
	"serverMonitor/pkg/typed"
	"testing"
	"time"
)

func Test_Register(t *testing.T) {
	config := &typed.ConfigYaml{}
	config.Etcd.EndPoints = []string{"192.168.0.100:2379"}
	logGrpcService := &typed.MicroService{ServiceName: constant.LogGrpcName}
	logGrpcService.Endpoints = append(logGrpcService.Endpoints, typed.Endpoint{IP: "192.168.0.101", Port: 10000})

	register(logGrpcService, config)
	go discoverServices(logGrpcService.ServiceName, nil, config)

	time.Sleep(5 * time.Second)
	logGrpcService.Endpoints = append(logGrpcService.Endpoints, typed.Endpoint{IP: "192.168.0.100", Port: 10000})
	register(logGrpcService, config)
	time.Sleep(5 * time.Second)

	deRegister(logGrpcService.ServiceName, config)
	time.Sleep(2 * time.Second)
}
