package boot

import (
	"github.com/BurntSushi/toml"
	"go-blog/pkg/library/db"
)

type config struct {
	App   *App
	Mysql *db.Config
}

type App struct {
	Name    string
	Port    int
	Version string
}

var Cfg config

func init() {

	//读取配置文件
	if _, err := toml.DecodeFile("./config/app.toml", &Cfg); err != nil {
		panic("decode config file err")

	}

	//初始化数据库
	db.InitDB(Cfg.Mysql)

	//初始化redis

}
