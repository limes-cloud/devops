package logic

import (
	"context"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddServiceLogic {
	return &AddServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddServiceLogic) AddService(req *types.AddServiceRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
