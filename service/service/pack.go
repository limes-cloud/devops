package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/limeschool/gin"
	"service/consts"
	"service/errors"
	"service/model"
	"service/tools"
	"service/tools/code_pack"
	"service/tools/code_registry"
	"service/tools/lock"
	"service/types"
	"strings"
	"time"
)

func PagePackLog(ctx *gin.Context, in *types.PagePackRequest) ([]model.PackLog, int64, error) {
	packLog := model.PackLog{}
	return packLog.Page(ctx, in.Page, in.Count, in)
}

func AddPack(ctx *gin.Context, in *types.AddPackRequest) error {
	in.CommitID = tools.GetCommitID(in.CommitID)

	log := model.PackLog{}
	if copier.Copy(&log, in) != nil {
		return errors.AssignError
	}

	// 获取服务信息
	service := model.Service{}
	if err := service.OneByKeyword(ctx, in.ServiceKeyword); err != nil {
		return err
	}

	// 获取dockerfile 信息
	dockerfile := model.Dockerfile{}
	if err := dockerfile.OneById(ctx, service.DockerfileID); err != nil {
		return err
	}

	// 获取代码信息
	code := model.CodeRegistry{}
	if err := code.OneById(ctx, service.CodeRegistryID); err != nil {
		return err
	}

	// 获取镜像信息
	image := model.ImageRegistry{}
	if err := image.OneById(ctx, service.ImageRegistryID); err != nil {
		return err
	}

	client, err := code_registry.NewCodeRegistry(code.Type, code.Host, code.Token)
	if err != nil {
		return err
	}

	var (
		url   string
		isShh bool
	)

	project, err := client.GetRepo(service.Owner, service.Repo)
	if err != nil {
		return err
	}
	// 获取代码仓库信息
	if code.CloneType == consts.SSHURL {
		isShh = true
		url = project.Ssh
	}

	if code.CloneType == consts.HTMLURL {
		isShh = false
		url = project.Url
	}

	// 进行打包数据控制
	key := fmt.Sprintf("lock_pack_%v_%v", in.ServiceKeyword, in.CloneValue)
	lk := lock.NewLockWithDuration(ctx, key, time.Minute*15)
	if !lk.TryAcquire() {
		return errors.New("存在相同任务正在处理中")
	}

	// 创建打包任务
	log.DockerfileName = dockerfile.Name
	log.ServiceName = service.Name
	log.ImageRegistryName = image.Name
	log.CodeRegistryName = code.Name
	log.ImageRegistryID = image.ID
	if err := log.Create(ctx); err != nil {
		lk.Release()
		return err
	}

	// 获取需要清理的数据
	historyList := log.OldHistory(ctx, image.HistoryCount)

	// 进行异步打包
	go func() {
		newLog := model.PackLog{}
		newLog.ID = log.ID
		newLog.Status = tools.Bool(true)

		var argStr []string
		startTime := time.Now().Unix()
		pk := code_pack.NewPack()
		pk.WorkDir = consts.WorkDir
		pk.GitUrl = url
		pk.IsSsh = isShh
		pk.GitToken = code.Token
		pk.RegistryUrl = image.Host
		pk.RegistryUser = image.Username
		pk.RegistryPass = image.Password
		pk.ServerName = service.Keyword
		pk.ServerBranch = in.CloneValue
		pk.ServerVersion = in.CommitID
		pk.RegistryProtocol = image.Protocol
		pk.Exec = consts.Exec
		pk.Dockerfile = dockerfile.Template
		pk.Args = map[string]string{
			consts.ListenPort: fmt.Sprint(service.ListenPort),
			consts.RunPort:    fmt.Sprint(service.RunPort),
		}

		// 进行镜像清理
		for _, item := range historyList {
			pk.RemoveRemoteImage(item.ServiceName, item.CommitID)
		}

		defer func() {
			lk.Release()
			newLog.Desc = strings.Join(argStr, "\n")
			newLog.IsFinish = true
			newLog.UseTime = time.Now().Unix() - startTime
			newLog.ImageName = pk.GetImageName()
			_ = newLog.UpdateByID(ctx)
		}()

		// 组装打包详细信息
		pk.SetWatch(func(s string) {
			s = strings.ReplaceAll(s, "\n", "")
			s = fmt.Sprintf("%v || %v", s, time.Now().Format("2006-01-02 15:04:05"))
			argStr = append(argStr, s)
		})

		if err = pk.Start(); err != nil {
			newLog.Status = tools.Bool(false)
		}

	}()

	return nil
}
