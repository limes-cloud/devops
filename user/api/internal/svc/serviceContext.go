package svc

import (
	"devops/common/drive/mysqlx"
	"devops/common/drive/redisx"
	"devops/common/tools"
	"devops/user/api/internal/config"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Orm    *gorm.DB
	Redis  *redis.Client
	Rbac   *casbin.Enforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := NewOrm(c)
	return &ServiceContext{
		Config: c,
		Orm:    db,
		Redis:  NewRedis(c),
		Rbac:   NewRbac(db),
	}
}

func NewOrm(c config.Config) *gorm.DB {
	conf := mysqlx.Config{}
	tools.Transform(c.Mysql, &conf)
	return mysqlx.NewOrm(conf)
}

// NewRedis 新增Redis
func NewRedis(c config.Config) *redis.Client {
	conf := redisx.Config{}
	tools.Transform(c.Redis, &conf)
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
