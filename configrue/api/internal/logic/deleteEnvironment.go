package logic

import (
	"context"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEnvironmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEnvironmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEnvironmentLogic {
	return &DeleteEnvironmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEnvironmentLogic) DeleteEnvironment(req *types.DeleteEnvironmentRequest) error {
	env := models.Environment{}
	env.ID = req.ID
	return env.DeleteByID(l.ctx)
}
