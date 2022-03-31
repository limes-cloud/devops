package svc

import {{.imports}}

type ServiceContext struct {
	Config config.Config
	Orm    *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Orm:    NewOrm(c),
		Redis:  NewRedis(c),
	}
}

// NewOrm 新增数据库
func NewOrm(c config.Config) *gorm.DB {
	conf := c.Mysql
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Duration(conf.SlowThreshold) * time.Second, // 慢 SQL 阈值
				LogLevel:      logger.LogLevel(conf.Level),                     // Log level
				Colorful:      false,                                           // 禁用彩色打印
			},
		),
	})
	if err != nil {
		panic(err)
	}
	sdb, _ := db.DB()
	sdb.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetime) * time.Second) //设置最大的链接时间
	sdb.SetMaxOpenConns(conf.MaxOpenConn)                                     //最大链接数量
	sdb.SetMaxIdleConns(conf.MaxIdleConn)                                     //最大闲置数量
	return db
}

// NewRedis 新增Redis
func NewRedis(c config.Config) *redis.Client {
	conf := c.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: conf.Pass,
	})

	if err := client.Ping(context.TODO()).Err(); err != nil {
		panic(err)
	}

	return client
}
