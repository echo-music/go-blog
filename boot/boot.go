package boot

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/echo-music/go-blog/internal/router"
	"github.com/echo-music/go-blog/pkg/cache"
	"github.com/echo-music/go-blog/pkg/db"
	"github.com/echo-music/go-blog/pkg/logs"
	"github.com/echo-music/go-blog/pkg/middleware"
	"github.com/echo-music/go-blog/swagger"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"syscall"
)

type config struct {
	App struct {
		Mode    string
		Name    string
		Port    int
		Version string
	}
	Mysql  db.Config
	Redis  cache.Config
	Logger logs.Config
}

var Cfg config

func Run() {

	//设置debug模式
	gin.SetMode(gin.ReleaseMode)

	//读取配置文件
	if _, err := toml.DecodeFile("./config/app.toml", &Cfg); err != nil {
		panic("decode config file err")
	}

	//初始化日志
	logs.Init(Cfg.Logger)
	defer logs.Sync()

	//初始化数据库
	db.Init(Cfg.Mysql)

	//初始化redis
	cache.Init(Cfg.Redis)

	r := gin.New()

	middleware.Init(r)
	router.Init(r)
	swagger.Init(r)

	server := endless.NewServer(fmt.Sprintf(":%d", Cfg.App.Port), r)
	server.BeforeBegin = func(add string) {
		log.Printf("server is listening at port %d , Actual pid is %d", Cfg.App.Port, syscall.Getpid())
	}

	log.Fatal(server.ListenAndServe())
}
