### 前言
在安装devops管理平台之前我们先来说一下整个平台的技术架构，平台目前涉及的服务主要有 `用户中心`、`配置中心`、`服务中心`组成。用户中心主要功能包括平台用户管理、菜单管理、部门管理等，另外还对整个系统的api进行鉴权。配置中心主要功能包括配置业务字段管理、配置模板管理，配置中心依靠consul\etcd\zk等中间件可以实现配置的动态变更，对多环境开发中统一配置管理非常友好。服务管理主要负责系统的服务构建、服务发布、网络配置等，并且兼容k8s和docker-compose两种打包方式，你可以在单机节点就可以依靠docker-compose实现自动构建打包。
主要架构如下：
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672155753912-67527df2-36e3-4c67-a76b-9a0c427a07a6.png#averageHue=%23f3f3f3&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=560&id=u279455cb&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1120&originWidth=2354&originalType=binary&ratio=1&rotation=0&showTitle=false&size=253975&status=done&style=none&taskId=u26120705-db6a-4279-82c9-13ce2f96871&title=&width=1177)
### 服务安装要求
由上图，我们可以看到我们需要安装后台管理系统、Nginx、用户中心、配置中心、服务中心、mysql、redis、consul（etcd、zk）、k8s（docker-compose）服务。看到这里是不是你被吓住了，别害怕，跟着教程来，我们很快就可以实现安装的，但是在安装之前，我们应该**拥有一台配置不低于2h4g的服务器**(你至少要安装mysql和redis和docker)，但是在最好的情况下，我们是最少需要多台服务器的，一台作为构建服务，mysql、redis也应该用专用服务器，配置中心底层依赖consul（etcd、zk）也应该有专用服务器，最后运行服务的应该是一个服务集群。
假如你是个人开发者，并不满足这些条件，那么我想需要重新设计一下方案了，我们可以用一台作为构建服务（nginx、devops系统、consul（etcd、zk）、mysql、redis）一台作为运行服务（docker-compose）。假如你真的只有一台的情况下，也可以把构建服务和运行服务全都在一台上进行构建，只是我在后续的文档中会假设你拥有两台服务器作为例子进行讲解，但这不影响你只有一台服务器进行搭建，你可以将构建服务器就当成运行服务器即可。

**服务器配置要求**
**运行服务器**：2h4g+ centos系统
**构建服务器**：2h4g+ centos系统
**已备案域名：**实在没有的话，也可以通过ip进行访问。

接下来我们要一步一步让devops平台运行起来了。

特别提示：假如后续过程中使用git拉取代码失败，可以进行如下操作
```shell
vim /etc/hosts

# 新增如下两条记录
192.30.255.112 github.com
192.30.255.112 raw.githubusercontent.com

# 退出保存
```
### 构建服务安装
#### 一、nginx 安装
nginx 主要是用作后端服务的网关，提供域名前后端域名访问。
```shell
# 1:安装nginx
yum install nginx -y 

# 2:验证是否安装成功
nginx -V
# nginx version: nginx/1.20.1

# 3:设置nginx为守护开机自启
systemctl enable nginx

# 4:启动nginx
systemctl start nginx 
```
#### 二、Git安装
由于需要使用到git在后续的教程中进行devops代码拉取，所以我们先提前进行安装git
```shell
# 1:安装git
yum install git -y

# 2:验证是否安装成功
git --version
# git version 1.8.3.1
```
#### 三、docker安装
```shell
# 1:安装docker
curl -sSL https://get.daocloud.io/docker | sh

# 2:验证docker是否安装成功
docker --version
# Docker version 20.10.16, build aa7e414

# 3:设置nginx为守护开机自启
systemctl enable docker

# 4:启动nginx
systemctl start docker 
```
#### 四、安装docker-compose
```shell
# 安装docker-compose 
sudo curl -L https://get.daocloud.io/docker/compose/releases/download/1.25.1/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose

# 执行加权限
sudo chmod +x /usr/local/bin/docker-compose

# 验证是否安装完成
docker-compose --version
# docker-compose version 1.25.1, build a82fef07

```
#### 四、安装服务所需的基础设施（mysql、redis、consul）
```shell
# 创建devops工作目录
mkdir /usr/local/devops && cd /usr/local/devops

# 创建基础设施工作目录
mkdir drive && cd drive

# 创建docker-compose yaml资源清单
vim docker-compose.yaml

# 注意下，资源清单中的mysql的密码是需要改成你自己的密码的

```
```yaml
version: '3'
services:
  redis:
    image: redis:6.2.5
    container_name: redis
    privileged: true
    volumes:
      - ./redis/data:/data
      - ./redis/logs:/logs
    environment:
      - TZ="Asia/Shanghai"
  mysql:
    container_name: mysql
    image: mysql:5.7 
    volumes:
      - ./mysql/data/db:/var/lib/mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: 123456 #这里改成你的密码

networks:
  devops_net:
    driver: bridge
```
```shell
# 验证基础设施是否安装成功
docker ps
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672215047580-3880eed1-f57d-4f69-944c-4fe0f2252edb.png#averageHue=%23141728&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=92&id=ufdbd8db5&margin=%5Bobject%20Object%5D&name=image.png&originHeight=184&originWidth=2340&originalType=binary&ratio=1&rotation=0&showTitle=false&size=77686&status=done&style=none&taskId=ucf7298a0-3cd9-40d6-bf43-a2cbe12d3dc&title=&width=1170)
如上 存在redis、mysql、consul 及代表基础设施服务已经安装成功。

#### 五、下载服务端代码
```yaml
# 创建服务端目录
mkdir /usr/local/devops/service && cd /usr/local/devops/service 

