package model

import (
	"github.com/limeschool/gin"
	"service/meta"
)

type Service struct {
	gin.BaseModel
	Keyword           string  `json:"keyword"`
	Name              string  `json:"name"`
	IsPrivate         *bool   `json:"is_private"`
	Description       *string `json:"description"`
	Operator          string  `json:"operator"`
	OperatorId        int64   `json:"operator_id"`
	TeamID            *int64  `json:"team_id"`
	EnvIds            []int64 `json:"env_ids" gorm:"-"`
	ReleaseID         int64   `json:"release_id" binding:"required"`
	DockerfileID      int64   `json:"dockerfile_id"`
	DockerfileName    string  `json:"dockerfile_name" gorm:"->"`
	CodeRegistryID    int64   `json:"code_registry_id"`
	CodeRegistryName  string  `json:"code_registry_name" gorm:"->"`
	ImageRegistryID   int64   `json:"image_registry_id"`
	ImageRegistryName string  `json:"image_registry_name"  gorm:"->"`
	RunPort           int64   `json:"run_port"`
	ListenPort        int64   `json:"listen_port"`
	Owner             string  `json:"owner"`
	Repo              string  `json:"repo"`
	Replicas          int64   `json:"replicas" binding:"required"`
	ProbeType         string  `json:"probe_type" binding:"required"`
	ProbeValue        string  `json:"probe_value" binding:"required"`
	ProbeInitDelay    int64   `json:"probe_init_delay" binding:"required"`
}

func (u *Service) Table() string {
	return "service"
}

func (u *Service) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *Service) OneById(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "id=?", id).Error)
}

func (u *Service) OneByKeyword(ctx *gin.Context, key string) error {
	return database(ctx).Table(u.Table()).First(u, "keyword=?", key).Error
}

func (u *Service) Page(ctx *gin.Context, page, count int, m interface{}, fs ...callback) ([]Service, int64, error) {
	var list []Service
	var total int64

	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	db = execCallback(db, fs...)
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}

	return list, total, transferErr(db.Offset((page - 1) * count).Limit(count).Find(&list).Error)
}

func (u *Service) Count(ctx *gin.Context, m interface{}, fs ...callback) int64 {
	var total int64
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	db = execCallback(db, fs...)
	db.Count(&total)
	return total
}

func (u *Service) Filter(ctx *gin.Context, page, count int, m interface{}, fs ...callback) ([]Service, int64, error) {
	var list []Service
	var total int64
	db := database(ctx).Table(u.Table()).Select("id,keyword,name")
	db = gin.GormWhere(db, u.Table(), m)
	db = execCallback(db, fs...)

	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}
	return list, total, transferErr(db.Offset((page - 1) * count).Limit(count).Find(&list).Error)
}

func (u *Service) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *Service) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}
