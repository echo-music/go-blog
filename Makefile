## make run 项目跑起来
## make swag 安装swag命令,此命令主要生成接口文档

init:
	go mod init github.com/echo-music/go-blog
	go mod tidy
	go get github.com/pilu/fresh
	go install github.com/pilu/fresh
	go install github.com/swaggo/swag/cmd/swag@latest

run:
	swag init;go mod tidy;fresh

