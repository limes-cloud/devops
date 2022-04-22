package logic

import (
	"context"
	"devops/common/tools"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEnvConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEnvConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEnvConfigLogic {
	return &GetEnvConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEnvConfigLogic) GetEnvConfig(req *types.GetEnvConfigRequest) (resp *types.GetEnvConfigResponse, err error) {
	info, err := GetEtcEnv(req.Env)
	if err != nil {
		return nil, err
	}
	resp = new(types.GetEnvConfigResponse)
	tools.Transform(info, &resp)
	return
}
