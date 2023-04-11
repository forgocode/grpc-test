package typed

import (
	"time"
)

type serviceStatus string

const (
	ServiceReady    serviceStatus = "serviceReady"
	ServiceNotReady serviceStatus = "serviceNotReady"
)

type Service struct {
	Name          string
	CreateTime    time.Time
	LastCheckTime time.Time
	Url           string
	Status        serviceStatus
}

func (o *Service) Copy(in *Service) {
	o.Name = in.Name
	o.Url = in.Url
}

func (Service) TableName() string {
	return "service"
}
