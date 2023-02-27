package boot

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"go-blog/pkg/cache"
	"go-blog/pkg/db"
)

type config struct {
	App   App
	Mysql db.Config
	Redis cache.Config
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

}
