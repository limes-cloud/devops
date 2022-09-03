package consts

import "github.com/spf13/viper"

type LoginLimitConfig struct {
	Enable        bool  `json:"enable" mapstructure:"enable"`
	IpLimit       int64 `json:"ip_limit" mapstructure:"ip_limit"`
	PasswordLimit int64 `json:"password_limit" mapstructure:"password_limit"`
}

var WhitelistApi = make(map[string]bool)
var LoginLimit LoginLimitConfig

func InitConfig(v *viper.Viper) {
	_ = v.UnmarshalKey("whitelist", &WhitelistApi)
	_ = v.UnmarshalKey("login_limit", &LoginLimit)
}
