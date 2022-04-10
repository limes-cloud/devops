package logic

import (
	"context"
	"devops/user/models"

	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleMenuIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleMenuIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleMenuIdsLogic {
	return &GetRoleMenuIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleMenuIdsLogic) GetRoleMenuIds(req *types.GetRoleMenuIdsRequest) (resp *types.GetRoleMenuIdsResponse, err error) {
	var ids []int64
	m := models.RoleMenu{}
	resp = new(types.GetRoleMenuIdsResponse)
	list, _, err := m.All(map[string]int64{"role_id": req.RoleID})
	for _, item := range list {
		ids = append(ids, item.MenuID)
	}
	resp.List = ids
	return resp, err
}
