package logic

import (
	"context"
	"devops/common/meta"
	"devops/user/models"
	"errors"

	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.UpdateUserPasswordRequest) error {
	user := models.User{}
	user.ID = meta.UserId(l.ctx)
	if user.OneByID() != nil {
		return errors.New("为查询到用户信息")
	}
	if !meta.CompareHashPwd(user.Password, req.Oldpass) {
		return errors.New("旧密码错误")
	}
	user.Password, _ = meta.ParsePwd(req.Pass)
	return user.Update(l.ctx)
}
