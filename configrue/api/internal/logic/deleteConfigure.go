package logic

import (
	"context"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigureLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteConfigureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigureLogic {
	return &DeleteConfigureLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteConfigureLogic) DeleteConfigure(req *types.DeleteConfigureRequest) error {
	conf := models.Configure{}
	conf.ID = req.ID
	return conf.DeleteByID(l.ctx)
}