# 拉去服务端代码
git clone https://github.com/limeschool/devops.git 

```
#### 六、导入数据库资源
```yaml
# 复制sql文件到数据库的挂载目录
cp devops/devops.sql ../drive/mysql/data/db/devops.sql

# 进入mysql容器
docker exec -it mysql /bin/sh

#连接mysql
mysql -u root -p
# Enter password: 
# 接下来输入密码，密码是在创建资源清单时所设置的
# mysql> 

# 导入文件
source /var/lib/mysql/devops.sql

# 验证文件导入
show databases;

# 输出一下devops_configure、devops_service、devops_ums 即为成功
# +--------------------+
# | Database           |
# +--------------------+
# | information_schema |
# | devops_configure   |
# | devops_service     |
# | devops_ums         |
# | mysql              |
# | performance_schema |
# | sys                |
# +--------------------+
# 7 rows in set (0.00 sec)

# 退出mysql
exit

# 退出docker
exit
```

#### 七、启动devops构建后台服务 devops_config、devops_service、devops_ums
##### 创建ums服务启动配置
```shell
# 创建目录
mkdir -p config/ums

# 创建配置文件
vim config/ums/conf.json

# 输入以下配置
# 首先我们需要进行一下配置进行修改
# mysql.dsn 中需要修改成自己的账号密码
# jwt.secret 中改成自己的加密密钥
```
```json
{
  "service": "user-center",
  "system": {
    "skip_request_log": {
      "get:/auth": true,
      "get:/ums/menu": true,
      "get:/ums/role/menu": true,
      "get:/ums/team": true
    }
  },
  "log": {
    "level": 0,
    "trace_key": "log-id",
    "service_key": "service"
  },
  "request":{
    "enable_log": true,
    "retry_count": 3,
    "retry_wait_time":"1s",
    "timeout": "10s",
    "request_msg": "http request",
    "response_msg": "http response"
  },
  "mysql": [
    {
      "enable": true,
      "name": "devops",
      "dsn": "root:123456@tcp(mysql:3306)/devops_ums?charset=utf8mb4&parseTime=True&loc=Local"
    }
  ],
  "redis": [
    {
      "enable": true,
      "name": "redis",
      "host": "redis:6379"
    }
  ],
  "jwt": {
    "expire": 7200,
    "max_expire": 172800,
    "secret": "limeschool"
  },
  "whitelist": {
    "post:/ums/user/login": true,
    "post:/ums/token/refresh": true,
    "get:/configure/config": true,
    "get:/ums/api/v1/user": true
  },
  "login_limit": {
    "enable":true,
    "ip_limit": 10,
    "password_error_limit": 3
  },
  "rsa": [
    {
      "enable": true,
      "name": "private",
      "path": "cert/private.pem"
    },
    {
      "enable": true,
      "name": "public",
      "path": "cert/public.pem"
    }
  ],
  "rbac": {
    "enable": true,
    "db": "devops"
  }
}

```
##### 创建service服务启动配置
```shell
# 创建目录
mkdir -p config/service

# 创建配置文件
vim config/service/conf.json

# 输入以下配置
# 首先我们需要进行一下配置进行修改
# mysql.dsn 中需要修改成自己的账号密码
# work_dir  docker构建的工作目录 /usr/local/docker
# exec_type 执行shell命令的方式 /bin/sh
# release_type 发布方式 可选["docker-compose","k8s"]
# ums_addr ums服务的请求的地址
```
```json
{
  "service": "service",
  "system": {
    "skip_request_log": {
      "get:/auth": true
    },
    "client_timeout":"10s"
  },
  "log": {
    "level": 0,
    "trace_key": "log-id",
    "service_key": "service"
  },
  "request":{
    "enable_log": true,
    "retry_count": 3,
    "retry_wait_time":"1s",
    "timeout": "120s",
    "request_msg": "http request",
    "response_msg": "http response"
  },
  "mysql": [
    {
      "enable": true,
      "name": "devops",
      "dsn": "root:123456@tcp(mysql:3306)/devops_service?charset=utf8mb4&parseTime=True&loc=Local"
    }
  ],
  "redis": [
    {
      "enable": true,
      "name": "redis",
      "host": "redis:6379"
    }
  ],
  "ums_addr": "http://ums:8080",
  "work_dir":  "/usr/local/docker",
  "exec_type": "/bin/sh",
  "release_type": ["docker-compose"]
}

```
##### 创建configure服务启动配置
```shell
# 创建目录
mkdir -p config/configure

# 创建配置文件
vim config/configure/conf.json

# 输入以下配置
# 首先我们需要进行一下配置进行修改
# mysql.dsn 中需要修改成自己的账号密码
# jwt.secret 中改成自己的加密密钥
```
```json
{
  "service": "configure",
  "system": {
    "skip_request_log": {
      "get:/auth": true,
      "get:/configure/config": true
    },
    "timeout": "20s"
  },
  "log": {
    "level": 0,
    "trace_key": "log-id",
    "service_key": "service"
  },
  "request":{
    "enable_log": true,
    "retry_count": 3,
    "retry_wait_time":"1s",
    "timeout": "10s",
    "request_msg": "http request",
    "response_msg": "http response"
  },
  "mysql": [
    {
      "enable": true,
      "name": "devops",
      "dsn": "root:123456@tcp(mysql:3306)/devops_configure?charset=utf8mb4&parseTime=True&loc=Local"
    }
  ],
  "redis": [
    {
      "enable": true,
      "name": "redis",
      "host": "redis:6379"
    }
  ]
}
```

#### 创建启动docker-compose资源清单
```yaml
version: '3'
services:
  ums:
    container_name: ums
    build:
      context: ./devops/ums/
      dockerfile: Dockerfile
    image: ums:v1.0
    restart: always
    ports:
      - 8080:8080
    volumes:
      - ./config/ums:/go/build/config
    environment:
      TZ: Asia/Shanghai
    networks:
      - devops_net

  service:
    container_name: service
    build:
      context: ./devops/service/
      dockerfile: Dockerfile
    image: service:v1.0
    restart: always
    ports:
      - 8081:8081
    volumes:
      - ./config/service:/go/build/config
    environment:
      TZ: Asia/Shanghai
    networks:
      - devops_net

  configure:
    container_name: configure
    build:
      context: ./devops/configure/
      dockerfile: Dockerfile
    image: configure:v1.0
    restart: always
    ports:
      - 8082:8082
    volumes:
      - ./config/configure:/go/build/config
    environment:
      TZ: Asia/Shanghai
    networks:
      - devops_net

