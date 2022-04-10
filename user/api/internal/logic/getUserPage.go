package logic

import (
	"context"
	"devops/common/tools"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPageLogic {
	return &GetUserPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPageLogic) GetUserPage(req *types.GetUserPageRequest) (resp *types.GetUserPageResponse, err error) {
	resp = new(types.GetUserPageResponse)
	user := models.User{}
	list, total, err := user.Page(req, req.Page, req.Count)
	tools.Transform(list, &resp.List)
	resp.Total = total
	return resp, err
}
