package logic

import (
	"context"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteServiceLogic {
	return &DeleteServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteServiceLogic) DeleteService(req *types.DeleteServiceRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
