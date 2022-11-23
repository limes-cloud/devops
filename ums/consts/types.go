package consts

import "github.com/spf13/viper"

type loginLimitConfig struct {
	Enable             bool  `json:"enable" mapstructure:"enable"`
	IpLimit            int64 `json:"ip_limit" mapstructure:"ip_limit"`
	PasswordErrorLimit int64 `json:"password_error_limit" mapstructure:"password_error_limit"`
}

type jwtConfig struct {
	Expire    int64  `json:"expire" mapstructure:"expire"`
	MaxExpire int64  `json:"max_expire" mapstructure:"max_expire"`
	Secret    string `json:"secret" mapstructure:"secret"`
}

var WhitelistApi = make(map[string]bool)
var LoginLimit loginLimitConfig
var Jwt jwtConfig

func InitConfig(v *viper.Viper) {
	_ = v.UnmarshalKey("whitelist", &WhitelistApi)
	_ = v.UnmarshalKey("login_limit", &LoginLimit)
	_ = v.UnmarshalKey("jwt", &Jwt)
}
