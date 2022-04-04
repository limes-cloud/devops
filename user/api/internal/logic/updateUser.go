package logic

import (
	"context"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) error {
	user := models.User{}
	tb := l.svcCtx.Orm.Table(user.Table())
	tb = tb.Where("id = ?", req.ID)
	return tb.Updates(req).Error
}
