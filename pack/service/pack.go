package service

import (
	"fmt"
	"os"
	"pack/tools/exec"
	"regexp"
	"strconv"
	"strings"
)

type pack struct {
	cmd           exec.Interface
	WorkDir       string   `json:"work_dir"`       // 工作目录
	GitUrl        string   `json:"git_url"`        // git代码地址
	RegistryUrl   string   `json:"registry_url"`   // 仓库地址
	RegistryUser  string   `json:"registry_user"`  // 仓库账号
	RegistryPass  string   `json:"registry_pass"`  // 仓库密码
	ServerName    string   `json:"server_name"`    // 服务名
	ServerBranch  string   `json:"server_branch"`  // 服务分支
	ServerVersion string   `json:"server_version"` // 服务版本
	Exec          string   `json:"exec"`           // 执行器
	Dockerfile    string   `json:"dockerfile"`     // 打包脚本
	Args          []string `json:"args"`           // 打包脚本变量
}

func NewPack() *pack {
	return &pack{
		cmd: exec.New(),
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

func (p *pack) GetImageName() string {
	return fmt.Sprintf("%v/%v:%v", p.RegistryUrl, p.ServerName, p.ServerVersion)
}

func (p *pack) GetServerWorkDir() string {
	name, _ := p.GetGitName()
	return p.WorkDir + "/" + name
}

// HasImage 是否存在镜像
func (p *pack) HasImage() bool {
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

// CreateWorkDir 创建工作目录
func (p *pack) CreateWorkDir() error {
	if !p.HasWorkDir() {
		if err := os.MkdirAll(p.WorkDir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

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
	// 拉去代码
	cmd := p.cmd.Command(p.Exec, "-c", fmt.Sprintf("git clone %v", p.GitUrl))
	cmd.SetDir(p.WorkDir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git clone error :%v", string(out))
	}

	// 切换分支
	cmd = p.cmd.Command(p.Exec, "-c", fmt.Sprintf("git checkout %v", p.ServerBranch))
	cmd.SetDir(p.GetServerWorkDir())
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git checkout error:%v", string(out))
	}

	return nil
}

// Pack 进行打包镜像
func (p *pack) Pack() error {
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

// Upload 进行镜像上传
func (p *pack) Upload() error {
	if !p.HasImage() {
		return fmt.Errorf("docker image not exist")
	}

	// 登陆仓库
	shell := fmt.Sprintf("docker login -u %v -p %v %v", p.RegistryUser, p.RegistryPass, p.RegistryUrl)
	cmd := p.cmd.Command(p.Exec, "-c", shell)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("docker login error :%v", string(out))
	}

	// 推送镜像
	cmd = p.cmd.Command("/bin/sh", "-c", fmt.Sprintf("docker push %v", p.GetImageName()))
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("docker push error :%v", string(out))
	}
	return nil
}

func (p *pack) RemoveImage() {
	if !p.HasImage() {
		return
	}

	p.cmd.Command(p.Exec, "-c", fmt.Sprintf("docker rmi -f %v .", p.GetImageName()))
}

func (p *pack) Start() error {
	// 环境检测
	dv, err := p.GetDockerVersion()
	if err != nil {
		return err
	}
	fmt.Println("docker version", dv)

	gv, err := p.GetGitVersion()
	if err != nil {
		return err
	}
	fmt.Println("git version", gv)

	// 清理工作痕迹
	defer func() {
		p.RemoveServerPath()
		p.RemoveImage()
	}()

	if p.HasImage() {
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
