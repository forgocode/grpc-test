package server

import (
	"context"
	"fmt"
	"net"
	conf "serverMonitor/internal/userService/config"
	pb "serverMonitor/internal/userService/proto"
	"serverMonitor/pkg/constant"

	"serverMonitor/pkg/config"
	"serverMonitor/pkg/serviceRegister/pkg/manager"
	"serverMonitor/pkg/serviceRegister/pkg/register"
	"serverMonitor/pkg/typed"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServer
}

var serviceRoot *manager.MicroServiceRoot = nil

func (s *Server) UserManager(ctx context.Context, msg *pb.UserInfo) {
	fmt.Printf("user info: %+v\n", msg)
	user := &typed.User{
		Name:   msg.Name,
		Passwd: msg.Passwd,
	}
	switch msg.Type {
	case "login":
		fmt.Printf("login\n")
		Login(*user)
	case "register":
		fmt.Printf("register\n")
		Regiter(*user)

	}
}

func StartUserRpc() {
	fmt.Println("user grpc start!")
	userGrpc := &typed.MicroService{}
	userGrpc.ServiceName = constant.UserGrpcName

	userGrpc.Endpoints = append(userGrpc.Endpoints, typed.Endpoint{IP: "127.0.0.1", Port: 20009})
	lis, err := net.Listen("tcp", ":20009")
	if err != nil {
		fmt.Printf("failed to listen 20008, err:%+v\n", err)
		return
	}

	s := grpc.NewServer()
	isConfigUpdated := make(chan int)
	conf.UserServiceConfig, conf.Logger = config.ConfigInit(conf.ConfigPath, conf.LogLevel, isConfigUpdated)
	fmt.Printf("config :%+v\n", conf.UserServiceConfig)
	register.Register(userGrpc, conf.UserServiceConfig)
	//注册到服务中心, 感觉没有必要
	//pb.RegisterUserServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to start grpc server, err: %+v\n", err)
	}
	outPut := make(chan *typed.MicroService, 100)
	go handleMicroService(outPut)
	for _, svc := range constant.MicroServiceList {
		if svc == constant.LogGrpcName {
			continue
		}
		go func() {
			register.DiscoverServices(svc, outPut, conf.UserServiceConfig)
		}()

	}
}

func GetUserServiceRoot() *manager.MicroServiceRoot {
	if serviceRoot != nil {
		return serviceRoot
	}
	userGrpc := &typed.MicroService{}
	userGrpc.ServiceName = constant.UserGrpcName
	userGrpc.Endpoints = append(userGrpc.Endpoints, typed.Endpoint{IP: "127.0.0.1", Port: 20009})
	serviceRoot = manager.InitServiceRoot(userGrpc)
	return serviceRoot
}

func handleMicroService(svcs chan *typed.MicroService) {
	for {
		select {
		case svc := <-svcs:
			switch svc.Action {
			case typed.MicroServiceAdd, typed.MicroServiceUpdate:
				GetUserServiceRoot().UpsertService(*svc)
			case typed.MicroServiceDel:
				GetUserServiceRoot().DeleteService(*svc)
			}
		}
	}
}