networks:
  devops_net:
     external:
        name: drive_devops_net
```
使用` docker ps` 验证是否安装成功，安装成功会得到如下的输出，这样以来我们服务端的代码就相当于安装成功了。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672324471246-34cb05d1-26b4-47e5-beae-2d2dd88795db.png#averageHue=%23141728&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=167&id=u6ff332af&margin=%5Bobject%20Object%5D&name=image.png&originHeight=334&originWidth=2068&originalType=binary&ratio=1&rotation=0&showTitle=false&size=131603&status=done&style=none&taskId=u3555dd87-87f0-49db-adab-10825d2b724&title=&width=1034)

#### 八、为后端服务绑定域名
##### 添加域名
这里我们需要使用到域名服务，这里我会以腾讯云为例进行域名的解析与绑定。进入腾讯云-域名管理平台，点击添加记录。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672325294281-00e35387-95bf-4d57-8d40-e6ee3a8a7379.png#averageHue=%23fbe8ce&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=530&id=u6f866569&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1060&originWidth=2228&originalType=binary&ratio=1&rotation=0&showTitle=false&size=378164&status=done&style=none&taskId=ub87192ea-db97-4c02-8627-00672baf263&title=&width=1114)
进行域名解析，如下我解析的域名前缀为devops，假如我的主站为qlime.cn，则我的后台服务的域名则为devops.qlime.cn 这里的记录值就是我们安装后台服务的服务器的ip。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672325423221-7fbdfd46-314f-4b8a-b2d8-faa80772b0cc.png#averageHue=%23fdfdfb&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=278&id=u7c6659e3&margin=%5Bobject%20Object%5D&name=image.png&originHeight=556&originWidth=2084&originalType=binary&ratio=1&rotation=0&showTitle=false&size=140687&status=done&style=none&taskId=u70706e87-28fd-4b74-94f6-5c482b0450e&title=&width=1042)
点击确定了之后保存即可。
##### 申请ssl证书
接下来我们进入证书申请[腾讯云-证书免费申请](https://console.cloud.tencent.com/ssl)（假如你需要使用https协议，则需要继续进行证书申请，如果你不需要，那后续的证书申请流程你可以直接选择跳过）
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672325756084-1d02f4f8-e6ab-4e64-ac63-9eb2efc44333.png#averageHue=%237bc387&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=550&id=ufceba336&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1100&originWidth=2606&originalType=binary&ratio=1&rotation=0&showTitle=false&size=386282&status=done&style=none&taskId=u768ae837-0607-43c7-b78d-be1a9ed452c&title=&width=1303)
进入之后进行如下的填写,填写完成之后直接提交即可。审批需要一点时间，一般10分钟左右。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672325953328-e9e8255a-c7d1-4724-bd7a-995921a40828.png#averageHue=%23fcfaf9&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=638&id=ud2164c0b&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1276&originWidth=1672&originalType=binary&ratio=1&rotation=0&showTitle=false&size=262859&status=done&style=none&taskId=u2d8a09f0-12ec-4dd7-97e8-8d3d97424f9&title=&width=836)
回到证书列表，对刚才已经颁发的证书文件进行下载（选择Nginx）。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672326091147-29715f59-a256-45eb-a60a-5d213f839de3.png#averageHue=%2388cb94&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=479&id=u3552130c&margin=%5Bobject%20Object%5D&name=image.png&originHeight=958&originWidth=2838&originalType=binary&ratio=1&rotation=0&showTitle=false&size=307788&status=done&style=none&taskId=u0e566acf-18e3-4e57-967e-9331283381f&title=&width=1419)

接下来我们需要在我们服务器上对域名进行配置了。还记得在第一步安装的nginx么。现在起就是它的发挥作用的时候了。
##### 进行nginx配置
```shell
# 进入nginx配置目录
cd /etc/nginx/conf.d/


# 创建ssl证书 [不需要则跳过]
mkdir -p ssl/devops
# 创建ssl.key
vim ssl/devops/ssl.key
# 这里打开刚才下载好的证书文件，将后缀为.key的文件的内容粘贴进去保存即可。
# 创建ssl.crt
vim ssl/devops/ssl.crt
# 这里打开刚才下载好的证书文件，将后缀为.crt的文件的内容粘贴进去保存即可。
# 创建证书结束

# 创建配置文件
vim devops.conf

