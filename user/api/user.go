package main

import (
	"devops/common/configx"
	"devops/common/meta"
	"devops/user/api/internal/config"
	"devops/user/api/internal/handler"
	"devops/user/api/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

const serviceName = "ums"

func main() {
	flag.Parse()

	var c config.Config
	c.Viper = configx.InitConfig(serviceName)
	restConf := rest.RestConf{}
	c.UnmarshalKey("system", &restConf)
	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(restConf, rest.WithCors())
	server.Use(meta.SetUserIdHandle)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", restConf.Host, restConf.Port)
	server.Start()
}
