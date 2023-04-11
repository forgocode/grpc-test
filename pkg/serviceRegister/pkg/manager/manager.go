package manager

import (
	"serverMonitor/pkg/typed"
	"sync"
)

type MicroServiceRoot struct {
	typed.MicroServiceRoot
	sync.RWMutex
}

func InitServiceRoot(service *typed.MicroService) *MicroServiceRoot {
	m := &MicroServiceRoot{}
	m.OtherSerices = make(map[string][]typed.Endpoint)
	m.RWMutex = sync.RWMutex{}
	m.Name = service.ServiceName
	m.Endpoints = service.Endpoints

	return m
}

func (m *MicroServiceRoot) UpsertService(s typed.MicroService) {
	m.Lock()
	defer m.Unlock()
	m.OtherSerices[s.ServiceName] = s.Endpoints
}

func (m *MicroServiceRoot) DeleteService(s typed.MicroService) {
	m.Lock()
	defer m.Unlock()
	delete(m.OtherSerices, s.ServiceName)
}

func (m *MicroServiceRoot) GetServiceEndpointsByName(name string) typed.Endpoint {
	m.Lock()
	defer m.Unlock()
	ep := m.OtherSerices[name]
	return ep[0]
}
