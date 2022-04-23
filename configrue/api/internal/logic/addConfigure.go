package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/api/internal/common"
	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"
	"devops/configrue/models"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"

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

	service := models.Service{}
	service.ID = req.ServiceId
	if service.OneByID() != nil {
		return errors.New("未找到对应的服务信息")
	}

	if _, err = common.ParseTemplate(req.ServiceId, service.Keyword, req.Template); err != nil {
		return err
	}
	config.Version = uuid.NewString()
	//在创建之前进行模板检测
	if err = config.Create(l.ctx); err == nil {
		//删除之前20个版本。
		list, _, _ := config.Page(nil, 20, 1, func(db *gorm.DB) *gorm.DB {
			return db.Order("id desc")
		})
		if len(list) > 0 {
			del := models.Configure{}
			del.Delete(l.ctx, nil, func(db *gorm.DB) *gorm.DB {
				return db.Where("id < ?", list[0].ID)
			})
		}
		//日志入库
		NewAddConfigureLog(l.ctx, l.svcCtx).Add(AddConfigureLogRequest{
			ServiceName: service.Name,
			Title:       CreateConfigure,
			Content:     req.Description,
		})

	}

	return err
}
