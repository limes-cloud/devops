package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigureFieldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddConfigureFieldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigureFieldLogic {
	return &AddConfigureFieldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddConfigureFieldLogic) AddConfigureField(req *types.AddConfigureFieldRequest) error {
	cf := models.ConfigureField{}
	tools.Transform(req, &cf)
	return cf.Create(l.ctx)
}
