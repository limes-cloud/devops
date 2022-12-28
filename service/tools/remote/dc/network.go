package dc

import (
	"github.com/go-resty/resty/v2"
	"github.com/limeschool/gin"
	"service/tools/remote/model"
)

// CreateNetwork 新增网络
func (c *client) CreateNetwork(ctx *gin.Context, config model.NetworkConfig) error {
	resp, err := ctx.Http().Option(func(request *resty.Request) *resty.Request {
		return request.SetHeader("Token", c.Token)
	}).PostJson(c.Host+"/api/v1/network", config)

	if err != nil {
		return err
	}

	_, err = c.ParseResponse(resp.Body())
	if err != nil {
		return err
	}
	return nil
}

// DeleteNetwork 删除网络
func (c *client) DeleteNetwork(ctx *gin.Context, config model.NetworkConfig) error {
	resp, err := ctx.Http().Option(func(request *resty.Request) *resty.Request {
		return request.SetHeader("Token", c.Token).SetBody(config)
	}).Delete(c.Host + "/api/v1/network")

	if err != nil {
		return err
	}

	_, err = c.ParseResponse(resp.Body())
	if err != nil {
		return err
	}
	return nil

}
