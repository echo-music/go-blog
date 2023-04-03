package boot

import (
	"context"
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
	"github.com/uptrace/uptrace-go/uptrace"
	"log"
	"syscall"
)

type config struct {
	App    App
	Mysql  db.Config
	Redis  cache.Config
	Logger logs.Config
}

type App struct {
	Name    string
	Port    int
	Version string
}

var Cfg config

func init() {
	//设置debug模式
	gin.SetMode(gin.ReleaseMode)

	//读取配置文件
	if _, err := toml.DecodeFile("./config/app.toml", &Cfg); err != nil {
		panic("decode config file err")
	}

	//初始化数据库
	db.Init(Cfg.Mysql)

	//初始化redis
	cache.Init(Cfg.Redis)

	//初始化日志
	logs.Init(Cfg.Logger)

}

func Run() {
	////初始化中间件,路由,swagger等
	ctx := context.Background()

	// Configure OpenTelemetry with sensible defaults.
	uptrace.ConfigureOpentelemetry(
		// copy your project DSN here or use UPTRACE_DSN env var
		uptrace.WithDSN("http://project2_secret_token@localhost:14317/2"),

		uptrace.WithServiceName("myservice"),
		uptrace.WithServiceVersion("1.0.0"),
	)

	// Send buffered spans and free resources.
	defer uptrace.Shutdown(ctx)

	r := gin.New()

	middleware.Init(r)
	router.Init(r)
	swagger.Init(r)

	server := endless.NewServer(fmt.Sprintf(":%d", Cfg.App.Port), r)

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	log.Fatal(server.ListenAndServe())

}
