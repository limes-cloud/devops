package address

import (
	"devops/common/tools"
	"devops/common/tools/request"
)

func GetAddress(ip string) string {
	if tools.CheckIP(ip) {
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
	if request.Get(url, &resp) != nil {
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
	request.Get(url, &resp)
	return resp.Addr
}

func IPApi(ip string) string {
	url := "ip-api.com/json/" + ip + "?lang=zh-CN"
	type response struct {
		RegionName string `json:"region_name"`
		City       string `json:"city"`
	}
	var resp response
	request.Get(url, &resp)
	return resp.RegionName + resp.City
}
