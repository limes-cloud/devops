package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListConfigureLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListConfigureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListConfigureLogic {
	return &ListConfigureLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListConfigureLogic) ListConfigure(req *types.ListConfigureRequest) (resp *types.ListConfigureResponse, err error) {
	conf := models.Configure{}
	resp = new(types.ListConfigureResponse)
	list, _, err := conf.All(req)
	tools.Transform(list, &resp.List)
	return resp, err
}
