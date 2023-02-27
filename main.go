package main

import (
	"fmt"
	_ "github.com/echo-music/go-blog/boot"
	"github.com/echo-music/go-blog/internal/router"
	"github.com/echo-music/go-blog/pkg/middleware"
	"github.com/echo-music/go-blog/swagger"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.New()

	middleware.Init(r)
	router.Init(r)
	swagger.Init(r)

	log.Fatal(r.Run(fmt.Sprintf(":%d", 8081)))
}
