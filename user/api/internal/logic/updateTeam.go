package logic

import (
	"context"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTeamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTeamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTeamLogic {
	return &UpdateTeamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTeamLogic) UpdateTeam(req *types.UpdateTeamRequest) error {
	team := models.Team{}
	tb := l.svcCtx.Orm.Table(team.Table())
	return tb.Updates(req).Error
}
