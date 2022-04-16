package etcdx

import (
	"context"
	"devops/configrue/api/internal/config"
	"errors"
	etcd "go.etcd.io/etcd/client/v3"
	"time"
)

var EtcdClient *etcd.Client

func Init(conf config.Config) {
	var err error
	EtcdClient, err = etcd.New(etcd.Config{
		Endpoints:   conf.Etcd.Hosts,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	//return client
	//kv := etcd.NewKV(client)
}

// GetConfig 获取指定服务的配置
func GetConfig(name string) (string, error) {
	resp, err := EtcdClient.Get(context.TODO(), "/configure/service/"+name)
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) != 0 {
		return string(resp.Kvs[0].Value), nil
	}
	return "", errors.New("为查找到值")
}
