package github

import (
	"context"
	"errors"
	"github.com/google/go-github/v48/github"
	"github.com/jinzhu/copier"
	"golang.org/x/oauth2"
	"service/tools/code_registry/model"
	"time"
)

type client struct {
	vm   *github.Client
	user *model.User
}

func NewClient(token string) (*client, error) {
	var err error
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	cli := &client{vm: github.NewClient(tc)}
	if cli.user, err = cli.CurUser(); err != nil {
		return nil, errors.New("token 验证失败")
	}
	return cli, nil
}

// CurUser 获取当前用户信息
func (c *client) CurUser() (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	user, _, err := c.vm.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	mu := model.User{}
	_ = copier.Copy(&mu, user)

	return &mu, err
}

func (c *client) GetRepo(owner, repo string) (*model.Project, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	project, _, err := c.vm.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, err
	}

	return &model.Project{
		ID:        project.GetID(),
		Owner:     owner,
		Repo:      repo,
		Desc:      project.GetDescription(),
		FullName:  project.GetFullName(),
		Name:      project.GetName(),
		IsPrivate: project.GetPrivate(),
		Url:       project.GetHTMLURL(),
		Ssh:       project.GetSSHURL(),
	}, nil
}

func (c *client) GetRepoTags(project *model.Project) ([]model.Tag, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	tags, _, err := c.vm.Repositories.ListTags(ctx, project.Owner, project.Repo, nil)
	if err != nil {
		return nil, errors.New("获取项目失败:" + err.Error())
	}

	var mt []model.Tag
	for _, item := range tags {
		mt = append(mt, model.Tag{
			Name:       item.GetName(),
			CommitID:   item.Commit.GetSHA(),
			CommitTime: item.Commit.Committer.GetDate().String(),
			CommitName: item.Commit.GetCommitter().GetName(),
		})
	}
	return mt, nil
}

func (c *client) GetRepoBranches(project *model.Project) ([]model.Branch, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	branches, _, err := c.vm.Repositories.ListBranches(ctx, project.Owner, project.Repo, nil)
	if err != nil {
		return nil, err
	}

	var mt []model.Branch
	for _, item := range branches {
		mt = append(mt, model.Branch{
			Name:       item.GetName(),
			CommitID:   item.Commit.GetSHA(),
			CommitTime: item.Commit.GetCommit().GetAuthor().GetDate().String(),
			CommitName: item.Commit.GetCommitter().GetName(),
		})
	}
	return mt, nil
}
