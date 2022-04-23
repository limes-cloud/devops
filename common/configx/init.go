package configx

import (
	"devops/common/drive/etcx"
	"devops/common/tools"
	"encoding/json"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"io"
	"net/http"
	"os"
	"time"
)

// InitConfig 2:获取链接配置
func InitConfig(serviceName string, watch etcx.CallFunc) *viper.Viper {
	info := GetEtc()
	info.Prefix = info.Prefix + serviceName
	v := etcx.Init(info)
	data := map[string]interface{}{}
	if err := v.Unmarshal(&data); err != nil {
		panic(err)
	}
	etcx.CallBack = watch
	return v
}

// GetEtc 1:通过url获取配置信息
func GetEtc() etcx.EtcEnv {
	configAddr := os.Getenv("CONFIG_ADDR")
	if configAddr == "" {
		panic("环境变量：配置中心地址未配置")
	}
	env := os.Getenv("ENV")
	if env == "" {
		panic("环境变量：当前环境未配置")
	}
	url := configAddr + "/api/v1/cms/envconfig/info?env=" + env
	client := http.Client{Timeout: 10 * time.Second}
	response, err := client.Get(url)
	if err != nil {
		panic("请求配置中心信息异常" + err.Error())
	}
	defer response.Body.Close()
	respData := struct {
		Code int64       `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{}

	b, _ := io.ReadAll(response.Body)
	info := etcx.EtcEnv{}
	if json.Unmarshal(b, &respData) != nil {
		panic("解析配置中心失败")
	}
	if respData.Code != 200 {
		panic("获取配置连接信息失败:" + respData.Msg)
	}
	tools.Transform(respData.Data, &info)
	return info
}
