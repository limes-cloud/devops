package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddEnvironmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddEnvironmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddEnvironmentLogic {
	return &AddEnvironmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddEnvironmentLogic) AddEnvironment(req *types.AddEnvironmentRequest) error {
	env := models.Environment{}
	tools.Transform(req, &env)
	return env.Create(l.ctx)
}
