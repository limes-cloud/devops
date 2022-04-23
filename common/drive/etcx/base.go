package etcx

import (
	"context"
	"errors"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	Timeout = 5 * time.Second
)

type CallFunc func(v *viper.Viper)

var CallBack func(v *viper.Viper)

type EtcEnv struct {
	Env      string `json:"env"`
	Type     string `json:"type"`
	Host     string `json:"host"`
	Prefix   string `json:"prefix"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Etc interface {
	Init(*viper.Viper)
	Watch(*viper.Viper)
}

func Init(info EtcEnv) *viper.Viper {
	var etc Etc
	var err error
	switch info.Type {
	case "etcd":
		if etc, err = NewEtcd(&info); err != nil {
			panic(err)
		}
	case "consul":
		info.Prefix = strings.TrimPrefix(info.Prefix, "/")
		if etc, err = NewConsul(&info); err != nil {
			panic(err)
		}
	default:
		panic("错误的配置环境变量")
	}
	v := viper.New()
	etc.Init(v)
	return v
}

func Update(info *EtcEnv, service, val string) error {

	info.Prefix = info.Prefix + service
	switch info.Type {
	case "etcd":
		if client, err := NewEtcd(info); err != nil {
			return err
		} else {
			_, err = client.Client.KV.Put(context.TODO(), info.Prefix, val)
			return err
		}
	case "consul":
		info.Prefix = strings.TrimPrefix(info.Prefix, "/")
		kv := &api.KVPair{Key: info.Prefix, Value: []byte(val)}
		if client, err := NewConsul(info); err != nil {
		} else {
			_, err = client.Client.KV().Put(kv, nil)
			return err
		}

	default:
		return errors.New("暂不支持的配置中间件")
	}
	return nil
}
