package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigureLog struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

const (
	SyncVersion     = "同步版本"
	CreateConfigure = "创建服务配置"
	UpdateConfigure = "同步配置"
)

func NewAddConfigureLog(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigureLog {
	return &AddConfigureLog{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type AddConfigureLogRequest struct {
	ServiceName string `json:"service_name"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}

func (l *AddConfigureLog) Add(req AddConfigureLogRequest) (err error) {
	log := models.ConfigureLog{}
	tools.Transform(req, &log)
	return log.Create(l.ctx)
}
