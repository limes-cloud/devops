package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"service/errors"
	"service/model"
	"service/types"
)

func AllRelease(ctx *gin.Context, in *types.AllReleaseRequest) ([]model.Release, error) {
	release := model.Release{}
	return release.All(ctx, in)
}

func AddRelease(ctx *gin.Context, in *types.AddReleaseRequest) error {
	release := model.Release{}
	if copier.Copy(&release, in) != nil {
		return errors.AssignError
	}
	return release.Create(ctx)
}

func UpdateRelease(ctx *gin.Context, in *types.UpdateReleaseRequest) error {
	release := model.Release{}
	if copier.Copy(&release, in) != nil {
		return errors.AssignError
	}
	return release.UpdateByID(ctx)
}

func DeleteRelease(ctx *gin.Context, in *types.DeleteReleaseRequest) error {
	release := model.Release{}
	return release.DeleteByID(ctx, in.ID)
}

func AllReleaseImages(ctx *gin.Context, in *types.AllReleaseImagesRequest) ([]model.PackLog, error) {
	log := model.PackLog{}
	logs, _, err := log.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Select("service_name,image_name,created_at").
			Where("service_keyword = ?", in.ServiceKeyword).
			Where("is_clear = false"). // 镜像未清理
			Where("status = true")     // 发布成功的
	})
	return logs, err
}
