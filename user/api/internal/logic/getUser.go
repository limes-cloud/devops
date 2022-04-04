package logic

import (
	"context"
	"devops/common/meta"
	"devops/common/tools"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser() (resp *types.GetUserResponse, err error) {
	userId := meta.UserId(l.ctx)
	resp = new(types.GetUserResponse)

	user := models.User{}
	if err = user.GetUser(userId); err != nil {
		return nil, errors.New("获取用户信息失败")
	}
	tools.Transform(user, &resp)
	return resp, nil
}
