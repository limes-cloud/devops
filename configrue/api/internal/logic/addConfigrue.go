package logic

import (
	"context"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigrueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddConfigrueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigrueLogic {
	return &AddConfigrueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddConfigrueLogic) AddConfigrue(req *types.AddConfigureRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
