package main

import (
	"serverMonitor/internal/userService/server"
)

func main() {
	//启动user grpc
	server.StartUserRpc()
}
