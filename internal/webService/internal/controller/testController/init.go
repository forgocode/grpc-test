package testController

import (
	"fmt"
	base "serverMonitor/internal/webService/internal/controller/baseController"

	"github.com/gin-gonic/gin"
)

type testResult struct {
	User string `json:"user"`
	Role string `json:"role"`
	Age  int16  `json:"age"`
}

func TestGetApi(c *gin.Context) {
	user := &testResult{User: "root", Role: "admin", Age: 25}
	base.ResponseWithJson(c, user)
}

func TestDeleteApi(c *gin.Context) {
	user := &testResult{User: "root", Role: "admin", Age: 25}
	base.ResponseWithJson(c, user)
}

func TestPostApi(c *gin.Context) {
	user := &testResult{}
	err := c.BindJSON(user)
	if err != nil {
		base.ResponseWithError(c, "error body")
		return
	}
	fmt.Printf("user: %+v\n", user)
	base.ResponseWithJson(c, user)
}

func TestPutApi(c *gin.Context) {
	user := &testResult{User: "root", Role: "admin", Age: 25}
	base.ResponseWithJson(c, user)
}
