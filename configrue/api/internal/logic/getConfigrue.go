package logic

import (
	"context"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigrueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigrueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigrueLogic {
	return &GetConfigrueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigrueLogic) GetConfigrue(req *types.GetConfigureRequest) (resp *types.GetConfigureResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
