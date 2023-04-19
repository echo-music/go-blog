package middleware

import (
	"bytes"
	"fmt"
	"github.com/echo-music/go-blog/pkg/known"
	"github.com/echo-music/go-blog/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"strings"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.RequestURI, "/swagger/") {
			return
		}
		start := time.Now()
		bodyLogWriter := &response.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		c.Next()

		logContent := []zapcore.Field{
			zap.Any(known.XRequestIDKey, c.Value(known.XRequestIDKey)),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("res", bodyLogWriter.Body.String()),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("cost", time.Since(start)),
		}
		if e := c.Errors.ByType(gin.ErrorTypePrivate); len(e) > 0 {
			stack := fmt.Sprintf("%+v", e.Last().Err)
			data := strings.Split(strings.ReplaceAll(stack, "\t", ""), "\n")
			d := make([]string, 5)
			copy(d, data)
			if gin.Mode() == gin.DebugMode {
				fmt.Println(stack)
			} else {
				logContent = append(logContent,
					zap.Any("stack", d),
				)
			}

		}

		msg := http.StatusText(c.Writer.Status())
		if c.Writer.Status() != 200 || c.Errors != nil {
			zap.L().Error(msg, logContent...)
		} else {
			zap.L().Info(msg, logContent...)
		}

	}
}
