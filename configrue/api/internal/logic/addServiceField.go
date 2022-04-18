package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddServiceFieldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddServiceFieldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddServiceFieldLogic {
	return &AddServiceFieldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddServiceFieldLogic) AddServiceField(req *types.AddServiceFieldRequest) error {
	sf := models.ServiceField{}
	tools.Transform(req, &sf)
	return sf.Create(l.ctx)
}
