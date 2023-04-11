package baseController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseGin gin.Engine

func ResponseWithJson(c *gin.Context, pyload interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{"code": http.StatusOK, "result": pyload})
}

func ResponseWithError(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{"code": http.StatusServiceUnavailable, "message": msg})
}
