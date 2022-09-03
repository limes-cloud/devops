package service

import (
	"configure/model"
	"configure/types"
	"github.com/limeschool/gin"
)

func PageOperatorLog(ctx *gin.Context, in *types.PageServiceFieldRequest) ([]model.OperatorLog, int64, error) {
	srv := model.OperatorLog{}
	return srv.Page(ctx, in.Page, in.Count, in)
}
