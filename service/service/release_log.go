package service

import (
	"github.com/limeschool/gin"
	"service/consts"
	"service/errors"
	"service/model"
	"service/tools/image_release"
	"service/tools/image_release/k8s"
	"service/types"
)

func PageReleaseLog(ctx *gin.Context, in *types.PageReleaseLogRequest) ([]model.ReleaseLog, int64, error) {
	packLog := model.ReleaseLog{}
	return packLog.Page(ctx, in.Page, in.Count, in)
}

func AddReleaseLog(ctx *gin.Context, in *types.AddReleaseLogRequest) error {

	// 查询服务信息
	service := model.Service{}
	if err := service.OneByKeyword(ctx, in.ServiceKeyword); err != nil {
		return err
	}

	image := model.ImageRegistry{}
	if err := image.OneById(ctx, service.ImageRegistryID); err != nil {
		return err
	}

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

	// 模板解析
	template := service.Replace(release.Template)

	// 连接到远程
	var client image_release.ImageRelease

	if release.Type == consts.Dc {

	} else {
		if env.K8sHost == nil || env.K8sToken == nil {
			return errors.New("此环境未设置k8s配置信息")
		}
		client, err = k8s.NewK8sClient(*env.K8sHost, *env.K8sToken, *env.K8sNamespace)
		if err != nil {
			return err
		}
	}

	// 创建任务日志
	log := model.ReleaseLog{
		EnvKeyword:        env.Keyword,
		EnvName:           env.Name,
		ServiceKeyword:    service.Keyword,
		ServiceName:       service.Name,
		ImageName:         in.ImageName,
		ImageRegistryName: image.Name,
	}

	if err = log.Create(ctx); err != nil {
		return err
	}

	// 异步监听执行
	go func() {
		newLog := model.ReleaseLog{}
		newLog.ID = log.ID

		defer func() {
			newLog.IsFinish = true
			_ = newLog.UpdateByID(ctx)
		}()

		// 执行创建
		if err = client.UpdateFromYaml(ctx, template); err != nil {
			newLog.Desc = err.Error()
			newLog.Status = consts.Fail
		}

		// 监听启动状态

	}()

	return nil
}
