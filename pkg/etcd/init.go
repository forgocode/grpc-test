package etcd

import (
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func GetClient(eps []string, timeOut time.Duration) *clientv3.Client {
	client, err := clientv3.New(GeneratedEtcdConfig(eps, timeOut))
	if err != nil {
		fmt.Printf("failed to get new etcd client, err:%+v\n", err)
		return nil
	}
	return client
}

func GeneratedEtcdConfig(eps []string, timeOut time.Duration) clientv3.Config {
	return clientv3.Config{Endpoints: eps, DialTimeout: timeOut}
}