# 输入配置文件内容如下
# 特别提示一下内容需要修改成自己的
server_name  xxx;
```
```shell
server {
    listen        80;
    server_name  devops.qlime.cn;
        
		
		location /api/auth {
			fastcgi_intercept_errors on;
			proxy_pass_request_body  off;
			if ($request_method = 'OPTIONS') {
			    return 204;
			}
			proxy_set_header X-Original-URI $request_uri;
			proxy_set_header X-Original-METHOD $request_method;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
			proxy_set_header X-Forwarded-Proto $scheme;
			proxy_pass http://127.0.0.1:8080/auth;
		}
		
		location /ums {
			auth_request /api/auth;
			auth_request_set $user $upstream_http_x_user;
			auth_request_set $error $upstream_http_x_error;
			auth_request_set $error_code $upstream_http_x_error_code;
			proxy_set_header X-Error $error;
			proxy_set_header X-User $user;
			error_page 400 401 403 404 501 500 502 503 504 = @error;
			
			proxy_set_header Host $http_host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header REMOTE-HOST $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		  proxy_pass http://127.0.0.1:8080;
		}
		
		location /service {
			auth_request /api/auth;
			auth_request_set $user $upstream_http_x_user;
			error_page 400 401 403 404 501 500 502 503 504 = @error;
			proxy_set_header X-User $user;
			proxy_set_header Host $http_host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header REMOTE-HOST $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		  proxy_pass http://127.0.0.1:8081;
		}
		
		location /configure {
			auth_request /api/auth;
			auth_request_set $user $upstream_http_x_user;
			error_page 400 401 403 404 501 500 502 503 504 = @error;
			proxy_set_header X-User $user;
			proxy_set_header Host $http_host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header REMOTE-HOST $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		  proxy_pass http://127.0.0.1:8082;
		}

		location @error{
			add_header Access-Control-Allow-Origin *;
			add_header Content-Type application/json;
			if ($error){
				return 200 '{"code":$error_code,"msg":"$error"}';
			}
			return 200 '{"code":400,"msg":"系统内部错误"}';
		}
}
```
```shell
server
{
    listen 80;
    listen 443 ssl http2;
    server_name devops.qlime.cn;
    if ($server_port !~ 443){
        rewrite ^(/.*)$ https://$host$1 permanent;
    }
    ssl_certificate    /etc/nginx/conf.d/ssl/devops/ssl.crt;
    ssl_certificate_key    /etc/nginx/conf.d/ssl/devops/ssl.key;
    ssl_protocols TLSv1.1 TLSv1.2 TLSv1.3;
    ssl_ciphers EECDH+CHACHA20:EECDH+CHACHA20-draft:EECDH+AES128:RSA+AES128:EECDH+AES256:RSA+AES256:EECDH+3DES:RSA+3DES:!MD5;
    ssl_prefer_server_ciphers on;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;
    add_header Strict-Transport-Security "max-age=31536000";


    location /api/auth {
			fastcgi_intercept_errors on;
			proxy_pass_request_body  off;
			if ($request_method = 'OPTIONS') {
			    return 204;
			}
			proxy_set_header X-Original-URI $request_uri;
			proxy_set_header X-Original-METHOD $request_method;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
			proxy_set_header X-Forwarded-Proto $scheme;
			proxy_pass http://127.0.0.1:8080/auth;
		}
		
		location /ums {
			auth_request /api/auth;
			auth_request_set $user $upstream_http_x_user;
			auth_request_set $error $upstream_http_x_error;
			auth_request_set $error_code $upstream_http_x_error_code;
			proxy_set_header X-Error $error;
			proxy_set_header X-User $user;
			error_page 400 401 403 404 501 500 502 503 504 = @error;
			
			proxy_set_header Host $http_host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header REMOTE-HOST $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		  proxy_pass http://127.0.0.1:8080;
		}
		
		location /service {
			auth_request /api/auth;
			auth_request_set $user $upstream_http_x_user;
			error_page 400 401 403 404 501 500 502 503 504 = @error;
			proxy_set_header X-User $user;
			proxy_set_header Host $http_host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header REMOTE-HOST $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		  proxy_pass http://127.0.0.1:8081;
		}
		
		location /configure {
			auth_request /api/auth;
			auth_request_set $user $upstream_http_x_user;
			error_page 400 401 403 404 501 500 502 503 504 = @error;
			proxy_set_header X-User $user;
			proxy_set_header Host $http_host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header REMOTE-HOST $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		  proxy_pass http://127.0.0.1:8082;
		}

		location @error{
			add_header Access-Control-Allow-Origin *;
			add_header Content-Type application/json;
			if ($error){
				return 200 '{"code":$error_code,"msg":"$error"}';
			}
			return 200 '{"code":400,"msg":"系统内部错误"}';
		}
}


```
使用命令 `nginx -s reload` 重新载入nginx配置，接下来在浏览器中直接访问配置好的url,出现如下即为安装成功。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672327193508-1b73ba17-cec1-4798-8a73-e62eba0578f6.png#averageHue=%23f4f4f4&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=457&id=u882e7958&margin=%5Bobject%20Object%5D&name=image.png&originHeight=914&originWidth=2726&originalType=binary&ratio=1&rotation=0&showTitle=false&size=431942&status=done&style=none&taskId=u9d27291b-01f7-42e4-b57a-b18145e7fa9&title=&width=1363)

#### 安装后台web端
```shell
# 下载后台web端代码
# 创建web目录
mkdir /usr/local/devops/service-web && cd /usr/local/devops/service-web
# 下载源码
git clone https://github.com/limeschool/devops-admin.git
#修改配置
vim devops-admin/dist/config.js

# baseUrl 需要改成你所配置的后台服务域名

# var SYSTEM_CONFIG = {
#    baseUrl : "https://devops.qlime.cn"
# }
```
为后台web端配置url，这里参考前面的搭建后端服务申请域名的方式进行搭建即可，如果你的后台服务申请了证书，那你在配置web端的时候也应该要申请证书。这里我申请的域名为`devops-admin.qlime.cn`
```shell
# 进入nginx配置目录
cd /etc/nginx/conf.d/


