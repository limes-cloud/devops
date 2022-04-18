package logic

import (
	"context"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigureFieldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateConfigureFieldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigureFieldLogic {
	return &UpdateConfigureFieldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateConfigureFieldLogic) UpdateConfigureField(req *types.UpdateConfigureFieldRequest) error {
	cf := models.ConfigureField{}
	cf.ID = req.ID
	return cf.UpdateByID(l.ctx, req)
}
