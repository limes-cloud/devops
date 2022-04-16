package logic

import (
	"context"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigrueLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigrueLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigrueLogLogic {
	return &GetConfigrueLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigrueLogLogic) GetConfigrueLog(req *types.GetConfigureLogRequest) (resp *types.GetConfigureLogResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
