# go-blog

这是用gin搭建的web用例,如果你是小白，那它适合你练手！

## gin 框架文档

https://gin-gonic.com/docs/examples/serving-data-from-reader/

1. 新增日志中间件

2. 新增recover中间件

3. 新增异常捕获中间件

4. 新增跨域中间件

5. 新增热启动(执行make run即可)

6. 新增平滑启动

## 一、项目结构说明

```
.
├── Makefile
├── README.md
├── boot
│   └── boot.go
├── config
│   └── app.toml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go-blog
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   └── user.go
│   ├── router
│   │   └── router.go
│   └── service
│       └── order.go
├── main.go
├── pkg
│   ├── gerror
│   │   ├── code.go
│   │   ├── gerror.go
│   │   └── msg.go
│   ├── logs
│   │   └── logs.go
│   ├── middleware
│   │   ├── catch.go
│   │   ├── logger.go
│   │   ├── middleware.go
│   │   └── recovery.go
│   ├── model
│   │   └── user.go
│   ├── response
│   │   └── response.go
│   └── store
│       ├── cache
│       │   ├── cache.go
│       │   └── config.go
│       └── mysql
│           ├── config.go
│           └── db.go
├── runtime
│   └── logs
│       └── a.log
├── swagger
│   └── swagger.go
└── tmp
    └── runner-build




```

### boot

```
项目启动包,配置初始化，mysql,redis初始化；
```

### config

```
项目的配置文件
```

### internal

```
内部包目录,禁止其他go项目引入里面的包
```

### pkg

```
公共包目录，任何go项目都可引入公共包
```

## 二、快速体验

### 1）搭建go环境（ mac下安装）

#### 1-1、安装：

http://c.biancheng.net/view/3994.html

#### 1-2、设置go环境变量:

```
  项目目录：{你的项目地址}/goweb,goweb目录下新建三个目录如下
  bin pkg src
  
  打开.bash_profile 文件写入以下go环境配置
 
  export GOROOT=/usr/local/go
  export GOPATH={你的项目地址}/goweb
  export GOBIN=$GOPATH/bin
  export PATH=$GOBIN:$PATH
``` 

#### 1-3、开启go mod

```
  export GO111MODULE=on
```

#### 1-4、设置拉取依赖库代理地址

```
go env -w GOPROXY=https://goproxy.cn,direct
```

### 2) 拉取项目

```
git clone git@github.com:echo-music/go-blog.git
```

### 3）更新依赖包（项目根目录下执行)

```
go mod download
go mod tidy

注意:确保go.mod文件存在，否则执行 go mod init github.com/echo-music/go-blog ）

```

### 4) 新建user表

```
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

### 4）项目配置文件,修改数据库，redis配置

```
[app]
name = "go_gin"
port = 8081
version = "1.0.0"

[mysql]
type = "mysql"
link = "root:123456@tcp(localhost:3306)/test"
charset = "utf8mb4"
maxIdle = 10
maxOpen = 100
maxLifetime = "30s"

[redis]
host = "127.0.0.1:6379"
password = "123456"
db = 0

[logger]
fileName = "./runtime/logs/a.log"
maxSize = 1
maxBackups = 30
maxAges = 7
compress = true
level = ""
```

### 5）项目跑起来（项目根目录下执行）

```
## 项目首次启动运行该命令
make init

## 项目跑起来
make run
```

### 6) 访问接口

http://127.0.0.1:8081/users/

### 7) 访问swagger接口文档

http://localhost:8081/swagger/index.html

### 8）接入链路追踪

接入jaeger链路追踪
https://github.com/open-telemetry/opentelemetry-go/tree/main/exporters/jaeger

监控系统UI

https://uptrace.dev/get/opentelemetry-go.html#getting-started

https://github.com/uptrace/uptrace/blob/master/README.zh.md

https://cloud.tencent.com/developer/article/2193995

监控系统访问地址

http://localhost:14318/overview/2/?time_gte=20230403T083600&time_dur=3600
