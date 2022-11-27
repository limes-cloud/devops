package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"gorm.io/gorm"
	"service/errors"
	"service/meta"
	"service/model"
	"service/tools"
	"service/types"
)

func AllServiceEnv(ctx *gin.Context, in *types.AllServiceEnvRequest) ([]model.Environment, error) {
	env := model.Environment{}
	return env.AllFilter(ctx, "id in (select env_id from service_env where srv_id = ?)", in.SrvId)
}

func PageService(ctx *gin.Context, in *types.PageServiceRequest) ([]model.Service, int64, error) {
	// 获取当前用户的部门id
	userTeamIds, err := UserTeamIds(ctx)
	if err != nil {
		return nil, 0, err
	}

	// 判断是否具有部门的操作权限
	if in.IsPrivate != nil && !*in.IsPrivate && in.TeamID != nil {
		if !tools.InList(userTeamIds, *in.TeamID) {
			return nil, 0, errors.New("暂无此部门的操作权限")
		}
	}

	// 查询用户权限内的服务
	srv := model.Service{}
	list, total, err := srv.Page(ctx, in.Page, in.Count, in, func(db *gorm.DB) *gorm.DB {
		if in.IsPrivate != nil && !*in.IsPrivate {
			if in.TeamID != nil {
				db = db.Where("team_id=?", in.TeamID)
			} else {
				db = db.Where("team_id in ?", userTeamIds)
			}
		}
		return db
	})

	if err != nil {
		return nil, total, err
	}

	for key, item := range list {
		envSrv := model.ServiceEnv{}
		envSrvList, _ := envSrv.All(ctx, "srv_id = ?", item.BaseModel.ID)
		var ids []int64
		for _, ite := range envSrvList {
			ids = append(ids, ite.EnvId)
		}
		list[key].EnvIds = ids
	}

	return list, total, nil
}

func AddService(ctx *gin.Context, in *types.AddServiceRequest) error {
	srv := model.Service{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	if err := srv.Create(ctx); err != nil {
		return err
	}
	return AddServiceEnv(ctx, srv.BaseModel.ID, in.EnvIds)
}

func AddServiceEnv(ctx *gin.Context, srvId int64, envIds []int64) error {
	srv := model.ServiceEnv{}
	_ = srv.Delete(ctx, "srv_id = ?", srvId)
	var list []model.ServiceEnv
	user := meta.User(ctx)
	for _, envId := range envIds {
		list = append(list, model.ServiceEnv{
			EnvId:      envId,
			SrvId:      srvId,
			OperatorId: user.UserId,
			Operator:   user.UserName,
		})
	}
	return srv.CreateAll(ctx, list)
}

func UpdateService(ctx *gin.Context, in *types.UpdateServiceRequest) error {
	srv := model.Service{}
	if copier.Copy(&srv, in) != nil {
		return errors.AssignError
	}
	if err := srv.UpdateByID(ctx); err != nil {
		return err
	}
	return AddServiceEnv(ctx, srv.BaseModel.ID, in.EnvIds)
}

func DeleteService(ctx *gin.Context, in *types.DeleteServiceRequest) error {
	srv := model.Service{}
	return srv.DeleteByID(ctx, in.ID)
}
