package boot

import (
	"github.com/spf13/viper"
	"log"
)

type config struct {
	AppName    string
	AppPort    string
	AppVersion string
	MysqlPort  string
	MysqlHost  string
}

var Cfg *config

func InitConf() {

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			log.Println("找不到配置文件。。。")
			return
		} else {
			log.Println("配置文件出错。。。")
			return
		}
	}

	Cfg = &config{
		AppName:    viper.GetString("app.name"),
		AppPort:    viper.GetString("app.port"),
		AppVersion: viper.GetString("app.version"),
		MysqlPort:  viper.GetString("mysql.port"),
		MysqlHost:  viper.GetString("mysql.host"),
	}
}
