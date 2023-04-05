## make run 项目跑起来
## make swag 安装swag命令,此命令主要生成接口文档

init:
	go get github.com/pilu/fresh
	go install github.com/pilu/fresh
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init
	go mod tidy


run:
	swag init;go mod tidy;fresh ## 热启动


restart:
	kill -1 `pgrep go-blog`;## 平滑重启


