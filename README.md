# go-blog

```
 这是用gin搭建的app用例

```

## 一、项目结构说明

```
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
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   └── order.go
│   ├── router
│   │   └── router.go
│   └── service
│       └── order.go
├── main.go
├── pkg
│   ├── cache
│   │   ├── cache.go
│   │   └── config.go
│   ├── db
│   │   ├── config.go
│   │   └── db.go
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
│   │   └── order.go
│   └── response
│       └── response.go
├── runtime
│   └── logs
│       ├── a-2023-02-27T10-28-57.646.log
│       ├── a-2023-02-27T10-29-01.987.log
│       ├── a-2023-02-27T10-29-05.857.log
│       ├── a-2023-02-27T10-41-41.139.log
│       └── a.log
└── swagger
    └── swagger.go

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

```
   http://c.biancheng.net/view/3994.html
``` 

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

### 2）安装swag命令用于生成接口文档

```
go install github.com/swaggo/swag/cmd/swag@latest

```

### 3) 拉取项目

```
git clone git@github.com:echo-music/go-blog.git
```

### 4）更新依赖包（项目根目录下执行)

```
go mod download
go mod tidy

注意:确保go.mod文件存在，否则执行 go mod init github.com/echo-music/go-blog ）

```

### 5）项目配置文件,修改数据库，redis配置

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

### 6）项目跑起来（项目根目录下执行）

```
make run
```

### 7) 访问接口

```
http://127.0.0.1:8081/orders/
```

### 8) 新增平滑启动

```
https://github.com/fvbock/endless
```

### 9）本地热启动

```
仓库地址:
https://github.com/gravityblast/fresh


安装fresh:

go install github.com/pilu/fresh

项目下之下以下命令:

fresh
```

### 10) 访问swagger接口文档

```
http://localhost:8081/swagger/index.html
```