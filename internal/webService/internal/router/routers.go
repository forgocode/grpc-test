package router

import (
	"serverMonitor/internal/webService/internal/controller/loginController"
	"serverMonitor/internal/webService/internal/controller/serviceController"
	"serverMonitor/internal/webService/internal/controller/testController"
	"serverMonitor/internal/webService/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Start() {
	//修改日志级别
	gin.SetMode(gin.ReleaseMode)
	//禁止打印到终端
	// gin.DefaultWriter = io.Discard

	r := gin.Default()
	//全局中间件
	r.Use(middleware.Logger(), gin.Recovery())

	r.POST("/login", LoginController.Login)
	r.POST("/register", LoginController.Regiter)

	r.GET("/home", middleware.Auth(), serviceController.ListService)

	r.GET("/service", serviceController.ListService)
	r.POST("/service", serviceController.AddService)
	r.PUT("/service", serviceController.UpdateService)
	r.DELETE("/service", serviceController.DeleteService)

	r.GET("/test", testController.TestGetApi)
	r.DELETE("/test", testController.TestDeleteApi)
	r.POST("/test", testController.TestPostApi)
	r.PUT("/test", testController.TestPutApi)

	r.Run(":8000")
}
