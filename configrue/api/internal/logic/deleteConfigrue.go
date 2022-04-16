package logic

import (
	"context"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigrueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteConfigrueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigrueLogic {
	return &DeleteConfigrueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteConfigrueLogic) DeleteConfigrue(req *types.DeleteConfigureRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
