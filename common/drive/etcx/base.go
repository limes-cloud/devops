package etcx

import (
	"os"
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
	Init(interface{})
}

func Init(conf interface{}) {
	var etc Etc
	var err error
	info := NewEtcEnv()
	switch info.Type {
	case "etcd":
		if etc, err = NewEtcd(info); err != nil {
			panic(err)
		}
	default:
		panic("错误的配置环境变量")
	}
	etc.Init(&conf)
}

func NewEtcEnv() *EtcEnv {
	return &EtcEnv{
		Env:      os.Getenv("env"),
		Type:     os.Getenv("etc_type"),
		Host:     os.Getenv("etc_host"),
		Username: os.Getenv("username"),
		Password: os.Getenv("Password"),
	}
}
