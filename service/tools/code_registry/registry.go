package code_registry

import (
	"service/consts"
	"service/errors"
	"service/tools/code_registry/gitee"
	"service/tools/code_registry/github"
	"service/tools/code_registry/gitlab"
	"service/tools/code_registry/model"
)

type CodeRegistry interface {
	// GetRepo 获取指定仓库
	GetRepo(owner, repo string) (*model.Project, error)

	// CurUser 获取当前用户
	CurUser() (*model.User, error)

	// GetRepoTags 获取指定项目的所有标签
	GetRepoTags(*model.Project) ([]model.Tag, error)

	// GetRepoBranches 获取指定项目的所有标签
	GetRepoBranches(*model.Project) ([]model.Branch, error)
}

func NewCodeRegistry(tp, host, token string) (CodeRegistry, error) {
	var codeRegistry CodeRegistry
	var err error
	switch tp {
	case consts.GITLAB:
		codeRegistry, err = gitlab.NewClient(token, host)
	case consts.GITEE:
		codeRegistry, err = gitee.NewClient(token, host)
	case consts.GITHUB:
		codeRegistry, err = github.NewClient(token)
	default:
		err = errors.New("不支持的代码仓库")
	}
	return codeRegistry, err
}
