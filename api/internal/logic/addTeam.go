package logic

import (
	"context"

	"configure/api/internal/svc"
	"configure/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTeamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTeamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTeamLogic {
	return &AddTeamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTeamLogic) AddTeam(req *types.AddTeamRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
