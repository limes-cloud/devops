package etcx

import (
	"bytes"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/api/v3/mvccpb"
	etcd "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

type Etcd struct {
	Client *etcd.Client
	Info   *EtcEnv
}

func NewEtcd(info *EtcEnv) (*Etcd, error) {
	client, err := etcd.New(etcd.Config{
		Endpoints:   strings.Split(info.Host, ","),
		DialTimeout: Timeout,
	})
	if err != nil {
		return nil, err
	}
	return &Etcd{
		Client: client,
		Info:   info,
	}, nil
}

// Init 初始化配置
func (e *Etcd) Init(v *viper.Viper) {
	b, err := e.GetConfig(e.Info.Prefix)
	if err != nil {
		panic(err)
	}
	v.SetConfigType("json")
	if err = v.ReadConfig(bytes.NewBuffer(b)); err != nil {
		panic(err)
	}
	go e.Watch(v)
}

func (e *Etcd) Watch(v *viper.Viper) {
	for {
		ctx, cancelFunc := context.WithCancel(context.TODO())
		time.AfterFunc(5*time.Second, func() {
			cancelFunc()
		})

		watchRespChan := e.Client.Watch(ctx, e.Info.Prefix)
		for watchResp := range watchRespChan {
			for _, event := range watchResp.Events {
				switch event.Type {
				case mvccpb.PUT:
					fmt.Printf("配置被修改：%s", event.Kv.Value)
					v.ReadConfig(bytes.NewBuffer(event.Kv.Value))
					if CallBack != nil {
						CallBack(v)
					}
				}
			}
		}
	}
}

func (e *Etcd) GetConfig(prefix string) ([]byte, error) {
	resp, err := e.Client.KV.Get(context.TODO(), prefix)
	if len(resp.Kvs) != 0 {
		return resp.Kvs[0].Value, nil
	}
	return nil, err
}
