package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageConfigureFieldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageConfigureFieldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageConfigureFieldLogic {
	return &PageConfigureFieldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageConfigureFieldLogic) PageConfigureField(req *types.PageConfigureFieldRequest) (resp *types.PageConfigureFieldResponse, err error) {
	cf := models.ConfigureField{}
	resp = new(types.PageConfigureFieldResponse)
	list, total, err := cf.Page(req, req.Page, req.Count)
	tools.Transform(list, &resp.List)
	resp.Total = total
	return resp, err
}
