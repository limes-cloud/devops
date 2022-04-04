package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Rsa struct {
		PublicFile  string
		PrivateFile string
		PublicKey   string
		PrivateKey  string
	}
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
}
