package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigureLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigureLogic {
	return &GetConfigureLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigureLogic) GetConfigure(req *types.GetConfigureRequest) (resp *types.Configure, err error) {
	conf := models.Configure{}
	conf.ID = req.ID
	resp = new(types.Configure)
	err = conf.OneByID()
	tools.Transform(conf, &resp)
	return resp, err
}
