package response

import "video/common/tools"

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResp(data interface{}) *Resp {
	return &Resp{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

func HandlerResp(in interface{}) interface{} {
	var data = make(map[string]interface{})
	tools.Transform(in, &data)
	if val, ok := data["data"]; ok {
		return NewResp(val)
	}
	return NewResp(data)
}
