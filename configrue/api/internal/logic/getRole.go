package logic

import (
	"configure/models"
	"context"

	"configure/api/internal/svc"
	"configure/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleLogic) GetRole(req *types.GetRoleRequest) (resp *types.GetRoleResponse, err error) {
	role := models.Role{}
	resp = new(types.GetRoleResponse)

	tb := l.svcCtx.Orm.Table(role.Table())
	if req.Name != "" {
		tb = tb.Where("name = ?", req.Name)
	}

	if req.KeyWord != "" {
		tb = tb.Where("keyword = ?", req.KeyWord)
	}

	if req.Status != nil {
		tb = tb.Where("status = ?", req.Status)
	}

	if req.OperatorID != 0 {
		tb = tb.Where("operator_id = ?", req.OperatorID)
	}

	return resp, tb.Find(&resp.List).Error
}
