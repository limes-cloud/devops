package code_pack

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"service/tools"
	"service/tools/exec"
	"strconv"
	"strings"
	"time"
)

type pack struct {
	call             func(string)
	cmd              exec.Interface
	WorkDir          string            `json:"work_dir"`          // 工作目录
	GitUrl           string            `json:"git_url"`           // git代码地址
	GitToken         string            `json:"git_token"`         // git token
	RegistryProtocol string            `json:"registry_protocol"` // 镜像仓库协议
	RegistryUrl      string            `json:"registry_url"`      // 仓库地址
	RegistryUser     string            `json:"registry_user"`     // 仓库账号
	RegistryPass     string            `json:"registry_pass"`     // 仓库密码
	ServerName       string            `json:"server_name"`       // 服务名
	ServerBranch     string            `json:"server_branch"`     // 服务分支
	ServerVersion    string            `json:"server_version"`    // 服务版本
	Exec             string            `json:"exec"`              // 执行器
	Dockerfile       string            `json:"dockerfile"`        // 打包脚本
	Args             map[string]string `json:"args"`              // 打包脚本变量
	IsSsh            bool              `json:"is_ssh"`
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

func (p *pack) Call(str string) {
	if p.call != nil {
		p.call(str)
	}
}

func (p *pack) SetWatch(f func(string)) {
	p.call = f
}

func (p *pack) GetImageName() string {
	return fmt.Sprintf("%v/%v:%v", p.RegistryUrl, p.ServerName, p.ServerVersion)
}

func (p *pack) GetServerWorkDir() string {
	name, _ := p.GetGitName()
	return p.WorkDir + "/" + name
}

// HasDevImage 是否存在本地镜像
func (p *pack) HasDevImage() (res bool) {
	p.Call(fmt.Sprintf("判断是否存在本地镜像"))
	defer func() {
		p.Call(fmt.Sprintf("判断是否存在本地镜像 => %v", res))
	}()

	cmd := p.cmd.Command(p.Exec, "-c", fmt.Sprintf("docker images %v|wc -l", p.GetImageName()))
	out, err := cmd.Output()
	if err != nil {
		res = false
		return
	}
	reg := regexp.MustCompile(`[1-9]+`)
	outStr := reg.FindString(string(out))
	length, _ := strconv.ParseInt(outStr, 10, 64)
	res = length == 2
	return
}

// HasRemoteImage 是否存在远程镜像
func (p *pack) HasRemoteImage() (res bool, err error) {
	p.Call(fmt.Sprintf("判断是否存在远程镜像"))
	defer func() {
		if err != nil {
			p.Call(fmt.Sprintf("判断是否存在远程镜像失败 => %v", err.Error()))
		}
		p.Call(fmt.Sprintf("判断是否存在远程镜像 => %v", res))
	}()

	shell := fmt.Sprintf("curl -X GET -u %v:%v %v/v2/%v/tags/list", p.RegistryUser, p.RegistryPass, p.RegistryUrl, p.ServerName)
	cmd := p.cmd.Command(p.Exec, "-c", shell)
	out, err := cmd.Output()
	if err != nil {
		res = false
		err = errors.New("镜像仓库访问失败")
		return
	}

	resp := struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}{}

	_ = json.Unmarshal(out, &resp)
	if resp.Name == "" {
		return
	}
	res = tools.InList(resp.Tags, p.ServerVersion)
	return
}

