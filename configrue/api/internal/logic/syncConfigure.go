package logic

import (
	"context"
	"devops/common/drive/etcx"
	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"
	"devops/configrue/models"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncConfigureLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncConfigureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncConfigureLogic {
	return &SyncConfigureLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncConfigureLogic) SyncConfigure(req *types.SyncConfigureRequest) error {
	parse := NewGetParseConfigureLogic(l.ctx, l.svcCtx)
	resp, err := parse.GetParseConfigure(&types.GetParseConfigureRequest{ID: req.ID, Env: req.Env})
	if err != nil {
		return err
	}

	//获取当前配置的服务
	config := models.Configure{}
	config.ID = req.ID
	config.One(nil, func(db *gorm.DB) *gorm.DB {
		return db.Select("service_id")
	})

	service := models.Service{}
	service.ID = config.ServiceId
	service.OneByID()

	etcEnv, err := GetEtcEnv(req.Env)
	if err != nil {
		return err
	}
	return etcx.Update(etcEnv, service.Keyword, resp.Template)
}

func UpdateEnv(list []etcx.EtcEnv, prefix []string, service, template string) error {
	for index, item := range list {
		if err := etcx.Update(&item, prefix[index]+service, template); err != nil {
			return fmt.Errorf("环境：%v,更新出错：%v", item.Env, err)
		}
	}
	return nil
}

func GetEtcEnv(env string) (*etcx.EtcEnv, error) {
	var info etcx.EtcEnv
	environment := models.Environment{}
	if environment.One(map[string]interface{}{
		"keyword": env,
	}) != nil {
		return nil, errors.New("未找到指定环境：" + env)
	}

	if json.Unmarshal([]byte(environment.Config), &info) != nil {
		return nil, errors.New("解析配置信息失败：" + env)
	}
	info.Type = environment.Drive
	info.Env = environment.Keyword
	info.Prefix = environment.Prefix
	return &info, nil
}
