package serviceController

import (
	base "serverMonitor/internal/webService/internal/controller/baseController"
	"serverMonitor/internal/webService/internal/monitor"
	asset "serverMonitor/pkg/typed"
	"time"

	"github.com/gin-gonic/gin"
)

func ListService(c *gin.Context) {
	services := monitor.ListService()
	base.ResponseWithJson(c, services)
}

func AddService(c *gin.Context) {
	service := &asset.Service{}
	service.CreateTime = time.Now()
	service.LastCheckTime = time.Now()
	err := c.BindJSON(service)
	if err != nil {
		base.ResponseWithError(c, "error body")
		return
	}
	err = monitor.AddService(service)
	if err != nil {
		base.ResponseWithError(c, err)
		return
	}
	base.ResponseWithJson(c, "post service successfully")
}

func UpdateService(c *gin.Context) {
	service := &asset.Service{}
	err := c.BindJSON(service)
	if err != nil {
		base.ResponseWithError(c, "error body")
		return
	}
	err = monitor.UpdateService(service)
	if err != nil {
		base.ResponseWithError(c, err)
		return
	}

	base.ResponseWithJson(c, "update service successfully")
}

func DeleteService(c *gin.Context) {
	service := &asset.Service{}
	err := c.BindJSON(service)
	if err != nil {
		base.ResponseWithError(c, "error body")
		return
	}
	err = monitor.DeleteService(service)
	if err != nil {
		base.ResponseWithError(c, err)
		return
	}
	base.ResponseWithJson(c, "delete service successfully")
}
