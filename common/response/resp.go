package response

import "devops/common/tools"

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewResp(data interface{}) *Resp {
	return &Resp{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

func HandlerResp(in interface{}) interface{} {
	if in == nil {
		return NewResp(nil)
	}
	var data = make(map[string]interface{})
	tools.Transform(in, &data)
	if val, ok := data["data"]; ok {
		return NewResp(val)
	}
	return NewResp(data)
}
