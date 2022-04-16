package mysqlx

import (
	"context"
	"devops/common/drive/redisx"
	"devops/common/tools"
	"encoding/base64"
	"fmt"
	"github.com/zeromicro/go-zero/core/hash"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type Config struct {
	DSN             string
	Level           int
	ConnMaxLifetime int
	MaxOpenConn     int
	MaxIdleConn     int
	SlowThreshold   int
}

var DB *gorm.DB

// NewOrm 新增数据库
func NewOrm(conf Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Duration(conf.SlowThreshold) * time.Second, // 慢 SQL 阈值
				LogLevel:      logger.LogLevel(conf.Level),                     // Log level
				Colorful:      false,                                           // 禁用彩色打印
			},
		),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		panic(err)
	}
	//registerCallBack(db)
	sdb, _ := db.DB()
	sdb.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetime) * time.Second) //设置最大的链接时间
	sdb.SetMaxOpenConns(conf.MaxOpenConn)                                     //最大链接数量
	sdb.SetMaxIdleConns(conf.MaxIdleConn)                                     //最大闲置数量
	DB = db
	return db
}

// 注册全局回调
func registerCallBack(db *gorm.DB) {
	type ModelID struct {
		ID int64 `json:"id"`
	}

	//目前只支持按照ID进行缓存。其他的缓存逻辑需要自定义实现
	db.Callback().Query().Replace("gorm:query", func(db *gorm.DB) {
		//key := ""
		//field := ""
		//if object, ok := db.Statement.Clauses["WHERE"]; ok {
		//	if cond, is := object.Expression.(clause.Where); is {
		//		for _, item := range cond.Exprs {
		//			if expr, isExpr := item.(clause.Expr); isExpr {
		//				if strings.Contains(expr.SQL, "id = ?") {
		//					key = joinWhere(expr.SQL, expr.Vars)
		//				}
		//			}
		//
		//			if expr, isIn := item.(clause.IN); isIn {
		//				if len(expr.Values) == 1 {
		//					key = fmt.Sprintf("id = %v", expr.Values[0])
		//				}
		//			}
		//		}
		//	}
		//}
		//
		//if key == "" {
		//	modelId := ModelID{}
		//	tools.Transform(db.Statement.Model, &modelId)
		//	if modelId.ID != 0 {
		//		key = fmt.Sprintf("id = %v", modelId.ID)
		//	}
		//}
		//if db.Statement.Selects != nil {
		//	field = strings.Join(db.Statement.Selects, ",")
		//}
		//
		//if key == "" { // 没有通过ID进行查询
		//	callbacks.Query(db)
		//	return
		//}
		//key = base64.StdEncoding.EncodeToString(hash.Md5([]byte(db.Statement.Table + key)))
		////进行缓存
		//redisKey := base64.StdEncoding.EncodeToString(hash.Md5([]byte(key + field)))
		//ctx := context.TODO()
		//str, err := redisx.Client.Get(ctx, redisKey).Result()
		//if str != "" && err == nil {
		//	_ = json.Unmarshal([]byte(str), db.Statement.Dest)
		//	return
		//}
		callbacks.Query(db)
		fmt.Print(db)
		//if db.Error == nil {
		//	b, _ := json.Marshal(db.Statement.Dest)
		//	redisx.Client.Set(ctx, redisKey, string(b), 2*time.Hour)
		//	redisx.Client.Set(ctx, key, redisKey, 2*time.Hour)
		//} else {
		//	redisx.Client.Del(ctx, key, redisKey)
		//}
	})

	db.Callback().Update().Replace("gorm:after_update", func(db *gorm.DB) {
		key := ""
		modelId := ModelID{}
		tools.Transform(db.Statement.Model, &modelId)
		if modelId.ID != 0 {
			key = fmt.Sprintf("id = %v", modelId.ID)
			key = base64.StdEncoding.EncodeToString(hash.Md5([]byte(db.Statement.Table + key)))
			ctx := context.TODO()
			if redisKey, _ := redisx.Client.Get(ctx, key).Result(); redisKey != "" {
				redisx.Client.Del(ctx, key, redisKey)
			}
		}
		callbacks.AfterUpdate(db)
	})

	db.Callback().Delete().Replace("gorm:after_delete", func(db *gorm.DB) {
		key := ""
		modelId := ModelID{}
		tools.Transform(db.Statement.Model, &modelId)
		if modelId.ID != 0 {
			key = fmt.Sprintf("id = %v", modelId.ID)
			key = base64.StdEncoding.EncodeToString(hash.Md5([]byte(db.Statement.Table + key)))
			ctx := context.TODO()
			if redisKey, _ := redisx.Client.Get(ctx, key).Result(); redisKey != "" {
				redisx.Client.Del(ctx, key, redisKey)
			}
		}
		callbacks.AfterUpdate(db)
	})
}

func joinWhere(sql string, vars []interface{}) string {
	for _, val := range vars {
		sql = strings.Replace(sql, "?", fmt.Sprintf("%v", val), 1)
	}
	//匹配出 id = ?
	reg := regexp.MustCompile(`id = \d+`)
	return reg.FindString(sql)
}
