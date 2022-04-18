package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"
	"github.com/google/uuid"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigureLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddConfigureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigureLogic {
	return &AddConfigureLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddConfigureLogic) AddConfigure(req *types.AddConfigureRequest) (err error) {
	config := models.Configure{}
	tools.Transform(req, &config)
	config.Version = uuid.NewString()
	if err = config.Create(l.ctx); err == nil {
		NewAddConfigureLog(l.ctx, l.svcCtx).Add(AddConfigureLogRequest{
			ServiceName: req.ServiceName,
			Title:       CreateConfigure,
			Content:     req.Description,
		})
	}
	return err
}
