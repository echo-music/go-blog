package middleware

import (
	"bytes"
	"github.com/echo-music/go-blog/pkg/response"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func Logger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		bodyLogWriter := &response.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		c.Next()
		span := trace.SpanFromContext(c.Request.Context())
		defer span.End()
		logContent := []zapcore.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("res", bodyLogWriter.Body.String()),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", time.Since(start)),
		}

		if c.Writer.Status() != 200 || c.Errors != nil {
			logger.Error(path, logContent...)
		} else {
			logger.Info(path, logContent...)
		}

	}
}
