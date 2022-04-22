package etcx

import (
	"context"
	"github.com/spf13/viper"
	"time"
)

const (
	Timeout     = 5 * time.Second
	ServiceName = "configure"
)

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
	}
	return nil
}
