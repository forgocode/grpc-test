package middleware

import (
	"time"

	conf "serverMonitor/internal/webService/pkg/config"

	"github.com/gin-gonic/gin"
)

func Logger()gin.HandlerFunc{
	return func(c *gin.Context){
		start := time.Now()
		host := c.RemoteIP()
		path := c.Request.URL.Path
		method := c.Request.Method
		c.Next()
		raw := c.Request.URL.RawQuery
		status := c.Writer.Status()
		conf.Logger.Infof("src ip: %s, path: %s, method: %s, consume: %+v, payload: %s, status: %d", host, path, method, time.Now().Sub(start), raw, status)
	}
}
