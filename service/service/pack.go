package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"service/tools"
	"service/tools/exec"
	"strconv"
	"strings"
)

type pack struct {
	call           func(string)
	cmd            exec.Interface
	WorkDir        string            `json:"work_dir"`        // 工作目录
	GitUrl         string            `json:"git_url"`         // git代码地址
	GitUser        string            `json:"git_user"`        // git账号
	GitPass        string            `json:"git_pass"`        // git密码
	RegistryUrl    string            `json:"registry_url"`    // 仓库地址
	RegistryUser   string            `json:"registry_user"`   // 仓库账号
	RegistryPass   string            `json:"registry_pass"`   // 仓库密码
	ServiceName    string            `json:"service_name"`    // 服务名
	ServiceBranch  string            `json:"service_branch"`  // 服务分支
	ServiceVersion string            `json:"service_version"` // 服务版本
	Exec           string            `json:"exec"`            // 执行器
	Dockerfile     string            `json:"dockerfile"`      // 打包脚本
	Args           map[string]string `json:"args"`            // 打包脚本变量
}

type PackCall func(string)

func NewPack() *pack {
	return &pack{
		cmd:  exec.New(),
		call: nil,
	}
}

// HasWorkDir 是否存在工作目录
func (p *pack) HasWorkDir() bool {
	fi, err := os.Stat(p.WorkDir)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func (p *pack) SetWatch(f func(string)) {
	p.call = f
}

func (p *pack) GetImageName() string {
	return fmt.Sprintf("%v/%v:%v", p.RegistryUrl, p.ServiceName, p.ServiceVersion)
}

func (p *pack) GetServerWorkDir() string {
	name, _ := p.GetGitName()
	return p.WorkDir + "/" + name
}

// HasDevImage 是否存在本地镜像
func (p *pack) HasDevImage() bool {
	cmd := p.cmd.Command(p.Exec, "-c", fmt.Sprintf("docker images %v|wc -l", p.GetImageName()))
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	reg := regexp.MustCompile(`[1-9]+`)
	outStr := reg.FindString(string(out))
	length, _ := strconv.ParseInt(outStr, 10, 64)
	return length == 2
}

// HasRemoteImage 是否存在本地镜像
func (p *pack) HasRemoteImage() (bool, error) {
	shell := fmt.Sprintf("curl -X GET -u %v:%v %v/v2/%v/tags/list", p.RegistryUser, p.RegistryPass, p.RegistryUrl, p.ServiceName)
	cmd := p.cmd.Command(p.Exec, "-c", shell)
	out, err := cmd.Output()
	if err != nil {
		return false, errors.New("镜像仓库访问失败")
	}

	resp := struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}{}

	_ = json.Unmarshal(out, &resp)
	if resp.Name == "" {
		return false, nil
	}
	return tools.InList(resp.Tags, p.ServiceVersion), nil
}

// CreateWorkDir 创建工作目录
func (p *pack) CreateWorkDir() error {
	if !p.HasWorkDir() {
		if err := os.MkdirAll(p.WorkDir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// GetGitName 获取git代码名
func (p *pack) GetGitName() (string, error) {
	url := p.GitUrl
	if len(url) < 4 {
		return "", fmt.Errorf("%v 不是一个合法仓库地址", url)
	}
	if url[len(url)-4:] != ".git" {
		return "", fmt.Errorf("%v 不是一个仓库地址", url)
	}

	index := strings.LastIndex(url, "/")
	return url[index+1 : len(url)-4], nil
}

// RemoveServerPath 删除服务所在目录
func (p *pack) RemoveServerPath() {
	_ = os.RemoveAll(p.GetServerWorkDir())
}

// GetGitVersion 获取git版本
func (p *pack) GetGitVersion() (string, error) {
	cmd := p.cmd.Command(p.Exec, "-c", "git version")
	version, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(version), nil
}

// GetDockerVersion 获取docker版本
func (p *pack) GetDockerVersion() (string, error) {
	cmd := p.cmd.Command(p.Exec, "-c", "docker -v")
	version, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(version), nil
}

// GitCloneCode 进行代码拉去
func (p *pack) GitCloneCode() error {
	arr := strings.Split(p.GitUrl, "//")
	if len(arr) != 2 {
		return errors.New("git url error")
	}

	if p.GitUser != "" && p.GitPass != "" {
		p.GitUrl = fmt.Sprintf("%v//%v:%v@%v", arr[0], p.GitUser, p.GitPass, arr[1])
	}

	// 拉去代码
	cmd := p.cmd.Command(p.Exec, "-c", fmt.Sprintf("git clone %v", p.GitUrl))
	cmd.SetDir(p.WorkDir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git clone error :%v", string(out))
	}

	// 切换分支
	cmd = p.cmd.Command(p.Exec, "-c", fmt.Sprintf("git checkout %v", p.ServiceBranch))
	cmd.SetDir(p.GetServerWorkDir())
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git checkout error:%v", string(out))
	}

	return nil
}

// Pack 进行打包镜像
func (p *pack) Pack() error {
	// 渲染dockerfile
	p.RenderDockerfile()
	// 创建dockerfile
	err := os.WriteFile(p.GetServerWorkDir()+"/Dockerfile", []byte(p.Dockerfile), os.ModePerm)
	if err != nil {
		return fmt.Errorf("create dockerfile error :%v", err.Error())
	}

	// 执行dockerfile
	cmd := p.cmd.Command(p.Exec, "-c", fmt.Sprintf("docker build -t %v .", p.GetImageName()))
	cmd.SetDir(p.GetServerWorkDir())
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("docker build error :%v", string(out))
	}

	return nil
}

func (p *pack) RenderDockerfile() {
	reg := regexp.MustCompile(`\{\w+}`)
	list := reg.FindAllString(p.Dockerfile, -1)
	for _, val := range list {
		key := val[1 : len(val)-1]
		p.Dockerfile = strings.ReplaceAll(p.Dockerfile, val, p.Args[key])
	}
}

// Upload 进行镜像上传
func (p *pack) Upload() error {
	if !p.HasDevImage() {
		return fmt.Errorf("docker image not exist")
	}

	// 登陆仓库
	shell := fmt.Sprintf("docker login -u %v -p %v %v", p.RegistryUser, p.RegistryPass, p.RegistryUrl)
	cmd := p.cmd.Command(p.Exec, "-c", shell)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("docker login error :%v", string(out))
	}

	// 推送镜像
	cmd = p.cmd.Command(p.Exec, "-c", fmt.Sprintf("docker push %v", p.GetImageName()))
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("docker push error :%v", string(out))
	}
	return nil
}

func (p *pack) RemoveImage() {
	if !p.HasDevImage() {
		return
	}

	p.cmd.Command(p.Exec, "-c", fmt.Sprintf("docker rmi -f %v .", p.GetImageName()))
}

func (p *pack) Start() error {

	if _, err := p.GetDockerVersion(); err != nil {
		return err
	}

	if _, err := p.GetGitVersion(); err != nil {
		return err
	}

	// 清理工作痕迹
	defer func() {
		p.RemoveServerPath()
		p.RemoveImage()
	}()

	has, err := p.HasRemoteImage()
	if err != nil {
		return err
	}
	if has {
		return nil
	}

	// 防止人为创建
	p.RemoveServerPath()
	if err := p.CreateWorkDir(); err != nil {
		return err
	}

	// 拉取代码
	if err := p.GitCloneCode(); err != nil {
		return err
	}

	// 打包
	if err := p.Pack(); err != nil {
		return err
	}

	// 上传
	if err := p.Upload(); err != nil {
		return err
	}

	return nil
}
