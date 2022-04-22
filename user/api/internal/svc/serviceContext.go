package svc

import (
	"devops/common/drive/mysqlx"
	"devops/common/drive/redisx"
	"devops/user/api/internal/config"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Orm    *gorm.DB
	Redis  *redis.Client
	Rbac   *casbin.Enforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := NewOrm(c.Viper)
	return &ServiceContext{
		Config: c,
		Orm:    db,
		Redis:  NewRedis(c.Viper),
		Rbac:   NewRbac(db),
	}
}

func NewOrm(c *viper.Viper) *gorm.DB {
	conf := mysqlx.Config{}
	if err := c.UnmarshalKey("mysql", &conf); err != nil {
		panic("获取数据库配置失败" + err.Error())
	}
	return mysqlx.NewOrm(conf)
}

// NewRedis 新增Redis
func NewRedis(c *viper.Viper) *redis.Client {
	conf := redisx.Config{}
	if err := c.UnmarshalKey("redis", &conf); err != nil {
		panic("获取redis配置失败" + err.Error())
	}
	return redisx.NewClient(conf)
}

// NewRbac 鉴权
func NewRbac(db *gorm.DB) *casbin.Enforcer {
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")

	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}
	e, _ := casbin.NewEnforcer(m, a)
	if err = e.LoadPolicy(); err != nil {
		panic(err)
	}
	return e
}
