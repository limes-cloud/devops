package logic

import (
	"context"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigureFieldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteConfigureFieldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigureFieldLogic {
	return &DeleteConfigureFieldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteConfigureFieldLogic) DeleteConfigureField(req *types.DeleteConfigureFieldRequest) error {
	cd := models.ConfigureField{}
	cd.ID = req.ID
	return cd.DeleteByID(l.ctx)
}
