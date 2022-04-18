package logic

import (
	"context"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEnvironmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEnvironmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEnvironmentLogic {
	return &UpdateEnvironmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEnvironmentLogic) UpdateEnvironment(req *types.UpdateEnvironmentRequest) error {
	env := models.Environment{}
	env.ID = req.ID
	return env.UpdateByID(l.ctx, req)
}
