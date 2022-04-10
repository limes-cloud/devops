package logic

import (
	"context"
	"devops/common/typex"
	"devops/user/models"

	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuLogic) UpdateMenu(req *types.UpdateMenuRequest) error {
	menu := models.Menu{}
	menu.ID = req.ID
	if menu.OneByID() == nil {
		if menu.Permission == typex.BaseApiKey {
			l.svcCtx.Rbac.RemovePolicy(typex.PublicRoleKey, menu.Path, menu.Method)
		}
	}

	if req.Permission == typex.BaseApiKey {
		l.svcCtx.Rbac.AddPolicy(typex.PublicRoleKey, menu.Path, menu.Method)
	}

	tb := l.svcCtx.Orm.Table(menu.Table())
	return tb.Updates(req).Error
}
