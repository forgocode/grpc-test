package register

import (
	"context"
	"encoding/json"
	"fmt"
	"serverMonitor/pkg/etcd"
	"serverMonitor/pkg/typed"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func register(service *typed.MicroService, conf *typed.ConfigYaml) {
	etcdClient := etcd.GetClient(conf.Etcd.EndPoints, 5*time.Second)
	defer etcdClient.Close()
	data, err := json.Marshal(service)
	if err != nil {
		fmt.Printf("failed marshal service, err: %+v\n", err)
	}
	//TODO etcd 连接不上卡住不报错 不超时 status
	_, err = etcdClient.Put(context.Background(), "/"+service.ServiceName, string(data), clientv3.WithPrevKV())
	if err != nil {
		fmt.Printf("failed to register micro service {%s}, err: %+v\n", service.ServiceName, err)
	}
	fmt.Println("register successfully!")
}

func deRegister(serviceName string, conf *typed.ConfigYaml) {
	etcdClient := etcd.GetClient(conf.Etcd.EndPoints, 5*time.Second)
	defer etcdClient.Close()
	_, err := etcdClient.Delete(context.Background(), "/"+serviceName)
	if err != nil {
		fmt.Printf("failed to delete register service {%s}, err: %+v\n", serviceName, err)
	}

}

//考虑做成接口，每个服务都实现这个接口
func discoverServices(serviceName string, service chan *typed.MicroService, conf *typed.ConfigYaml) {
	etcdClient := etcd.GetClient(conf.Etcd.EndPoints, 5*time.Second)
	go func() {
		result, err := etcdClient.Get(context.Background(), "/"+serviceName)
		for _, s := range result.Kvs {
			svc := &typed.MicroService{}
			err = json.Unmarshal(s.Value, svc)
			if err != nil {
				fmt.Printf("add/update Unmarshal service failed, err:%+v\n", err)
				continue
			}
			svc.Action = typed.MicroServiceAdd
			service <- svc
		}

	}()

	serviceChan := etcdClient.Watch(context.Background(), "/"+serviceName)
	for {
		select {
		case services := <-serviceChan:
			for _, s := range services.Events {
				switch s.Type {
				case mvccpb.PUT:
					fmt.Printf("service add/update, key: %s, service:%+v\n", string(s.Kv.Key), string(s.Kv.Value))
					svc := &typed.MicroService{}
					err := json.Unmarshal(s.Kv.Value, svc)
					if err != nil {
						fmt.Printf("add/update Unmarshal service failed, err:%+v\n", err)
						continue
					}
					svc.Action = typed.MicroServiceAdd
					service <- svc
				case mvccpb.DELETE:
					fmt.Printf("service delete, key: %s, service:%+v\n", string(s.Kv.Key), string(s.Kv.Value))
					svc := &typed.MicroService{}
					err := json.Unmarshal(s.Kv.Value, svc)
					if err != nil {
						fmt.Printf("delete Unmarshal service failed, err:%+v\n", err)
						continue
					}
					svc.Action = typed.MicroServiceDel
					service <- svc
				}
			}
		}
	}
}

type MicroService typed.MicroService

func Register(m *typed.MicroService, conf *typed.ConfigYaml) {
	register(m, conf)
}

func DeRegister(m *typed.MicroService, conf *typed.ConfigYaml) {
	deRegister(m.ServiceName, conf)
}

func DiscoverServices(serviceName string, outPut chan *typed.MicroService, conf *typed.ConfigYaml) {
	discoverServices(serviceName, outPut, conf)
}
