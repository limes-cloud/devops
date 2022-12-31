package service

import (
	"bytes"
	"dc/errors"
	"dc/tools/exec"
	"dc/types"
	"fmt"
	"github.com/limeschool/gin"
	"io"
	"os"
)

func AddNetwork(ctx *gin.Context, in *types.AddNetworkRequest) (err error) {

	_ = DeleteNetwork(ctx, &types.DeleteNetworkRequest{Host: in.Host, ServiceName: in.ServiceName})

	certPath := ""
	keyPath := ""
	hostPath := ctx.Config.GetString("nginx_hosts_path")
	fileName := "template/http.conf"
	if in.Cert != "" && in.Key != "" {
		fileName = "template/https.conf"

		tlsDir := fmt.Sprintf("%v/ssl/%v", hostPath, in.ServiceName)
		if err = os.MkdirAll(tlsDir, os.ModePerm); err != nil {
			return err
		}

		// 写入 cert
		certPath = fmt.Sprintf("%v/%v.cert", tlsDir, in.Host)
		if err = os.WriteFile(certPath, []byte(in.Cert), os.ModePerm); err != nil {
			return err
		}

		// 写入key
		keyPath = fmt.Sprintf("%v/%v.key", tlsDir, in.Host)
		if err = os.WriteFile(keyPath, []byte(in.Key), os.ModePerm); err != nil {
			return err
		}
	}

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	template, _ := io.ReadAll(file)
	// 生成替换模板变量
	servicesInfo := ""
	for i := 0; i < in.Replicas; i++ {
		servicesInfo += fmt.Sprintf("\tserver  127.0.0.1:%v max_fails=5 fail_timeout=100;\n", in.RunPort+i)
	}

	// 替换模板变量
	template = bytes.Replace(template, []byte("[service]"), []byte(servicesInfo), -1)
	template = bytes.Replace(template, []byte("[service_name]"), []byte(in.ServiceName), -1)
	template = bytes.Replace(template, []byte("[host]"), []byte(in.Host), -1)
	if certPath != "" && keyPath != "" {
		template = bytes.Replace(template, []byte("[cert_path]"), []byte(certPath), -1)
		template = bytes.Replace(template, []byte("[key_path]"), []byte(keyPath), -1)
		extra := ""
		if in.Redirect {
			extra = `if ($server_port !~ 443){
				rewrite ^(/.*)$ https://$host$1 permanent;
			}`
		}
		template = bytes.Replace(template, []byte("[extra]"), []byte(extra), -1)
	}

	// 写入nginx配置
	if err = os.WriteFile(fmt.Sprintf("%v/%v.conf", hostPath, in.Host), template, os.ModePerm); err != nil {
		return err
	}

	// 执行nginx重启
	e := exec.New()
	execType := ctx.Config.GetDefaultString("exec_type", "/bin/sh")

	cmd := e.Command(execType, "-c", "nginx -s reload -c "+ctx.Config.GetString("nginx_conf_path"))
	byteData, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(byteData))
	}
	return nil
}

func DeleteNetwork(ctx *gin.Context, in *types.DeleteNetworkRequest) error {
	confPath := ctx.Config.GetString("nginx_conf_path")
	certPath := fmt.Sprintf("%v/ssl/%v/%v.cert", confPath, in.ServiceName, in.Host)
	keyPath := fmt.Sprintf("%v/ssl/%v/%v.key", confPath, in.ServiceName, in.Host)
	tempPath := fmt.Sprintf("%v/%v.conf", confPath, in.Host)
	_ = os.Remove(certPath)
	_ = os.Remove(keyPath)
	_ = os.Remove(tempPath)
	return nil
}
