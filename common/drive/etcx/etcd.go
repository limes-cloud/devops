package etcx

import (
	"context"
	"encoding/json"
	etcd "go.etcd.io/etcd/client/v3"
	"strings"
)

type Etcd struct {
	Client *etcd.Client
}

func NewEtcd(info *EtcEnv) (*Etcd, error) {
	client, err := etcd.New(etcd.Config{
		Endpoints:   strings.Split(info.Host, ","),
		DialTimeout: Timeout,
	})
	if err != nil {
		panic(err)
	}
	return &Etcd{
		Client: client,
	}, nil
}

// Init 初始化配置
func (e *Etcd) Init(conf interface{}) {
	str, err := e.GetConfig()
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(str, conf); err != nil {
		panic(err)
	}

}

func (e *Etcd) GetConfig() ([]byte, error) {
	resp, err := e.Client.KV.Get(context.TODO(), "/configure/service/configure")
	if len(resp.Kvs) != 0 {
		return resp.Kvs[0].Value, nil
	}
	return nil, err
}
