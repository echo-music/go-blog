package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var cache *redis.Client
var once sync.Once

func Redis() *redis.Client {
	return cache
}

func Init(cfg Config) {

	once.Do(func() {
		cache = redis.NewClient(&redis.Options{
			Addr:     cfg.Host,
			Password: cfg.Password, // no password set
			DB:       cfg.DB,       // use default DB
		})
		_, err := cache.Ping(context.Background()).Result()
		if err != nil {
			panic(fmt.Sprintf("Redis connect ping failed, err:%+v", err))
		}
	})
}

func Set(key string, value interface{}, expire time.Duration) (err error) {
	_, err = cache.Set(context.Background(), key, value, expire).Result()
	return
}
