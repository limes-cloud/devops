package logic

import (
	"context"
	"devops/common/tools"
	"devops/common/typex"
	"devops/user/models"

	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMenuLogic) AddMenu(req *types.AddMenuRequest) error {
	menu := models.Menu{}
	tools.Transform(req, &menu)
	if menu.Permission == typex.BaseApiKey {
		l.svcCtx.Rbac.AddPolicy(typex.PublicRoleKey, menu.Path, menu.Method)
	}
	return menu.Create(l.ctx)
}
