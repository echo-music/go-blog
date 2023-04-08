package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
)

var cli *redis.Client
var once sync.Once

func Redis() *redis.Client {
	if cli == nil {
		panic("未初始化redis")
	}
	return cli
}

func Init(cfg Config) {
	once.Do(func() {
		cli = redis.NewClient(&redis.Options{
			Addr:     cfg.Host,
			Password: cfg.Password, // no password set
			DB:       cfg.DB,       // use default DB
		})
		_, err := cli.Ping(context.Background()).Result()
		if err != nil {
			panic(fmt.Sprintf("Redis connect ping failed, err:%+v", err))
		}
	})
}
