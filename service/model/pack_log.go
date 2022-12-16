package model

import (
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"service/meta"
)

type PackLog struct {
	gin.CreateModel
	ServiceKeyword    string `json:"service_keyword,omitempty"`
	ServiceName       string `json:"service_name,omitempty"`
	DockerfileName    string `json:"dockerfile_name,omitempty"`
	CodeRegistryName  string `json:"code_registry_name,omitempty"`
	ImageRegistryName string `json:"image_registry_name,omitempty"`
	ImageRegistryID   int64  `json:"image_registry_id"`
	CloneType         string `json:"clone_type,omitempty"`
	CloneValue        string `json:"clone_value,omitempty"`
	CommitID          string `json:"commit_id,omitempty"`
	ImageName         string `json:"image_name,omitempty"`
	UseTime           int64  `json:"use_time"`
	Desc              string `json:"desc,omitempty"`
	IsClear           bool   `json:"is_clear"`
	IsFinish          bool   `json:"is_finish"`
	Status            *bool  `json:"status"`
	Operator          string `json:"operator,omitempty"`
	OperatorId        int64  `json:"operator_id,omitempty"`
}

func (u *PackLog) Table() string {
	return "pack_log"
}

func (u *PackLog) Create(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Create(&u).Error)
}

func (u *PackLog) OneById(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).First(u, "id=?", id).Error)
}

func (u *PackLog) OldHistory(ctx *gin.Context, max int64) []PackLog {
	temp := PackLog{}
	db := database(ctx).Table(u.Table()).Session(&gorm.Session{NewDB: true})
	var list []PackLog
	if db.Where("service_keyword=?", u.ServiceKeyword).
		Where("image_registry_id=?", u.ImageRegistryID).
		Where("clone_value=? and status=true", u.CloneValue).
		Offset(int(max)).Limit(1).Order("id desc").
		First(&temp).Error == nil {

		// 查询所有小于这id的服务镜像
		db.Where("service_keyword=?", u.ServiceKeyword).
			Where("image_registry_id=?", u.ImageRegistryID).
			Where("clone_value=? and status=true and id <= ?", u.CloneValue, temp.ID).Find(&list).Delete(&PackLog{})

	}
	return list
}

func (u *PackLog) Page(ctx *gin.Context, page, count int, m interface{}, fs ...callback) ([]PackLog, int64, error) {
	var list []PackLog
	var total int64

	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	db = execCallback(db, fs...).Order("created_at desc")
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}

	return list, total, transferErr(db.Offset((page - 1) * count).Limit(count).Find(&list).Error)
}

func (u *PackLog) All(ctx *gin.Context, fs ...callback) ([]PackLog, int64, error) {
	var list []PackLog
	var total int64

	db := database(ctx).Table(u.Table())
	db = execCallback(db, fs...).Order("created_at desc")
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}

	return list, total, transferErr(db.Find(&list).Error)
}

func (u *PackLog) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *PackLog) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}
