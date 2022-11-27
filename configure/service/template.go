package service

import (
	"configure/errors"
	"configure/model"
	"configure/tools"
	"configure/types"
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
)

func AllTemplate(ctx *gin.Context, in *types.AllTemplateRequest) ([]model.Template, error) {
	m := model.Template{}
	return m.All(ctx, in)
}

func GetTemplate(ctx *gin.Context, in *types.GetTemplateRequest) (*model.Template, error) {
	m := model.Template{}
	if in.ID != 0 {
		return &m, m.OneById(ctx, in.ID)
	} else {
		return &m, m.One(ctx, "service_keyword = ? and is_use = true", in.Keyword)
	}
}

func AddTemplate(ctx *gin.Context, in *types.AddTemplateRequest) error {
	m := model.Template{}
	if copier.Copy(&m, in) != nil {
		return errors.AssignError
	}
	if err := CheckTemplate(ctx, in.ServiceKeyword, in.Content); err != nil {
		return err
	}
	m.IsUse = true
	m.Version = tools.GenVersion()
	return m.Create(ctx)
}

func UpdateTemplate(ctx *gin.Context, in *types.UpdateTemplateRequest) error {
	m := model.Template{}
	m.ID = in.ID
	return m.UpdateVersionByID(ctx)
}
