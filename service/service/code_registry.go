package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"service/errors"
	"service/model"
	"service/tools"
	"service/tools/code_registry"
	codeModel "service/tools/code_registry/model"
	"service/types"
)

func AllCodeRegistries(ctx *gin.Context) ([]model.CodeRegistry, error) {
	code := model.CodeRegistry{}
	list, err := code.All(ctx, nil)
	if err != nil {
		return []model.CodeRegistry{}, nil
	}
	for key, item := range list {
		list[key].Token = tools.HideStar(item.Token)
	}
	return list, nil
}

func AllCodeRegistryFilter(ctx *gin.Context) ([]model.CodeRegistry, error) {
	code := model.CodeRegistry{}
	return code.All(ctx, nil, func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name")
	})
}

func AddCodeRegistry(ctx *gin.Context, in *types.AddCodeRegistryRequest) error {
	code := model.CodeRegistry{}
	if copier.Copy(&code, in) != nil {
		return errors.AssignError
	}
	return code.Create(ctx)
}

func UpdateCodeRegistry(ctx *gin.Context, in *types.UpdateCodeRegistryRequest) error {
	code := model.CodeRegistry{}
	if copier.Copy(&code, in) != nil {
		return errors.AssignError
	}
	return code.UpdateByID(ctx)
}

func DeleteCodeRegistry(ctx *gin.Context, in *types.DeleteCodeRegistryRequest) error {
	code := model.CodeRegistry{}
	return code.DeleteByID(ctx, in.ID)
}

func ConnectCodeRegistry(ctx *gin.Context, in *types.ConnectCodeRegistryRequest) error {
	code := model.CodeRegistry{}
	if err := code.OneById(ctx, in.ID); err != nil {
		return err
	}
	_, err := code_registry.NewCodeRegistry(code.Type, code.Host, code.Token)
	if err != nil {
		return err
	}
	return nil
}

func GetCodeRegistryProject(ctx *gin.Context, in *types.GetCodeRegistryProjectRequest) (*codeModel.Project, error) {
	code := model.CodeRegistry{}
	if err := code.OneById(ctx, in.ID); err != nil {
		return nil, err
	}
	client, err := code_registry.NewCodeRegistry(code.Type, code.Host, code.Token)
	if err != nil {
		return nil, err
	}
	return client.GetRepo(in.Owner, in.Repo)
}

func AllCodeRegistryBranches(ctx *gin.Context, in *types.AllCodeRegistryBranchesRequest) ([]codeModel.Branch, error) {
	service := model.Service{}
	if err := service.OneByKeyword(ctx, in.ServiceKeyword); err != nil {
		return nil, err
	}

	code := model.CodeRegistry{}
	if err := code.OneById(ctx, service.CodeRegistryID); err != nil {
		return nil, err
	}

	client, err := code_registry.NewCodeRegistry(code.Type, code.Host, code.Token)
	if err != nil {
		return nil, err
	}

	project, err := client.GetRepo(service.Owner, service.Repo)
	if err != nil {
		return nil, err
	}

	return client.GetRepoBranches(project)
}

func AllCodeRegistryTags(ctx *gin.Context, in *types.AllCodeRegistryTagsRequest) ([]codeModel.Tag, error) {
	service := model.Service{}
	if err := service.OneByKeyword(ctx, in.ServiceKeyword); err != nil {
		return nil, err
	}

	code := model.CodeRegistry{}
	if err := code.OneById(ctx, service.CodeRegistryID); err != nil {
		return nil, err
	}

	client, err := code_registry.NewCodeRegistry(code.Type, code.Host, code.Token)
	if err != nil {
		return nil, err
	}

	project, err := client.GetRepo(service.Owner, service.Repo)
	if err != nil {
		return nil, err
	}

	return client.GetRepoTags(project)
}
