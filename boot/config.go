package boot

import (
	"github.com/BurntSushi/toml"
	"go-blog/pkg/library/db"
)

type config struct {
	App   App
	Mysql db.Config
}

type App struct {
	Name    string
	Port    int
	Version string
}

var Cfg config

func InitConf() {

	if _, err := toml.DecodeFile("./config/app.toml", &Cfg); err != nil {
		panic("decode config file err")

	}

}
