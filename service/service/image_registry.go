package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"service/errors"
	"service/model"
	"service/tools"
	"service/tools/code_pack"
	"service/types"
)

func AllImageRegistries(ctx *gin.Context) ([]model.ImageRegistry, error) {
	image := model.ImageRegistry{}
	list, err := image.All(ctx, nil)
	if err != nil {
		return []model.ImageRegistry{}, nil
	}
	for key, item := range list {
		list[key].Password = tools.HideStar(item.Password)
	}
	return list, nil
}

func AllImageRegistryFilter(ctx *gin.Context) ([]model.ImageRegistry, error) {
	image := model.ImageRegistry{}
	return image.All(ctx, nil, func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name")
	})
}

func AddImageRegistry(ctx *gin.Context, in *types.AddImageRegistryRequest) error {
	image := model.ImageRegistry{}
	if copier.Copy(&image, in) != nil {
		return errors.AssignError
	}
	return image.Create(ctx)
}

func UpdateImageRegistry(ctx *gin.Context, in *types.UpdateImageRegistryRequest) error {
	image := model.ImageRegistry{}
	if copier.Copy(&image, in) != nil {
		return errors.AssignError
	}
	return image.UpdateByID(ctx)
}

func DeleteImageRegistry(ctx *gin.Context, in *types.DeleteImageRegistryRequest) error {
	image := model.ImageRegistry{}
	return image.DeleteByID(ctx, in.ID)
}

func ConnectImageRegistry(ctx *gin.Context, in *types.ConnectImageRegistryRequest) error {
	image := model.ImageRegistry{}
	if err := image.OneById(ctx, in.ID); err != nil {
		return err
	}
	pack := code_pack.NewPack()
	pack.RegistryUrl = image.Host
	pack.RegistryUser = image.Username
	pack.RegistryPass = image.Password
	pack.Exec = "/bin/sh"

	if _, err := pack.GetDockerVersion(); err != nil {
		return errors.New("安装主机必须存在docker运行环境")
	}

	if err := pack.Login(); err != nil {
		return errors.New("镜像仓库登陆失败")
	}
	return nil
}
