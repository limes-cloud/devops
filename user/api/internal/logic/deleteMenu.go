package logic

import (
	"context"
	"devops/common/typex"
	"devops/user/models"

	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuLogic) DeleteMenu(req *types.DeleteMenuRequest) error {
	menu := models.Menu{}
	menu.ID = req.ID
	//获取之前的数据是否是开放菜单
	if menu.OneByID() == nil {
		if menu.Permission == typex.BaseApiKey {
			l.svcCtx.Rbac.RemovePolicy(typex.PublicRoleKey, menu.Path, menu.Method)
		}
	}
	return menu.DeleteByID(l.ctx)
}
