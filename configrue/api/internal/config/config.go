package config

import "github.com/zeromicro/go-zero/rest"

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
}
