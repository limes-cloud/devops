package service

import (
	"fmt"
	"github.com/limeschool/gin"
	"regexp"
	"service/consts"
	"service/errors"
	"service/model"
	"service/tools/remote"
	remoteModel "service/tools/remote/model"
	"service/types"
	"strings"
	"time"
)

func PageReleaseLog(ctx *gin.Context, in *types.PageReleaseLogRequest) ([]model.ReleaseLog, int64, error) {
	packLog := model.ReleaseLog{}
	return packLog.Page(ctx, in.Page, in.Count, in)
}

func AddReleaseLog(ctx *gin.Context, in *types.AddReleaseLogRequest) error {
	// 查询打包日志信息
	pack := model.PackLog{}
	if err := pack.OneById(ctx, in.PackID); err != nil {
		return err
	}
	if (pack.Status == nil || !*pack.Status) || pack.IsClear || pack.ServiceKeyword != in.ServiceKeyword {
		return errors.New("未找到构建镜像信息")
	}

	// 查询服务信息
	service := model.Service{}
	if err := service.OneByKeyword(ctx, in.ServiceKeyword); err != nil {
		return err
	}

	image := model.ImageRegistry{}
	if err := image.OneById(ctx, service.ImageRegistryID); err != nil {
		return err
	}
	service.ImageRegistryName = image.Name

	// 查询环境信息
	env := model.Environment{}
	if err := env.OneByKeyword(ctx, in.EnvKeyword); err != nil {
		return err
	}

	// 查询服务所属环境
	srvEnv := model.ServiceEnv{}
	srvEnvs, err := srvEnv.All(ctx, "srv_id = ?", service.ID)
	if err != nil {
		return err
	}

	// 判断服务是否存在所属环境中
	isExists := false
	for _, item := range srvEnvs {
		if item.EnvId == env.ID {
			isExists = true
		}
	}

	if !isExists {
		return errors.NewF("该服务不归属环境%v(%v)", env.Name, env.Keyword)
	}

	// 查询发布模板
	release := model.Release{}
	if err = release.OneById(ctx, service.ReleaseID); err != nil {
		return err
	}
	if release.Type != env.Type {
		return errors.NewF("%v只支持%v发布方式", env.Name, env.Type)
	}

	// 模板解析
	template := Replace(env, service, pack, release.Template)

	// 连接到远程
	client, err := remote.NewClient(env.Type, env.Host, env.Token)
	if err != nil {
		return err
	}
	// 创建任务日志
	log := model.ReleaseLog{
		EnvKeyword:        env.Keyword,
		EnvName:           env.Name,
		ServiceKeyword:    service.Keyword,
		ServiceName:       service.Name,
		ImageName:         pack.ImageName,
		ImageRegistryName: image.Name,
	}

	if err = log.Create(ctx); err != nil {
		return err
	}

	// 异步监听执行
	go func() {
		newLog := model.ReleaseLog{}
		newLog.ID = log.ID
		startTime := time.Now().Unix()
		config := remoteModel.ServiceConfig{
			Type:          service.RunType,
			Yaml:          template,
			Namespace:     env.Namespace,
			ServiceName:   service.Keyword,
			Replicas:      service.Replicas,
			RunPort:       service.RunPort,
			ListenPort:    service.ListenPort,
			ImageUser:     image.Username,
			ImagePass:     image.Password,
			ImageRegistry: image.Host,
		}

		defer func() {
			newLog.IsFinish = true
			newLog.UseTime = time.Now().Unix() - startTime
			_ = newLog.UpdateByID(ctx)
		}()

		// 执行创建
		if err = client.CreateService(ctx, config); err != nil {
			newLog.Desc = err.Error()
			newLog.Status = consts.Fail
			return
		}

		// 监听启动状态
		if err = client.GetServiceRelease(ctx, config); err != nil {
			newLog.Desc = err.Error()
			newLog.Status = consts.Fail
		} else {
			newLog.Desc = "发布成功"
			newLog.Status = consts.Success
		}
	}()

	return nil
}

func Replace(env model.Environment, srv model.Service, pack model.PackLog, text string) string {
	reg := regexp.MustCompile(`\{\w+}`)
	args := reg.FindAllString(text, -1)
	for _, val := range args {
		key := val[1 : len(val)-1]
		key = strings.ReplaceAll(key, " ", "")
		switch key {
		case consts.RunPort:
			text = strings.ReplaceAll(text, val, fmt.Sprint(srv.RunPort))
		case consts.ListenPort:
			text = strings.ReplaceAll(text, val, fmt.Sprint(srv.ListenPort))
		case consts.Replicas:
			text = strings.ReplaceAll(text, val, fmt.Sprint(srv.Replicas))
		case consts.ProbeValue:
			text = strings.ReplaceAll(text, val, srv.ProbeValue)
		case consts.ProbeInitDelay:
			text = strings.ReplaceAll(text, val, fmt.Sprint(srv.ProbeInitDelay))
		case consts.ServiceName:
			text = strings.ReplaceAll(text, val, srv.Keyword)
		case consts.Image:
			text = strings.ReplaceAll(text, val, pack.ImageName)
		case consts.Namespace:
			namespace := "default"
			if env.Namespace != "" {
				namespace = env.Namespace
			}
			text = strings.ReplaceAll(text, val, namespace)
		}
	}
	return text
}
