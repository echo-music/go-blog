run:
	swag init;go mod tidy;go run main.go
swag:
	go install github.com/swaggo/swag/cmd/swag@latest