# 创建ssl证书 [不需要则跳过]
mkdir -p ssl/devops-admin
# 创建ssl.key
vim ssl/devops-admin/ssl.key
# 这里打开刚才下载好的证书文件，将后缀为.key的文件的内容粘贴进去保存即可。
# 创建ssl.crt
vim ssl/devops-admin/ssl.crt
# 这里打开刚才下载好的证书文件，将后缀为.crt的文件的内容粘贴进去保存即可。
# 创建证书结束

# 创建配置文件
vim devops-admin.conf

# 输入配置文件内容如下
# 特别提示一下内容需要修改成自己的
# server_name  xxx;
```

```shell
server
{
    listen 80;
    server_name devops-admin.qlime.cn;
    root /usr/local/devops/service-web/devops-admin/dist;

    #禁止访问的文件或目录
    location ~ ^/(\.user.ini|\.htaccess|\.git|\.env|\.svn|\.project|LICENSE|README.md)
    {
        return 404;
    }

    location ~ .*\.(gif|jpg|jpeg|png|bmp|swf)$
    {
        expires      30d;
        error_log /dev/null;
        access_log /dev/null;
    }
}


```
```shell
server
{
    listen 80;
    listen 443 ssl http2;
    server_name devops-admin.qlime.cn;
    root /usr/local/devops/service-web/devops-admin/dist;
    if ($server_port !~ 443){
        rewrite ^(/.*)$ https://$host$1 permanent;
    }
    ssl_certificate    /etc/nginx/conf.d/ssl/devops-admin/ssl.crt;
    ssl_certificate_key    /etc/nginx/conf.d/ssl/devops-admin/ssl.key;
    ssl_protocols TLSv1.1 TLSv1.2 TLSv1.3;
    ssl_ciphers EECDH+CHACHA20:EECDH+CHACHA20-draft:EECDH+AES128:RSA+AES128:EECDH+AES256:RSA+AES256:EECDH+3DES:RSA+3DES:!MD5;
    ssl_prefer_server_ciphers on;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;
    add_header Strict-Transport-Security "max-age=31536000";

    #禁止访问的文件或目录
    location ~ ^/(\.user.ini|\.htaccess|\.git|\.env|\.svn|\.project|LICENSE|README.md)
    {
        return 404;
    }

    location ~ .*\.(gif|jpg|jpeg|png|bmp|swf)$
    {
        expires      30d;
        error_log /dev/null;
        access_log /dev/null;
    }
}


```

使用命令 `nginx -s reload` 重新载入nginx配置，接下来在浏览器中直接访问配置好的url,出现如下即为安装成功，到此为止我们的基本平台就搭建好了。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672328864730-b5827114-2b01-46f1-889f-bf51648576f1.png#averageHue=%23f4f6eb&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=823&id=uca75dbdc&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1646&originWidth=2836&originalType=binary&ratio=1&rotation=0&showTitle=false&size=943998&status=done&style=none&taskId=uc31f21a8-74fb-478f-830b-e8d8491cb0f&title=&width=1418)

### 服务中心安装
##### 发布流程简介
进入管理平台之后，找到服务中心菜单，我可以看到如下的菜单选项。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672417263422-233fa55d-872f-4e97-93e4-47e66795a5eb.png#averageHue=%23fefefe&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=598&id=u8fa7afa2&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1196&originWidth=2622&originalType=binary&ratio=1&rotation=0&showTitle=false&size=184769&status=done&style=none&taskId=ub0e7bddf-0967-4cf3-8840-14a18b333fa&title=&width=1311)
先大概讲一下，进行代码发布的流程。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672417794329-2ba69a2e-ee39-4185-a1ed-d237c86932cd.png#averageHue=%23f8f6f4&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=253&id=u290fa4fa&margin=%5Bobject%20Object%5D&name=image.png&originHeight=506&originWidth=1832&originalType=binary&ratio=1&rotation=0&showTitle=false&size=69799&status=done&style=none&taskId=uea9fa891-01cb-4529-b7d5-465e53c6d29&title=&width=916)
主要分为以下几个步骤，括号后面是具体的配置位置。
1：将代码发布到代码仓库中（代码仓库）
2：使用dockerfile 对代码进行打包（dockerfile管理）
3：将打包的docker镜像推送到镜像仓库（镜像仓库）
4：使用资源清单定于服务运行策略。（发布管理）
5：指定的机器使用 k8s/docker-compose 通过资源清单运行服务。（环境管理）
接下来我将带你进行一步一步配置
##### 代码仓库
代码仓库目前支持三种仓库github、gitlab、gitee。这里我们使用gitee作为示范进行讲解，其他的都操作基本都是一样的。
1）登陆gitee之后点击设置
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672418346517-c049272a-a59e-4cd7-95a9-3c4d7bac7c80.png#averageHue=%23ea8440&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=358&id=u992ebfa0&margin=%5Bobject%20Object%5D&name=image.png&originHeight=716&originWidth=2490&originalType=binary&ratio=1&rotation=0&showTitle=false&size=217089&status=done&style=none&taskId=u56feba4a-0f0d-43ae-bb67-06bc52564f9&title=&width=1245)
2）找到安全设置-》私人令牌
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672418411532-d8eed318-95ec-423f-99fd-20ef0e98991e.png#averageHue=%23fcfbfa&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=353&id=u07341303&margin=%5Bobject%20Object%5D&name=image.png&originHeight=706&originWidth=2408&originalType=binary&ratio=1&rotation=0&showTitle=false&size=197076&status=done&style=none&taskId=u41b1ffe5-4ce6-4195-9092-8117e6763ac&title=&width=1204)
3）点击新建令牌
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672418447300-8aa601db-943a-4519-84f5-ea8270f7392f.png#averageHue=%23fcfaf9&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=243&id=u189b82f9&margin=%5Bobject%20Object%5D&name=image.png&originHeight=486&originWidth=2040&originalType=binary&ratio=1&rotation=0&showTitle=false&size=78851&status=done&style=none&taskId=u036f3fe5-5b10-4732-b6c0-a0dab055693&title=&width=1020)
4）提交进行生成令牌
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672418511027-4042bb1f-0621-4afd-8361-b7759ca1cfa9.png#averageHue=%23faf8f6&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=454&id=uf33d8786&margin=%5Bobject%20Object%5D&name=image.png&originHeight=908&originWidth=2590&originalType=binary&ratio=1&rotation=0&showTitle=false&size=259916&status=done&style=none&taskId=u62d2ad2f-bad7-4252-8e00-ecae3969801&title=&width=1295)
5）使用token在devops管理平台配置代码仓库。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672418951661-43f91488-838a-40cb-8d83-a6a7fecd4de1.png#averageHue=%23a9a7a7&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=471&id=u1852b87a&margin=%5Bobject%20Object%5D&name=image.png&originHeight=942&originWidth=2598&originalType=binary&ratio=1&rotation=0&showTitle=false&size=220358&status=done&style=none&taskId=u03615efd-3904-481b-b81e-6d0f5cc9bc2&title=&width=1299)
6）进行代码仓库连接测试，提示成功则为配置成功。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672419026113-e270d450-cd42-40df-83ed-9b7af44469d0.png#averageHue=%23fefcfb&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=238&id=u77aa7527&margin=%5Bobject%20Object%5D&name=image.png&originHeight=476&originWidth=2200&originalType=binary&ratio=1&rotation=0&showTitle=false&size=96159&status=done&style=none&taskId=u2eae635c-e1e2-4fe4-8e28-f18d2e43ec6&title=&width=1100)

##### 镜像仓库配置
镜像仓库我们可以选择阿里云，腾讯云进行托管，由于小编对镜像的域名有洁癖，我希望我的镜像仓库的地址是我喜欢的，所以这里我选择自己搭建私有镜像仓库，其实这个过程一点也不复杂。这里我期望的镜像仓库地址为`docker-hub.qlime.cn` ，你们需要改成自己的域名。
```yaml
# 创建私有仓库工作目录
mkdir -p /usr/local/devops/registry && cd /usr/local/devops/registry

