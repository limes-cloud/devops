package logic

import (
	"context"
	"devops/common/tools"
	"devops/user/models"
	"gorm.io/gorm"

	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogLogic {
	return &GetLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogLogic) GetLog(req *types.GetLogRequest) (resp *types.GetLogResponse, err error) {
	log := models.LoginLog{}
	resp = new(types.GetLogResponse)
	list, count, err := log.Page(req, req.Page, req.Count, func(db *gorm.DB) *gorm.DB {
		if req.Start != 0 && req.End != 0 {
			return db.Where("created_at between ? and ?", req.Start, req.End)
		}
		return db
	})
	tools.Transform(list, &resp.List)
	resp.Total = count
	return resp, err
}
