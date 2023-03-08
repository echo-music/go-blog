## make run 项目跑起来
## make swag 安装swag命令,此命令主要生成接口文档


run:
	swag init;go mod tidy;go run main.go
swag:
	go install github.com/swaggo/swag/cmd/swag@latest