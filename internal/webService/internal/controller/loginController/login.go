package LoginController

import (
	"context"
	"fmt"
	base "serverMonitor/internal/webService/internal/controller/baseController"
	"serverMonitor/internal/webService/internal/microService"
	"serverMonitor/pkg/constant"
	"serverMonitor/pkg/util"

	pb "serverMonitor/internal/logService/proto"

	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginInfo := make(map[string]string)
	err := c.BindJSON(&loginInfo)
	fmt.Printf("%+v\n", c)
	fmt.Println(loginInfo)
	if err != nil {
		fmt.Println(err)
		base.ResponseWithError(c, "error body")
	}

	if loginInfo["userName"] == "123456" && loginInfo["passWord"] == "123456" {
		endpoints := microservice.GetWebServiceServiceRoot().GetServiceEndpointsByName(constant.LogGrpcName)
		conn, err := grpc.Dial(util.GenerateGrpcClientStr(endpoints), grpc.WithInsecure())
		if err != nil {
			fmt.Printf("can't connect to %s", constant.LogGrpcName)
		}
		defer conn.Close()

		client := pb.NewRecordLogClient(conn)
		client.RecordLogMsg(context.Background(), &pb.Msg{User: loginInfo["userName"], Module: constant.LoginModule, Msg: fmt.Sprintf("user %s login successfully", loginInfo["userName"])})

		base.ResponseWithJson(c, "login successfully")
		return
	}
	base.ResponseWithError(c, "user or passwd error, please try again")
}

func Regiter(c *gin.Context) {

}
