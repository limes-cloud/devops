package logic

import (
	"context"
	"devops/common/meta"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"
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
	var err error
	user := models.User{}
	orm := l.svcCtx.Orm
	if req.Password != "" {
		req.Password, err = meta.ParsePwd(req.Password)
		if err != nil {
			return err
		}
	}
	orm.Table(user.Table()).Create(req)
	return nil
}
