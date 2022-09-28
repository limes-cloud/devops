package service

import (
	"github.com/limeschool/gin"
	"notice/errors"
	"notice/model"
	"notice/types"
)

func PageLog(ctx *gin.Context, in *types.GetLogRequest) ([]model.Log, int64, error) {
	m := model.Log{}
	list, total, err := m.Page(ctx, in.Page, in.Count, in)
	if err != nil {
		return nil, 0, errors.DBError
	}
	return list, total, nil
}
