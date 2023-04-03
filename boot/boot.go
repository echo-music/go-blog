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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"log"
	"syscall"
	"time"
)

const (
	service     = "trace-demo"
	environment = "production"
	id          = 1
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
	gin.SetMode(gin.DebugMode)

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
	//初始化中间件,路由,swagger等
	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

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
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}
