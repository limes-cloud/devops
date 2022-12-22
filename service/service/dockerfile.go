package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"service/errors"
	"service/model"
	"service/types"
)

func PageDockerfile(ctx *gin.Context, in *types.PageDockerfileRequest) ([]model.Dockerfile, int64, error) {
	image := model.Dockerfile{}
	return image.Page(ctx, in.Page, in.Count, in)
}

func AllDockerfileFilter(ctx *gin.Context) ([]model.Dockerfile, error) {
	image := model.Dockerfile{}
	return image.All(ctx, nil, func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name")
	})
}

func AddDockerfile(ctx *gin.Context, in *types.AddDockerfileRequest) error {
	image := model.Dockerfile{}
	if copier.Copy(&image, in) != nil {
		return errors.AssignError
	}
	return image.Create(ctx)
}

func UpdateDockerfile(ctx *gin.Context, in *types.UpdateDockerfileRequest) error {
	image := model.Dockerfile{}
	if copier.Copy(&image, in) != nil {
		return errors.AssignError
	}
	return image.UpdateByID(ctx)
}

func DeleteDockerfile(ctx *gin.Context, in *types.DeleteDockerfileRequest) error {
	image := model.Dockerfile{}
	return image.DeleteByID(ctx, in.ID)
}
