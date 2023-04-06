package middleware

import (
	"github.com/echo-music/go-blog/pkg/gerror"
	"github.com/echo-music/go-blog/pkg/logs"
	"github.com/echo-music/go-blog/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func Recovery(logger *zap.Logger) gin.HandlerFunc {
	out := logs.Writer()
	if gin.Mode() == gin.DebugMode {
		out = os.Stderr
	}
	return gin.RecoveryWithWriter(out, func(c *gin.Context, err interface{}) {
		logger.Error(gerror.ResponseMsg.Exception, zap.Any("err", err), zap.Stack("caller"))
		response.Error(c, gerror.Exception(""))
	})
}
