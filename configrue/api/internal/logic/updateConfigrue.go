package logic

import (
	"context"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigrueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateConfigrueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigrueLogic {
	return &UpdateConfigrueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateConfigrueLogic) UpdateConfigrue(req *types.UpdateConfigureRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
