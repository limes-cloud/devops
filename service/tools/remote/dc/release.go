package dc

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/limeschool/gin"
	"service/errors"
	"service/tools/remote/model"
	"time"
)

//DeleteService 删除服务
func (c *client) DeleteService(ctx *gin.Context, srv model.ServiceConfig) error {
	resp, err := ctx.Http().Option(func(request *resty.Request) *resty.Request {
		return request.SetHeader("Token", c.Token).SetBody(srv)
	}).Delete(c.Host + "/api/v1/service")

	if err != nil {
		return err
	}

	_, err = c.ParseResponse(resp.Body())
	if err != nil {
		return err
	}
	return nil
}

// CreateService 创建服务
func (c *client) CreateService(ctx *gin.Context, srv model.ServiceConfig) error {
	resp, err := ctx.Http().Option(func(request *resty.Request) *resty.Request {
		return request.SetHeader("Token", c.Token)
	}).PostJson(c.Host+"/api/v1/service", srv)

	if err != nil {
		return err
	}

	_, err = c.ParseResponse(resp.Body())
	if err != nil {
		return err
	}
	return nil
}

// GetServiceRelease 获取服务状态
func (c *client) GetServiceRelease(ctx *gin.Context, srv model.ServiceConfig) error {
	for {
		time.Sleep(2 * time.Second)
		info, err := c.ServiceRelease(ctx, srv)
		if err != nil {
			return err
		}
		if info["status"] == true {
			return nil
		}

		if info["state"] == "exited" || info["state"] == "restarting" {
			return errors.New("容器启动失败")
		}
	}
}

func (c *client) ServiceRelease(ctx *gin.Context, srv model.ServiceConfig) (map[string]any, error) {
	byteData, _ := json.Marshal(srv)
	query := map[string]string{}
	_ = json.Unmarshal(byteData, &query)

	resp, err := ctx.Http().Option(func(request *resty.Request) *resty.Request {
		return request.SetHeader("Token", c.Token).SetQueryParams(query)
	}).Get(c.Host + "/api/v1/service/release")

	if err != nil {
		return nil, err
	}

	data, err := c.ParseResponse(resp.Body())
	if err != nil {
		return nil, err
	}

	info, _ := data.(map[string]any)
	return info, nil
}

// GetServicePods 获取服务状态
func (c *client) GetServicePods(ctx *gin.Context, srv model.ServiceConfig) ([]model.Pod, error) {
	var resp []model.Pod
	byteData, _ := json.Marshal(srv)
	query := map[string]string{}
	_ = json.Unmarshal(byteData, &query)

	response, err := ctx.Http().Option(func(request *resty.Request) *resty.Request {
		return request.SetHeader("Token", c.Token).SetQueryParams(query)
	}).Get(c.Host + "/api/v1/service/pods")
	if err != nil {
		return nil, err
	}
	return resp, response.Result(&resp)
}
