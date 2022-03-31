package tools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func Transform(src, dst interface{}) {
	b, _ := json.Marshal(src)
	_ = json.Unmarshal(b, dst)
}

func Get(url string, dst interface{}) error {
	cli := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := cli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	return json.Unmarshal(body, dst)
}

func Post(url string, data interface{}, dest interface{}) error {
	return post(url, data, dest, "application/x-www-form-urlencoded")
}

func PostJson(url string, data interface{}, dest interface{}) error {
	return post(url, data, dest, "application/json")
}

func post(url string, data interface{}, dest interface{}, ct string) error {
	cli := http.Client{
		Timeout: 5 * time.Second,
	}
	b, _ := json.Marshal(data)
	resp, err := cli.Post(url, ct, bytes.NewBuffer(b))
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, dest)
}

func InListStr(list []string, val string) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}
