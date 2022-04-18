package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageServiceFieldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageServiceFieldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageServiceFieldLogic {
	return &PageServiceFieldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageServiceFieldLogic) PageServiceField(req *types.PageServiceFieldRequest) (resp *types.PageServiceFieldResponse, err error) {
	cf := models.ServiceField{}
	resp = new(types.PageServiceFieldResponse)
	list, total, err := cf.Page(req, req.Page, req.Count)
	tools.Transform(list, &resp.List)
	resp.Total = total
	return resp, err
}
