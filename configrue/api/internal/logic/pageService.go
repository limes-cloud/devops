package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"
	"gorm.io/gorm"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageServiceLogic {
	return &PageServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageServiceLogic) PageService(req *types.PageServiceRequest) (resp *types.PageServiceResponse, err error) {
	service := models.Service{}
	resp = new(types.PageServiceResponse)
	list, count, err := service.Page(nil, req.Page, req.Count, func(db *gorm.DB) *gorm.DB {
		if req.Keyword != "" {
			return db.Where("keyword like ?", "%"+req.Keyword+"%")
		}
		return db
	})
	tools.Transform(list, &resp.List)
	resp.Total = count
	return resp, err
}
