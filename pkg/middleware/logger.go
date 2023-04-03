package middleware

import (
	"bytes"
	"fmt"
	"github.com/echo-music/go-blog/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func Logger(logger *otelzap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		bodyLogWriter := &response.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		c.Next()
		span := trace.SpanFromContext(c.Request.Context())
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
			zap.String("trace_id", span.SpanContext().TraceID().String()),
			zap.String("span_id", span.SpanContext().SpanID().String()),
		}

		attr := []attribute.KeyValue{
			attribute.Int("status", c.Writer.Status()),
			attribute.String("method", c.Request.Method),
			attribute.String("path", path),
			attribute.String("query", query),
			attribute.String("res", bodyLogWriter.Body.String()),
			attribute.String("ip", c.ClientIP()),
			attribute.String("user-agent", c.Request.UserAgent()),
			attribute.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			attribute.String("cost", time.Since(start).String()),
			attribute.String("trace_id", span.SpanContext().TraceID().String()),
			attribute.String("span_id", span.SpanContext().SpanID().String()),
		}

		span.SetAttributes(attr...)
		defer span.End()

		if c.Writer.Status() != 200 || c.Errors != nil {
			fmt.Println(c.Writer.Status())
			logger.Ctx(c.Request.Context()).Error(path, logContent...)
		} else {
			logger.Ctx(c.Request.Context()).Info(path, logContent...)
		}

	}
}
