package etcx

import (
	"bytes"
	consul "github.com/hashicorp/consul/api"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"time"
)

type Consul struct {
	Client    *consul.Client
	Info      *EtcEnv
	waitIndex uint64
}

func NewConsul(info *EtcEnv) (*Consul, error) {
	client, err := consul.NewClient(&consul.Config{
		Address: info.Host,
		Token:   info.Password,
	})
	if err != nil {
		return nil, err
	}
	return &Consul{
		Client: client,
		Info:   info,
	}, nil
}

// Init 初始化配置
func (e *Consul) Init(v *viper.Viper) {
	b, err := e.GetConfig(nil)
	if err != nil {
		panic(err)
	}
	v.SetConfigType("json")
	if err = v.ReadConfig(bytes.NewBuffer(b)); err != nil {
		panic(err)
	}
	go e.Watch(v)
}

func (e *Consul) Watch(v *viper.Viper) {
	for {
		opts := consul.QueryOptions{
			WaitIndex: e.waitIndex,
		}
		data, err := e.GetConfig(&opts)
		if err != nil {
			log.Info("获取配置信息失败")
			continue
		}
		time.Sleep(1 * time.Second)
		v.ReadConfig(bytes.NewBuffer(data))
		if CallBack != nil {
			CallBack(v)
		}
	}
}

func (e *Consul) GetConfig(q *consul.QueryOptions) ([]byte, error) {
	data, meta, err := e.Client.KV().Get(e.Info.Prefix, q)
	if err != nil {
		return nil, err
	}
	e.waitIndex = meta.LastIndex
	return data.Value, nil
}
