package address

import (
	"bytes"
	"encoding/json"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func GetAddress(ip string) string {
	if CheckIP(ip) {
		if ip == "127.0.0.1" {
			return "本地登陆"
		}
		return GetAddressByIP(ip)
		//调用三方解析接口
	}
	return "非法ip地址"
}

func GetAddressByIP(ip string) (address string) {
	if address = IPWhois(ip); address != "" {
		return
	}
	return "地址查询失败"
}

func IPWhois(ip string) string {
	type response struct {
		Addr string `json:"addr"`
	}
	var resp response
	url := "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	Get(url, &resp, true)
	return resp.Addr
}

func Get(url string, dst interface{}, toUtf8 bool) error {
	cli := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := cli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if toUtf8 {
		body, _ = GbkToUtf8(body)
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(body, dst)
}

func CheckIP(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
