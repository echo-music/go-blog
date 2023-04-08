package g

import (
	"github.com/echo-music/go-blog/pkg/cache"
	"github.com/echo-music/go-blog/pkg/db"
	"github.com/echo-music/go-blog/pkg/logs"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Redis() *redis.Client {
	return cache.Redis()
}

func DB() *gorm.DB {
	return db.DB()
}

func Zap(c *gin.Context) *zap.Logger {
	return logs.Ctx(c)
}
