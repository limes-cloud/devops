package service

import (
	"github.com/go-resty/resty/v2"
	"github.com/limeschool/gin"
	"service/errors"
)

func UserTeamIds(ctx *gin.Context) ([]int64, error) {
	request := ctx.Http().Option(func(request *resty.Request) *resty.Request {
		request.Header.Set("Authorization", ctx.Request.Header.Get("Authorization"))
		return request
	})

	resp := struct {
		Code int
		Msg  string
		Data []int64
	}{}

	response, err := request.Get(ctx.Config.GetString("ums_addr"))
	if err != nil {
		return nil, err
	}

	if response.Result(&resp); err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, errors.New(resp.Msg)
	}

	return resp.Data, nil
}
