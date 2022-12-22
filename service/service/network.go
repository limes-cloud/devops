package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"service/consts"
	"service/errors"
	"service/model"
	"service/tools/image_release"
	"service/tools/image_release/k8s"
	"service/types"
)

func PageNetwork(ctx *gin.Context, in *types.PageNetworkRequest) ([]model.Network, int64, error) {
	network := model.Network{}
	list, total, err := network.Page(ctx, in.Page, in.Count, in)
	if err != nil {
		return list, total, err
	}
	for key, item := range list {
		// 查询service
		service := model.Service{}
		_ = service.OneById(ctx, item.SrvID)

		// 查询env
		env := model.Environment{}
		_ = env.OneById(ctx, item.EnvID)

		list[key].ServiceName = service.Name
		list[key].EnvName = env.Name
	}
	return list, total, err
}

func AddNetwork(ctx *gin.Context, in *types.AddNetworkRequest) error {
	network := model.Network{}
	if copier.Copy(&network, in) != nil {
		return errors.AssignError
	}

	// 查询服务
	service := model.Service{}
	if err := service.OneById(ctx, in.SrvID); err != nil {
		return err
	}

	// 查询环境
	env := model.Environment{}
	if err := env.OneById(ctx, in.EnvID); err != nil {
		return err
	}

	// 连接
	var client image_release.ImageRelease
	var err error
	if env.Type == consts.Dc {

	} else {
		client, err = k8s.NewK8sClient(env.K8sHost, env.K8sToken, env.K8sNamespace)
		if err != nil {
			return err
		}
	}
	// 生成config
	config := image_release.NetworkConfig{
		Namespace:  env.K8sNamespace,
		Service:    service.Keyword,
		Host:       in.Host,
		Cert:       in.Cert,
		Key:        in.Key,
		Redirect:   in.Redirect,
		TargetPort: service.ListenPort,
		RunPort:    service.RunPort,
	}
	_ = client.DeleteNetwork(ctx, config)
	if err = client.AddNetwork(ctx, config); err != nil {
		return err
	}

	return network.Create(ctx)
}

func UpdateNetwork(ctx *gin.Context, in *types.UpdateNetworkRequest) error {
	network := model.Network{}
	if copier.Copy(&network, in) != nil {
		return errors.AssignError
	}

	// 清除旧版本ingress
	old := model.Network{}
	if err := old.OneById(ctx, in.ID); err != nil {
		return err
	}
	oldService := model.Service{}
	if err := oldService.OneById(ctx, old.SrvID); err != nil {
		return err
	}

	// 查询旧版环境
	oldEnv := model.Environment{}
	if err := oldEnv.OneById(ctx, old.EnvID); err != nil {
		return err
	}

	// 查询旧版服务
	service := model.Service{}
	if err := service.OneById(ctx, in.SrvID); err != nil {
		return err
	}

	// 查询环境
	env := model.Environment{}
	if err := env.OneById(ctx, in.EnvID); err != nil {
		return err
	}

	// 连接
	var client image_release.ImageRelease
	var err error
	if env.Type == consts.Dc {

	} else {
		client, err = k8s.NewK8sClient(env.K8sHost, env.K8sToken, env.K8sNamespace)
		if err != nil {
			return err
		}
	}

	// 生成config
	config := image_release.NetworkConfig{
		Namespace:  oldEnv.K8sNamespace,
		Service:    oldService.Keyword,
		Host:       old.Host,
		Cert:       old.Cert,
		Key:        old.Key,
		Redirect:   old.Redirect,
		TargetPort: oldService.ListenPort,
		RunPort:    oldService.RunPort,
	}
	// 删除旧版本
	_ = client.DeleteNetwork(ctx, config)

	// 新增新版本
	config.Namespace = env.K8sNamespace
	config.Service = service.Keyword
	config.TargetPort = service.ListenPort
	config.RunPort = service.RunPort
	if in.Host != "" {
		config.Host = in.Host
	}

	if in.Cert != nil {
		config.Cert = *in.Cert
	}

	if in.Key != nil {
		config.Key = *in.Key
	}

	if in.Redirect != nil {
		config.Redirect = *in.Redirect
	}

	if err = client.AddNetwork(ctx, config); err != nil {
		return err
	}

	return network.UpdateByID(ctx)
}

func DeleteNetwork(ctx *gin.Context, in *types.DeleteNetworkRequest) error {
	network := model.Network{}
	if err := network.OneById(ctx, in.ID); err != nil {
		return err
	}

	service := model.Service{}
	if err := service.OneById(ctx, network.SrvID); err != nil {
		return err
	}

	// 查询环境
	env := model.Environment{}
	if err := env.OneById(ctx, network.EnvID); err != nil {
		return err
	}

	config := image_release.NetworkConfig{
		Namespace:  env.Keyword,
		Service:    env.Keyword,
		Host:       network.Host,
		Cert:       network.Cert,
		Key:        network.Key,
		Redirect:   network.Redirect,
		TargetPort: service.ListenPort,
		RunPort:    service.RunPort,
	}

	var client image_release.ImageRelease
	var err error
	if env.Type == consts.Dc {

	} else {
		client, err = k8s.NewK8sClient(env.K8sHost, env.K8sToken, env.K8sNamespace)
		if err != nil {
			return err
		}
	}

	if err = client.DeleteNetwork(ctx, config); err != nil {
		return err
	}
	// 执行yaml进行删除
	return network.DeleteByID(ctx, in.ID)
}
