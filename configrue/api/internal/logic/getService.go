package logic

import (
	"context"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceLogic {
	return &GetServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServiceLogic) GetService(req *types.GetServiceRequest) (resp *types.GetServiceResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
