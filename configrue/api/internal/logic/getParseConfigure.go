package logic

import (
	"context"
	"devops/configrue/api/internal/common"
	"devops/configrue/models"
	"errors"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetParseConfigureLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetParseConfigureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetParseConfigureLogic {
	return &GetParseConfigureLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetParseConfigure 获取解析之后的数据
func (l *GetParseConfigureLogic) GetParseConfigure(req *types.GetParseConfigureRequest) (resp *types.Configure, err error) {
	config := models.Configure{}
	resp = new(types.Configure)
	config.ID = req.ID
	if config.OneByID() != nil {
		return nil, errors.New("暂无此配置")
	}
	resp.Template, err = common.ParseTemplate(config.ServiceId, req.Env, config.Template)
	return
}
