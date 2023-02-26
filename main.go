package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/boot"
	"go-blog/internal/router"
	"go-blog/pkg/middleware"
	"log"
)

func main() {
	r := gin.New()

	boot.InitConf()
	middleware.Register(r)
	router.Register(r)

	log.Fatal(r.Run(fmt.Sprintf(":%d", boot.Cfg.App.Port)))
}
