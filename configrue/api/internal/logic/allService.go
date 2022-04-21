package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllServiceLogic {
	return &AllServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllServiceLogic) AllService() (resp *types.AllServiceResponse, err error) {
	service := models.Service{}
	resp = new(types.AllServiceResponse)
	list, _, err := service.All(nil)
	tools.Transform(list, &resp.List)
	return resp, err
}
