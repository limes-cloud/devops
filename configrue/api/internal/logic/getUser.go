package logic

import (
	"configure/api/internal/svc"
	"configure/api/internal/types"
	"configure/common/meta"
	"configure/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser() (resp *types.GetUserResponse, err error) {
	userId := meta.UserId(l.ctx)
	user := models.User{}
	resp = new(types.GetUserResponse)
	tb := l.svcCtx.Orm.Table(user.Table()).Select("user.*,team.name team_name,role.name role_name")
	return resp, tb.First(resp, userId).Error
}
