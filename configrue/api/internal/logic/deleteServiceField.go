package logic

import (
	"context"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteServiceFieldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteServiceFieldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteServiceFieldLogic {
	return &DeleteServiceFieldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteServiceFieldLogic) DeleteServiceField(req *types.DeleteServiceFieldRequest) error {
	sf := models.ServiceField{}
	sf.ID = req.ID
	return sf.DeleteByID(l.ctx)
}