# 进行htpasswd工具安装
yum -y install httpd-tools

# 创建密钥迷离
mkdir auth

# 创建私有镜像仓库的账号密码,这里的username 和 password 需要替换成你自己的账号和密码
# htpasswd -Bbn username password > auth/htpasswd
htpasswd -Bbn xxxx xxxx > auth/htpasswd

# 进入阿里云对象存储创建存储桶，使用阿里云对象存储可以介于构建主机的硬盘成本，具体操作百度即可。

vim docker-compose.yaml
# 使用docker-compose 安装镜像，以下有几个地方需要改成自己的。
# - REGISTRY_STORAGE_OSS_ACCESSKEYID=xxx key
# - REGISTRY_STORAGE_OSS_ACCESSKEYSECRET=xxx 密钥
# - REGISTRY_STORAGE_OSS_REGION=xxx 所在域
# - REGISTRY_STORAGE_OSS_BUCKET=xxx 存储桶 

# 启动创建
docker-compose up -d
```
使用docker-compose 在构建服务上搭建私有仓库
```yaml
version: '3'
services:
  registry:
    restart: always
    image: "registry:latest"
    container_name: registry
    ports:
      - 8001:5000
    environment:
      - REGISTRY_AUTH=htpasswd
      - REGISTRY_AUTH_HTPASSWD_REALM=Registry Realm
      - REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd
      - REGISTRY_STORAGE=oss
      - REGISTRY_STORAGE_OSS_ACCESSKEYID=xxx
      - REGISTRY_STORAGE_OSS_ACCESSKEYSECRET=xxx
      - REGISTRY_STORAGE_OSS_REGION=xxx
      - REGISTRY_STORAGE_OSS_BUCKET=xxx
      - REGISTRY_STORAGE_OSS_INTERNAL=false 
      - REGISTRY_STORAGE_OSS_SECURE=false
      - REGISTRY_STORAGE_DELETE_ENABLED=true
    volumes:
      - ./auth:/auth
```
验证是否安装成功
```yaml
curl 127.0.0.1:8001/v2/_catalog
```
出现如下输出则为安装成功。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672420981885-e014d41a-45e5-4e82-b8dd-144ee771b433.png#averageHue=%23141728&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=41&id=u9aa70114&margin=%5Bobject%20Object%5D&name=image.png&originHeight=82&originWidth=1872&originalType=binary&ratio=1&rotation=0&showTitle=false&size=38053&status=done&style=none&taskId=u64adc180-5864-4542-a5ed-06d1c6e9fe7&title=&width=936)
接下来为这个镜像仓库配置域名，这里参考之前服务搭建配置域名的步骤，具体的步骤我就不再细说了。

```yaml
# 进入nginx配置目录
cd /etc/nginx/conf.d/


# 创建ssl证书 [不需要则跳过]
mkdir -p ssl/registry
# 创建ssl.key
vim ssl/registry/ssl.key
# 这里打开刚才下载好的证书文件，将后缀为.key的文件的内容粘贴进去保存即可。
# 创建ssl.crt
vim ssl/registry/ssl.crt
# 这里打开刚才下载好的证书文件，将后缀为.crt的文件的内容粘贴进去保存即可。
# 创建证书结束

# 创建配置文件
vim registry.conf

