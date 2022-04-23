package config

import (
	"devops/common/configx"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/rest"
	"io"
	"os"
)

type Config struct {
	*viper.Viper
	rest.RestConf
}

func Init(serviceName string) *Config {
	conf := &Config{}
	conf.Viper = configx.InitConfig(serviceName, Watch)
	RsaInit(conf.Viper)
	if conf.UnmarshalKey("system", &conf.RestConf) != nil {
		panic("系统信息必须设置")
	}
	return conf
}

func Watch(v *viper.Viper) {
	fmt.Printf("配置中心发生改变")
	RsaInit(v)
}

func RsaInit(v *viper.Viper) {
	public, err := os.Open(v.GetString("rsa.public_file"))
	if err != nil {
		panic("初始化rsa-public :" + err.Error())
	}
	private, err := os.Open(v.GetString("rsa.private_file"))
	if err != nil {
		panic("初始化rsa-private :" + err.Error())
	}

	defer private.Close()
	defer public.Close()

	pb, _ := io.ReadAll(public)
	rb, _ := io.ReadAll(private)
	v.Set("rsa.public_key", string(pb))
	v.Set("rsa.private_key", string(rb))
}
