### 启动容器
docker-compose up -d

### 写入密码
如果没有htpasswd 工具需要进行安装
yum -y install httpd-tools

# 密码写入
htpasswd -Bbn 账号 密码 > auth/htpasswd

# 访问仓库
 http://ip:5000/v2/_catalog

# 在推送仓库之前我们需要给镜像打标签
docker tag docker.io/busybox  127.0.0.1:8001/test/busybox

这里推镜像必须的是 镜像仓库地址/:仓库名/:镜像名 格式的，但是127.xxx.xxx 这明显不太优雅。
我们可以通过nginx 去做反向代理。

# 这里需要强调一点的是，nginx 直接在宿主机进行安装，在docker进行安装映射80端口会出各种bug...
你可以直接 
yun install -y nginx

systemctl enable nginx

将registry.conf 移动到 /etc/nginx/conf.d 目录下。
或者 vim /etc/nginx/conf.d/registry.conf 

systemctl start nginx

然后你可以尝试curl registry.com 

到这里，我们还不能直接进行推送仓库
cat /etc/docker/daemon.json
{
"insecure-registries": ["registry.com"],
"log-driver":"json-file",
"log-opts": {"max-size":"500m", "max-file":"3"}
}

insecure-registries 这个字段里面需要加入我们刚才加入的域名。

# 这时候在推送仓库之前我们需要给镜像打标签
docker tag docker.io/busybox  registry.com/busybox



docker login -u 账号 -p 密码 仓库地址
