package svc

import (
	"devops/common/drive/mysqlx"
	"devops/common/drive/redisx"
	"devops/common/tools"
	"devops/configrue/api/internal/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Orm    *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := NewOrm(c)
	return &ServiceContext{
		Config: c,
		Orm:    db,
		Redis:  NewRedis(c),
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
	tools.Transform(c.Mysql, &conf)
	return redisx.NewClient(conf)
}
