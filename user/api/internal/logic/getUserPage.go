package logic

import (
	"context"
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

	tb := l.svcCtx.Orm.Table(user.Table()).Select("user.*,team.name team_name,role.name role_name")

	if req.TeamID != 0 {
		tb = tb.Where("team_id = ?", req.TeamID)
	}

	if req.RoleID != 0 {
		tb = tb.Where("role_id = ?", req.RoleID)
	}

	if req.Name == "" {
		tb = tb.Where("name = ?", req.Name)
	}

	if req.OperatorID != 0 {
		tb = tb.Where("operator_id = ?", req.OperatorID)
	}

	if req.Status != nil {
		tb = tb.Where("status = ?", req.Status)
	}

	return resp, tb.First(resp.List).Error
}
