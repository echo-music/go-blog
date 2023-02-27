package main

import (
	"fmt"
	"github.com/echo-music/go-blog/boot"
	"github.com/echo-music/go-blog/internal/router"
	"github.com/echo-music/go-blog/pkg/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.New()
	middleware.Register(r)
	router.Register(r)

	log.Fatal(r.Run(fmt.Sprintf(":%d", boot.Cfg.App.Port)))
}
