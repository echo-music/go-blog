package g

import (
	"github.com/echo-music/go-blog/pkg/cache"
	"github.com/echo-music/go-blog/pkg/db"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Redis() *redis.Client {
	return cache.Redis()
}

func DB() *gorm.DB {
	return db.DB()
}
