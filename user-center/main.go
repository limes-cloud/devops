package main

import (
	"github.com/limeschool/gin"
	"github.com/spf13/viper"
	"ums/consts"
	"ums/router"
)

func main() {
	engin := router.Init()
	gin.WatchConfig(func(v *viper.Viper) {
		consts.InitConfig(v)
	})
	engin.Run(":8080")
}
