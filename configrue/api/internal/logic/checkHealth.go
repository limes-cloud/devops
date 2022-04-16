package logic

import (
	"context"

	"devops/configrue/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckHealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckHealthLogic {
	return &CheckHealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckHealthLogic) CheckHealth() error {
	// todo: add your logic here and delete this line

	return nil
}
