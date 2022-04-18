package logic

import (
	"context"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateServiceFieldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateServiceFieldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateServiceFieldLogic {
	return &UpdateServiceFieldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateServiceFieldLogic) UpdateServiceField(req *types.UpdateServiceFieldRequest) error {
	sf := models.ServiceField{}
	sf.ID = req.ID
	return sf.UpdateByID(l.ctx, req)
}
