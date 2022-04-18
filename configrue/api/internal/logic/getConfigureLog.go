package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigureLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigureLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigureLogLogic {
	return &GetConfigureLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigureLogLogic) GetConfigureLog(req *types.GetConfigureLogRequest) (resp *types.GetConfigureLogResponse, err error) {
	conf := models.ConfigureLog{}
	resp = new(types.GetConfigureLogResponse)
	list, total, err := conf.Page(req, req.Page, req.Count)
	tools.Transform(list, &resp.List)
	resp.Total = total
	return resp, err
}
