package logic

import (
	"context"
	"devops/common/meta"
	"devops/common/tools"
	"devops/common/typex"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserRequest) error {
	user := models.User{}
	tools.Transform(req, &user)
	user.Password, _ = meta.ParsePwd(user.Password)
	if user.Password == req.Password {
		return errors.New("密码生成失败")
	}
	if err := user.Create(l.ctx); err != nil {
		return err
	}
	// 新添加的角色继承公共接口
	if user.OneByID() == nil {
		l.svcCtx.Rbac.AddGroupingPolicy(user.Keyword, typex.PublicRoleKey)
	}
	return nil
}
