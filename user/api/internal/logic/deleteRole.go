package logic

import (
	"context"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoleLogic) DeleteRole(req *types.DeleteRoleRequest) error {
	role := models.Role{}
	tb := l.svcCtx.Orm.Table(role.Table())
	return tb.Delete(&role, req.ID).Error
}