// CreateWorkDir 创建工作目录
func (p *pack) CreateWorkDir() (err error) {
	p.Call(fmt.Sprintf("创建全局工作目录"))
	defer func() {
		if err != nil {
			p.Call(fmt.Sprintf("创建全局工作目录失败 => %v", err.Error()))
		}
	}()

	if !p.HasWorkDir() {
		if err = os.MkdirAll(p.WorkDir, os.ModePerm); err != nil {
			return
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
func (p *pack) GetGitVersion() (res string, err error) {
	p.Call("获取git版本")
	defer func() {
		if err != nil {
			p.Call(fmt.Sprintf("获取git版本失败 => %v", err.Error()))
		}
		p.Call(fmt.Sprintf("git版本 => %v", res))
	}()

	shell := "git version"

	p.Call(shell)

	cmd := p.cmd.Command(p.Exec, "-c", shell)
	version, err := cmd.Output()
	if err != nil {
		return
	}
	res = string(version)
	return
}

// GetDockerVersion 获取docker版本
func (p *pack) GetDockerVersion() (res string, err error) {
	p.Call(fmt.Sprintf("获取docker版本"))
	defer func() {
		if err != nil {
			p.Call(fmt.Sprintf("获取docker版本 => %v", err.Error()))
		}
		p.Call(fmt.Sprintf("docker版本 => %v", res))
	}()

	shell := "docker -v"

	p.Call(shell)

	cmd := p.cmd.Command(p.Exec, "-c", shell)
	version, err := cmd.Output()
	if err != nil {
		return
	}
	res = string(version)
	return
}

// GitCloneCode 进行代码拉去
func (p *pack) GitCloneCode() (err error) {
	p.Call(fmt.Sprintf("进行代码拉取"))
	defer func() {
		if err != nil {
			p.Call(fmt.Sprintf("进行代码拉取失败 => %v", err.Error()))
		}
	}()

	arr := strings.Split(p.GitUrl, "//")
	if len(arr) != 2 {
		return errors.New("git url error")
	}

	if p.GitToken != "" {
		p.GitUrl = fmt.Sprintf("%v//oauth2:%v@%v", arr[0], p.GitToken, arr[1])
	}

	// 拉去代码
	cmd := p.cmd.Command(p.Exec, "-c", fmt.Sprintf("git clone %v", p.GitUrl))
	cmd.SetDir(p.WorkDir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("git clone error :%v", string(out))
		return
	}

	// 切换分支
	cmd = p.cmd.Command(p.Exec, "-c", fmt.Sprintf("git checkout %v", p.ServerBranch))
	cmd.SetDir(p.GetServerWorkDir())
	out, err = cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("git checkout error:%v", string(out))
		return
	}

	return
}

func (p *pack) RemoveRemoteImage(name string, version string) {
	var err error
	defer func() {
		if err != nil {
			p.Call(fmt.Sprintf("清理镜像 %v/%v:%v 失败 => %v", p.RegistryUrl, name, version, err.Error()))
		}
	}()

	url := fmt.Sprintf("%v://%v/v2/%v/manifests/%v", p.RegistryProtocol, p.RegistryUrl, name, version)
	client := http.Client{Timeout: 10 * time.Second}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	if err != nil {
		return
	}
	request.SetBasicAuth(p.RegistryUser, p.RegistryPass)
	response, err := client.Do(request)
	if err != nil {
		return
	}
	sha := response.Header.Get("Docker-Content-Digest")
	if sha == "" {
		return
	}

	delUrl := fmt.Sprintf("%v://%v/v2/%v/manifests/%v", p.RegistryProtocol, p.RegistryUrl, name, sha)
	request, err = http.NewRequest(http.MethodDelete, delUrl, nil)
	if err != nil {
		return
	}
	request.SetBasicAuth(p.RegistryUser, p.RegistryPass)
	response, err = client.Do(request)
}

// Pack 进行打包镜像
func (p *pack) Pack() (err error) {

	p.Call(fmt.Sprintf("进行打包镜像"))
	defer func() {
		if err != nil {
			p.Call(fmt.Sprintf("进行打包镜像 => %v", err.Error()))
		}
	}()

	// 渲染dockerfile
	p.RenderDockerfile()
	// 创建dockerfile
	err = os.WriteFile(p.GetServerWorkDir()+"/Dockerfile", []byte(p.Dockerfile), os.ModePerm)
	if err != nil {
		err = fmt.Errorf("create dockerfile error :%v", err.Error())
		return
	}

	// 执行dockerfile
	cmd := p.cmd.Command(p.Exec, "-c", fmt.Sprintf("docker build -t %v .", p.GetImageName()))
	cmd.SetDir(p.GetServerWorkDir())
	if out, cErr := cmd.CombinedOutput(); cErr != nil {
		err = fmt.Errorf("docker build error :%v", string(out))
		return
	}

	return
}

func (p *pack) RenderDockerfile() {

	p.Call("进行docker模板渲染")
	defer func() {
		p.Call("进行docker模板渲染成功")
		p.Call(p.Dockerfile)
	}()

	reg := regexp.MustCompile(`\{\w+}`)
	list := reg.FindAllString(p.Dockerfile, -1)
	for _, val := range list {
		key := val[1 : len(val)-1]
		p.Dockerfile = strings.ReplaceAll(p.Dockerfile, val, p.Args[key])
	}
}

// Login 登陆仓库
func (p *pack) Login() (err error) {
	p.Call("进行docker镜像仓库登陆")
	defer func() {
		if err != nil {
			p.Call(fmt.Sprintf("进行docker镜像仓库登陆失败 => %v", err.Error()))
		}
	}()

	shell := fmt.Sprintf("docker login -u %v -p %v %v", p.RegistryUser, p.RegistryPass, p.RegistryUrl)
	cmd := p.cmd.Command(p.Exec, "-c", shell)
	if out, cErr := cmd.CombinedOutput(); cErr != nil {
		err = fmt.Errorf("docker login error :%v", string(out))
		return
	}
	return
}

// Upload 进行镜像上传
func (p *pack) Upload() (err error) {
	p.Call("进行镜像上传")
	defer func() {
		if err != nil {
			p.Call(fmt.Sprintf("进行镜像上传 => %v", err.Error()))
		}
	}()

	if !p.HasDevImage() {
		err = fmt.Errorf("docker image not exist")
		return
	}

	// 登陆仓库
	if err = p.Login(); err != nil {
		return
	}

	// 推送镜像
	cmd := p.cmd.Command(p.Exec, "-c", fmt.Sprintf("docker push %v", p.GetImageName()))
	if out, cErr := cmd.CombinedOutput(); cErr != nil {
		err = fmt.Errorf("docker push error :%v", string(out))
		return
	}
	return nil
}

func (p *pack) RemoveImage() {

	if !p.HasDevImage() {
		return
	}
	p.Call(fmt.Sprintf("删除本地镜像:%v", p.GetImageName()))
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
		p.Call("清理工作痕迹")
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

	p.Call("pack success")
	return nil
}
