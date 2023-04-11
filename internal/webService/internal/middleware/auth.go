package middleware

import (
	conf "serverMonitor/internal/webService/pkg/config"

	"github.com/gin-gonic/gin"
)

func Auth()gin.HandlerFunc{
	return func(c *gin.Context){
		conf.Logger.Infoln("认证成功")
	}
}
