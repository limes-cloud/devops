package gitlab

import (
	"errors"
	"github.com/xanzy/go-gitlab"
	"service/tools/code_registry/model"
)

type client struct {
	vm   *gitlab.Client
	user *model.User
}

func NewClient(token, url string) (*client, error) {
	var err error
	cli, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	c := &client{vm: cli}
	if c.user, err = c.CurUser(); err != nil {
		return nil, errors.New("token 验证失败")
	}
	return c, err
}

// CurUser 获取当前用户信息
func (c *client) CurUser() (*model.User, error) {
	user, _, err := c.vm.Users.CurrentUser()
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:        int64(user.ID),
		Name:      user.Name,
		AvatarUrl: user.AvatarURL,
		SiteAdmin: user.IsAdmin,
		Email:     user.Email,
	}, err
}

func (c *client) GetRepo(owner, repo string) (*model.Project, error) {

	var group *gitlab.Group
	groups, _, err := c.vm.Groups.ListGroups(&gitlab.ListGroupsOptions{Search: &owner})
	if err != nil {
		return nil, errors.New("获取命名空间失败:" + err.Error())
	}
	for _, item := range groups {
		if item.Name == owner {
			group = item
		}
	}
	if group == nil {
		return nil, errors.New("获取命名空间失败")
	}

	var project *gitlab.Project
	projects, _, err := c.vm.Groups.ListGroupProjects(group.ID, &gitlab.ListGroupProjectsOptions{
		Search: &repo,
	})
	if err != nil {
		return nil, err
	}
	for _, item := range projects {
		if item.Name == repo {
			project = item
		}
	}
	if project == nil {
		return nil, errors.New("获取项目失败")
	}

	return &model.Project{
		ID:        int64(project.ID),
		FullName:  project.PathWithNamespace,
		Desc:      project.Description,
		Name:      project.Name,
		IsPrivate: !project.Public,
		Url:       project.HTTPURLToRepo,
		Ssh:       project.SSHURLToRepo,
		Owner:     owner,
		Repo:      repo,
	}, err
}

func (c *client) GetRepoTags(project *model.Project) ([]model.Tag, error) {
	tags, _, err := c.vm.Tags.ListTags(int(project.ID), nil, nil)
	if err != nil {
		return nil, err
	}

	var mt []model.Tag
	for _, item := range tags {
		mt = append(mt, model.Tag{
			Name:       item.Name,
			CommitID:   item.Commit.ID,
			CommitTime: item.Commit.CommittedDate.String(),
			CommitName: item.Commit.CommitterName,
		})
	}
	return mt, nil
}

func (c *client) GetRepoBranches(project *model.Project) ([]model.Branch, error) {
	tags, _, err := c.vm.Branches.ListBranches(int(project.ID), nil, nil)
	if err != nil {
		return nil, err
	}

	var mt []model.Branch
	for _, item := range tags {
		mt = append(mt, model.Branch{
			Name:       item.Name,
			CommitID:   item.Commit.ID,
			CommitTime: item.Commit.CommittedDate.String(),
			CommitName: item.Commit.CommitterName,
		})
	}
	return mt, nil
}
