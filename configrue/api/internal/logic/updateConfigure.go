package logic

import (
	"context"
	"devops/configrue/models"
	"gorm.io/gorm"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigureLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateConfigureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigureLogic {
	return &UpdateConfigureLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateConfigureLogic) UpdateConfigure(req *types.UpdateConfigureRequest) error {
	conf := models.Configure{}
	conf.ID = req.ID
	conf.Update(l.ctx, nil, map[string]interface{}{"is_use": false}, func(db *gorm.DB) *gorm.DB {
		return db.Where("id != ?", req.ID)
	})
	_ = conf.OneByID()

	if conf.UpdateByID(l.ctx, map[string]interface{}{"is_use": true}) == nil {
		NewAddConfigureLog(l.ctx, l.svcCtx).Add(AddConfigureLogRequest{
			ServiceName: req.ServiceName,
			Title:       UpdateConfigure,
			Content:     "同步配置版本:" + conf.Version,
		})
	}
	//conf.up
	//这里应该获取字段。然后把数据发送到mQ中。
	//由mq去进行发送同步到ZK
	return conf.UpdateByID(l.ctx, map[string]interface{}{"is_use": true})

}
