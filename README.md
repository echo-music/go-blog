# go-blog
```
 这是用gin搭建的app用例

```
## 一、项目结构说明
### boot
```
项目配置初始化，mysql,redis初始化；
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
go mod tidy

注意:确保go.mod文件存在，否则执行 go mod init github.com/echo-music/go-blog ）

```

### 5）项目跑起来（项目根目录下执行）
```
make run
```


