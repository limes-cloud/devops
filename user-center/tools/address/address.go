package address

import (
	"encoding/json"
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
	if address = IP360(ip); address != "" {
		return
	}
	if address = IPWhois(ip); address != "" {
		return
	}
	if address = IPApi(ip); address != "" {
		return
	}
	return "地址查询失败"
}

func IP360(ip string) string {
	type response struct {
		Errno int    `json:"errno"`
		Data  string `json:"data"`
	}
	var resp response
	url := "ip.360.cn/IPQuery/ipquery?ip=" + ip
	if Get(url, &resp) != nil {
		return ""
	}
	if resp.Errno == 0 {
		return resp.Data
	}
	return ""
}

func IPWhois(ip string) string {
	type response struct {
		Addr string `json:"addr"`
	}
	var resp response
	url := "whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	Get(url, &resp)
	return resp.Addr
}

func IPApi(ip string) string {
	url := "ip-api.com/json/" + ip + "?lang=zh-CN"
	type response struct {
		RegionName string `json:"region_name"`
		City       string `json:"city"`
	}
	var resp response
	Get(url, &resp)
	return resp.RegionName + resp.City
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

func CheckIP(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}
