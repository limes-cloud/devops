package config

import "github.com/spf13/viper"

type Config struct {
	*viper.Viper
	//rest.RestConf
	//Auth struct { //token密钥和过期时间
	//	AccessSecret string
	//	AccessExpire int64
	//}
	//
	//WhitePath []string //路由白名单
	//
	//Rsa struct {
	//	PublicFile  string
	//	PrivateFile string
	//	PublicKey   string
	//	PrivateKey  string
	//}
	//Mysql struct { //数据库
	//	DSN             string
	//	Level           int
	//	ConnMaxLifetime int
	//	MaxOpenConn     int
	//	MaxIdleConn     int
	//	SlowThreshold   int
	//}
	//Redis struct {
	//	Host string
	//	Pass string
	//}
}
