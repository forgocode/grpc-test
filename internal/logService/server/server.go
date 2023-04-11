package server

import (
	"context"
	"fmt"
	"net"
	conf "serverMonitor/internal/logService/config"
	pb "serverMonitor/internal/logService/proto"
	"serverMonitor/pkg/config"
	"serverMonitor/pkg/constant"
	"serverMonitor/pkg/mongo"
	"serverMonitor/pkg/serviceRegister/pkg/manager"
	"serverMonitor/pkg/serviceRegister/pkg/register"
	"serverMonitor/pkg/typed"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedRecordLogServer
}

type HandleLogInfo struct{
	Time time.Time
	User  string
	Module string
	Msg   string
}

func (s *Server) RecordLogMsg(ctx context.Context, msg *pb.Msg) (*pb.Reponse, error) {
	fmt.Printf("receive msg from client, msg info: %+v\n", msg)
	//TODO入库
	newMsg := HandleLogInfo{
		Time: time.Now(),
		User: msg.GetUser(),
		Module: msg.GetModule(),
		Msg: msg.GetMsg(),
	}
	c := mongo.NewMgo(constant.MongoHandleLogConnnection)
	err := c.InsertOne(newMsg)

	return &pb.Reponse{Result: 1}, err
}

var serviceRoot *manager.MicroServiceRoot = nil

func StartLogRpc() {
	fmt.Println("log rpc start!")
	logGrpc := &typed.MicroService{}
	logGrpc.ServiceName = constant.LogGrpcName

	logGrpc.Endpoints = append(logGrpc.Endpoints, typed.Endpoint{IP: "127.0.0.1", Port: 20008})
	lis, err := net.Listen("tcp", ":20008")
	if err != nil {
		fmt.Printf("failed to listen 20008, err:%+v\n", err)
		return
	}

	s := grpc.NewServer()
	isConfigUpdated := make(chan int)
	conf.LogServiceConfig, conf.Logger = config.ConfigInit(conf.ConfigPath, conf.LogLevel, isConfigUpdated)
	fmt.Printf("config :%+v\n", conf.LogServiceConfig)
	register.Register(logGrpc, conf.LogServiceConfig)
	//注册到服务中心, 感觉没有必要
	pb.RegisterRecordLogServer(s, &Server{})
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
			register.DiscoverServices(svc, outPut, conf.LogServiceConfig)
		}()

	}
}

func GetLogServiceRoot() *manager.MicroServiceRoot {
	if serviceRoot != nil {
		return serviceRoot
	}
	logGrpc := &typed.MicroService{}
	logGrpc.ServiceName = constant.LogGrpcName
	logGrpc.Endpoints = append(logGrpc.Endpoints, typed.Endpoint{IP: "127.0.0.1", Port: 20008})
	serviceRoot = manager.InitServiceRoot(logGrpc)
	return serviceRoot
}

func handleMicroService(svcs chan *typed.MicroService) {
	for {
		select {
		case svc := <-svcs:
			switch svc.Action {
			case typed.MicroServiceAdd, typed.MicroServiceUpdate:
				GetLogServiceRoot().UpsertService(*svc)
			case typed.MicroServiceDel:
				GetLogServiceRoot().DeleteService(*svc)
			}
		}
	}
}
