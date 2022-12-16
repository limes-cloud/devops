package model

import (
	"github.com/limeschool/gin"
	"service/meta"
)

type ReleaseLog struct {
	gin.CreateModel
	ServiceKeyword    string `json:"service_keyword"`
	ServiceName       string `json:"service_name"`
	ImageName         string `json:"image_name"`
	ImageRegistryName string `json:"image_registry_name"`
	UseTime           int64  `json:"use_time"`
	Desc              string `json:"desc"`
	IsFinish          bool   `json:"is_finish"`
	EnvName           string `json:"env_name"`
	EnvKeyword        string `json:"env_keyword"`
	Status            string `json:"status"`
	Operator          string `json:"operator"`
	OperatorId        int64  `json:"operator_id"`
}

func (u *ReleaseLog) Table() string {
	return "release_log"
}

func (u *ReleaseLog) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *ReleaseLog) OneById(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "id=?", id).Error)
}

func (u *ReleaseLog) Page(ctx *gin.Context, page, count int, m interface{}, fs ...callback) ([]ReleaseLog, int64, error) {
	var list []ReleaseLog
	var total int64

	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	db = execCallback(db, fs...).Order("created_at desc")
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}

	return list, total, transferErr(db.Offset((page - 1) * count).Limit(count).Find(&list).Error)
}

func (u *ReleaseLog) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *ReleaseLog) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}
