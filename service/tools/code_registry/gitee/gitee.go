package gitee

import (
	"context"
	"errors"
	"github.com/limeschool/go-gitee/gitee"
	"golang.org/x/oauth2"
	"service/tools/code_registry/model"
	"time"
)

type client struct {
	vm   *gitee.APIClient
	user *model.User
}

// NewClient 仓库客户端
func NewClient(token, url string) (*client, error) {
	var err error
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	conf := gitee.NewConfiguration()
	if url != "" {
		conf.BasePath = url
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	conf.HTTPClient = oauth2.NewClient(ctx, ts)
	cli := &client{vm: gitee.NewAPIClient(conf)}
	if cli.user, err = cli.CurUser(); err != nil {
		return nil, errors.New("token 验证失败")
	}
	return cli, nil
}

// CurUser 获取当前用户信息
func (c *client) CurUser() (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	user, _, err := c.vm.UsersApi.GetV5User(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        int64(user.Id),
		Name:      user.Name,
		AvatarUrl: user.AvatarUrl,
		SiteAdmin: user.SiteAdmin,
		Email:     user.Email,
	}, err
}

func (c *client) GetRepo(owner, repo string) (*model.Project, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	project, _, err := c.vm.RepositoriesApi.GetV5ReposOwnerRepo(ctx, owner, repo, nil)
	if err != nil {
		return nil, errors.New("获取项目失败:" + err.Error())
	}
	return &model.Project{
		ID:        int64(project.Id),
		FullName:  project.FullName,
		Desc:      project.Description,
		Name:      project.Name,
		Owner:     owner,
		Repo:      repo,
		IsPrivate: project.Private,
		Url:       project.HtmlUrl,
		Ssh:       project.SshUrl,
	}, err
}

func (c *client) GetRepoTags(project *model.Project) ([]model.Tag, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	tags, _, err := c.vm.RepositoriesApi.GetV5ReposOwnerRepoTags(ctx, project.Owner, project.Repo, nil)
	if err != nil {
		return nil, err
	}

	var mt []model.Tag
	for _, item := range tags {
		commit, _, err := c.vm.RepositoriesApi.GetV5ReposOwnerRepoCommitsSha(ctx, project.Owner, project.Repo, item.Commit.Sha, nil)
		if err != nil {
			continue
		}

		mt = append(mt, model.Tag{
			Name:       item.Name,
			CommitID:   item.Commit.Sha,
			CommitTime: item.Commit.Date,
			CommitName: commit.Commit.Author.Name,
		})
	}
	return mt, nil
}

func (c *client) GetRepoBranches(project *model.Project) ([]model.Branch, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	branches, _, err := c.vm.RepositoriesApi.GetV5ReposOwnerRepoBranches(ctx, project.Owner, project.Repo, nil)
	if err != nil {
		return nil, errors.New("获取分支失败:" + err.Error())
	}

	var mt []model.Branch
	for _, item := range branches {
		commit, _, err := c.vm.RepositoriesApi.GetV5ReposOwnerRepoCommitsSha(ctx, project.Owner, project.Repo, item.Commit.Sha, nil)
		if err != nil {
			continue
		}

		mt = append(mt, model.Branch{
			Name:       item.Name,
			CommitID:   item.Commit.Sha,
			CommitTime: commit.Commit.Author.Date.String(),
			CommitName: commit.Commit.Author.Name,
		})
	}
	return mt, nil
}
