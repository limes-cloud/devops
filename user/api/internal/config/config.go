package config

import (
	"github.com/zeromicro/go-zero/rest"
	"io"
	"os"
)

type Config struct {
	rest.RestConf
	Mysql struct { //数据库
		DSN             string
		Level           int
		ConnMaxLifetime int
		MaxOpenConn     int
		MaxIdleConn     int
		SlowThreshold   int
	}
	Redis struct {
		Host string
		Pass string
	}
	Rsa struct {
		PrivateKey  string
		PublicKey   string
		PrivateFile string
		PublicFile  string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Whitelist []string
}

//
//func Init(serviceName string) *Config {
//	conf := &Config{}
//	conf.Viper = configx.InitConfig(serviceName, Watch)
//	RsaInit(conf.Viper)
//	if conf.UnmarshalKey("system", &conf.RestConf) != nil {
//		panic("系统信息必须设置")
//	}
//	return conf
//}
//
//func Watch(v *viper.Viper) {
//	log.Info("配置中心发生改变")
//	RsaInit(v)
//}

// RsaInit 初始化rsa
func RsaInit(c *Config) {
	public, err := os.Open(c.Rsa.PublicFile)
	if err != nil {
		panic("初始化rsa-public :" + err.Error())
	}
	private, err := os.Open(c.Rsa.PrivateFile)
	if err != nil {
		panic("初始化rsa-private :" + err.Error())
	}

	defer private.Close()
	defer public.Close()

	pb, _ := io.ReadAll(public)
	rb, _ := io.ReadAll(private)
	c.Rsa.PublicKey = string(pb)
	c.Rsa.PrivateKey = string(rb)
}