# 输入配置文件内容如下
# 特别提示一下内容需要修改成自己的
# server_name  xxx;
```
```shell
server
{
    listen 80;
    server_name docker-hub.qlime.cn;

    location / {
			proxy_pass http://127.0.0.1:8001;
		}
}


```
```shell
server
{
    listen 80;
    listen 443 ssl http2;
    server_name docker-hub.qlime.cn;
    if ($server_port !~ 443){
        rewrite ^(/.*)$ https://$host$1 permanent;
    }
    ssl_certificate    /etc/nginx/conf.d/ssl/registry/ssl.crt;
    ssl_certificate_key    /etc/nginx/conf.d/ssl/registry/ssl.key;
    ssl_protocols TLSv1.1 TLSv1.2 TLSv1.3;
    ssl_ciphers EECDH+CHACHA20:EECDH+CHACHA20-draft:EECDH+AES128:RSA+AES128:EECDH+AES256:RSA+AES256:EECDH+3DES:RSA+3DES:!MD5;
    ssl_prefer_server_ciphers on;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;
    add_header Strict-Transport-Security "max-age=31536000";

    location / {
			proxy_pass http://127.0.0.1:8001;
		}
}


```
使用命令 `nginx -s reload` 重新载入nginx配置，接下来在浏览器中直接访问`[https://docker-hub.qlime.cn/v2/_catalog](https://docker-hub.qlime.cn/v2/_catalog)`,出现如下即为安装成功，到此为止我们的基本平台就搭建好了。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672421791203-9e99e2b4-47e0-4045-8531-d4f91547729c.png#averageHue=%23c9c9c9&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=360&id=u2200cb95&margin=%5Bobject%20Object%5D&name=image.png&originHeight=720&originWidth=1186&originalType=binary&ratio=1&rotation=0&showTitle=false&size=71557&status=done&style=none&taskId=u02b1c41c-fd03-43f9-a753-1ed777697d3&title=&width=593)
接下来回到管理平台进行镜像仓库配置
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672421949068-1c3e1f7e-49c0-428b-a917-cec81548c50d.png#averageHue=%23f18580&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=492&id=u4aaf02b7&margin=%5Bobject%20Object%5D&name=image.png&originHeight=984&originWidth=2228&originalType=binary&ratio=1&rotation=0&showTitle=false&size=188667&status=done&style=none&taskId=udbd39a42-3a40-498f-81c3-92a78ab022d&title=&width=1114)
这里的记录数量是设置构建的时候同一个服务最多存储多少个副本，以节约存储空间。然后点击链接测试进行连通性测试，出现如下提示则为安装成功了。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672422061075-13d817a7-de6f-4ae4-85ca-d7741460db19.png#averageHue=%23fefdfd&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=253&id=uf72a6dd9&margin=%5Bobject%20Object%5D&name=image.png&originHeight=506&originWidth=2134&originalType=binary&ratio=1&rotation=0&showTitle=false&size=115604&status=done&style=none&taskId=uf0f3e84b-0190-4ebd-8838-b1e7441dbe9&title=&width=1067)
##### dockerfile 管理
dockerfile管理是一个管理镜像打包的dockerfile模板库，我们可以提前约定一些dockerfile模板，针对同一类的应用进行打包。比如golang的一类、java的一类、nodejs的一类，现在问题就来了，怎么才能做到同一类的打包呢，每个服务都有不同的端口，副本数量，怎么进行统一呢，这时候模板变量起到了关键性的作用，他把关键性的数据抽象成一个变量，针对与不用的服务进行不同的设置，比如下面这个dockerfile模板
```yaml
FROM golang:alpine AS build
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/build
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o entry main.go
FROM alpine
EXPOSE {ListenPort}
WORKDIR /go/build
COPY --from=build /go/build/entry /go/build/entry
ADD ./config /go/build/config
CMD ["./entry"]
```
则是一个golang的具有多阶段缓存的打包模板，在`EXPOSE {ListenPort}`里把暴露的端口变成了一个变量，根据服务管理里面的配置在打包的时候会渲染成服务设置的特定值。
这样的变量目前在devops中主要有如下几个
```
RunPort:        "运行端口"
ListenPort:     "监听端口"
Replicas:       "副本数量"
ProbeValue:     "探测值"
ProbeInitDelay: "探测延迟时间"
ServiceName:    "服务标志"
Image:          "发布镜像"
Namespace:      "所属环境"
```
##### 发布管理
发布管理主要是管理资源清单，用于k8s/docker-compose在服务部署的时候进行使用，发布管理和dockerfile管理同样具有模板变量，具体的变量值和dockerfile中的保持一致的。我们来看看以下这个docker-compose资源清单。
```yaml
version: '3.6'
services:
  {ServiceName}:
    container_name: {ServiceName}
    image: {Image}
    restart: always
    ports:
      - {RunPort}:{ListenPort}

```
在这个资源清单中，我们把服务名以及映射的端口号、镜像都做成了变量，假如我的服务的参数如下
```yaml
ServiceName: test
Image: nginx
RunPort: 8080
ListenPort: 80
```
当这个服务要求的运行副本只有一个的时候，这个docker-compose的资源清单则会被渲染成如下
```yaml
version: '3.6'
services:
  test:
    container_name: test
    image: nginx
    restart: always
    ports:
      - 8080:80

```
当这个服务要求运行多个副本的时候（假如说现在有3个），这个docker-compose的资源清单则会被渲染成如下
```yaml
version: '3.6'
services:
  test-1:
    container_name: test-1
    image: nginx
    restart: always
    ports:
      - 8080:80
  test-2:
    container_name: test-2
    image: nginx
    restart: always
    ports:
      - 8081:80
  test-3:
    container_name: test-3
    image: nginx
    restart: always
    ports:
      - 8082:80
```
在k8s中可以通过service直接进行负载均衡，那在docker-compose中多副本的情况在，我们又该如何进行负载均衡呢，别担心，这一点在后面的网络管理中会讲到。
##### 环境管理
在我们的企业开发中，一般都会具有，测试环境、预发布环境、生产环境等。这里的环境设置就是把不同的环境进行划分，这里我们做一个约束，在使用docker-compose的情况下，每一台主机就是一个单独的环境，在使用k8s的情况下，每一个命名空间就是一个单独的环境。这里有人就问了，我想在使用docker-compose的情况下单机区别多环境怎么办呢，其实也是可以的，但是由于我们的docker-compose最终是做端口映射到宿主机上的，只要我们能够对映射的端口进行区分就可以了的，比如A服务，我们可以把A服务的测试环境端口映射到18008，生产环境的映射到8008，这样也是可以的。
废话不多说了，来看一下怎么进行环境配置的。
1）点击新增按钮进行添加环境，如下图：
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672456676220-007dd2d3-f1af-4702-927f-fa7a212802d3.png#averageHue=%23b5b3b3&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=555&id=ub25b797d&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1110&originWidth=2608&originalType=binary&ratio=1&rotation=0&showTitle=false&size=219540&status=done&style=none&taskId=uc8a2a9fb-43b2-439d-b555-3473dd79d66&title=&width=1304)
这里需要注意一点的就是运行模式这个选项，默认是docker-compose,当然我们也可以开启k8s，以及docker-compose和k8s都打开，但是小编建议只开其中一种就可以了，以免造成资源混用，具体的开启方式就在我们搭建后台服务service的配置中有一个release_type。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/26359342/1672456957044-70b9922e-f5be-44c4-b4d7-03c5c02d5365.png#averageHue=%23f9f7f6&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=148&id=u1afcef5f&margin=%5Bobject%20Object%5D&name=image.png&originHeight=296&originWidth=1068&originalType=binary&ratio=1&rotation=0&showTitle=false&size=51417&status=done&style=none&taskId=u29d5cc70-ae64-4187-9666-4ef4829e3a7&title=&width=534)
使用docker-compose的情况下，要求环境所在的机器必须存在docker以及docker-compose的运行环境。
在使用k8s的情况下，要求所在的机器必须存在k8s的运行环境。

###### 使用docker-compose运行环境时搭建
这里我们准备一台机器作为测试环境的运行服务器，当然如果你的机器十分紧缺，只有一台的情况下，你也可以时候构建服务器进行运行服务。
1）进行docker以及docker-compose的环境搭建，这里的搭建细节我不再描述，因为前面在搭建构建服务器器的时候我已经说过了。
2）安装构建插件，构建插件主要是用于接收处理构建服务器的部署请求的，实现在运行服务器上运行。
```yaml
# 创建运行目录
mkdir -p /usr/local/devops/dc-service && cd /usr/local/devops/dc-service

# 下载构建插件源码
git clone https://github.com/limeschool/devops.git

# 进入构建服务源码目录
cd devops/dc-service/

# 修改配置文件
vim config/conf.json

# 主要修改配置如下
  "request-token": "limeschool", #连接时使用的token，需要修改成自己的。
  "nginx_hosts_path": "/etc/nginx/conf.d", # nginx的hosts配置目录，默认安装不用修改
  "nginx_nginx_path": "/etc/nginx/nginx.conf", # nginx的服务配置目录，默认安装不用修改
  "exec_type": "/bin/sh",  # 默认安装不用修改
  "work_dir": "/usr/local/dc-service" # 插件的工作目录，默认安装不用修改

```
```yaml
{
  "service": "dc-service",
  "system": {
    "skip_request_log": {
      "get:/auth": true
    },
    "client_timeout":"60s"
  },
  "log": {
    "level": 0,
    "trace_key": "log-id",
    "service_key": "service"
  },
  "request-token": "limeschool",
  "nginx_hosts_path": "/etc/nginx/conf.d",
  "nginx_nginx_path": "/etc/nginx/nginx.conf",
  "exec_type": "/bin/sh",
  "work_dir": "/usr/local/dc-service"
}
```
配置保存了之后则可以进行插件运行了
```yaml
nohup ./dc-service -c config/conf.json &
```
接下来在后台管理中进行连接配置并保存
![image.png](https://cdn.nlark.com/yuque/0/2023/png/26359342/1672597255680-90d19cfd-f2c8-4e89-9343-e87ef1e0aea4.png#averageHue=%23e3452c&clientId=u640b5be3-27b9-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=590&id=u46eb58e2&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1180&originWidth=2190&originalType=binary&ratio=1&rotation=0&showTitle=false&size=214829&status=done&style=none&taskId=u47959ec5-5feb-4606-b285-23caef808f9&title=&width=1095)

###### 使用k8s运行环境时搭建
1）使用资源清单在k8s中创建角色账号
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kube-system

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: admin-user
    namespace: kube-system
```
2）查看token
```yaml
# 获取密钥
kubectl get secret -n kube-system 

# 我们可以看见有一个admin-user-token-xxxx的资源
# admin-user-token-jx9pl
kubectl describe secret admin-user-token-jx9pl  -n kube-system 

# 复制里面的token信息作为在后台的token中添加即可。
```

##### 服务管理
##### 服务构建
##### 服务发布
##### 网络管理









### 配置中心安装
我们登陆到后台之后
```yaml
version: '3'
services:
  consul-test:
    container_name: consul-test
    image: consul:1.13.3
    restart: always
    environment:
      TZ: Asia/Shanghai
    ports:
      - 8500:8500
    volumes:
      - ./consul/conf:/consul/conf
      - ./consul/data:/consul/data
    privileged: true
```

