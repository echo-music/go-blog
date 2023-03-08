# go-blog
```
使用gin 搭建的博客项目
```

## boot
```
配置文件读取，初始化项目依赖的服务，如mysql，redis等
```

## config
```
项目配置文件，如mysql,redis配置等
```

## internal
```
内部包目录,禁止其他go项目引入里面的包
```

## pkg
```
公共包目录，任何go项目都可引入公共包
```

## 快速体验
```
1）搭建go环境（ mac下安装）
  a、安装：
  http://c.biancheng.net/view/3994.html
  
  
  b、设置go环境变量:
  项目目录：{你的项目地址}/goweb,goweb目录下新建三个目录如下
  bin pkg src
  
  打开.bash_profile 文件写入以下go环境配置
 
  export GOROOT=/usr/local/go
  export GOPATH={你的项目地址}/goweb
  export GOBIN=$GOPATH/bin
  export PATH=$GOBIN:$PATH
  
  c、开启go mod
  export GO111MODULE=on
  
  d、设置拉取依赖库代理地址
  go env -w GOPROXY=https://goproxy.cn,direct

2）安装swag命令用于生成接口文档
go install github.com/swaggo/swag/cmd/swag@latest

3）更新依赖包（项目根目录下执行,确保go.mod文件存在，否则执行 go mod init github.com/echo-music/go-blog ）
go mod tidy

4）项目跑起来（项目根目录下执行）
make run

```