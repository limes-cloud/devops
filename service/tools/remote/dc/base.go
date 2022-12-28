package dc

import (
	"encoding/json"
	"service/errors"
)

type client struct {
	Host  string
	Token string
}

func NewClient(host, token string) (*client, error) {
	return &client{
		Host:  host,
		Token: token,
	}, nil
}

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (client) ParseResponse(data []byte) (any, error) {
	var resp response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, errors.New(resp.Msg)
	}
	return resp.Data, nil
}
