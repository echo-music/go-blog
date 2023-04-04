package middleware

import (
	"bytes"
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
			zap.String("trace_id", span.SpanContext().TraceID().String()),
		}

		attr := []attribute.KeyValue{
			attribute.String("res", bodyLogWriter.Body.String()),
			attribute.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			attribute.String("trace_id", span.SpanContext().TraceID().String()),
		}

		span.SetAttributes(attr...)
		if c.Errors != nil {
			span.RecordError(c.Errors.ByType(gin.ErrorTypePrivate)[0], trace.WithTimestamp(time.Now()), trace.WithStackTrace(true))
		}

		if gin.Mode() == gin.DebugMode {
			if c.Writer.Status() != 200 || c.Errors != nil {
				logger.Error(path, logContent...)
			} else {
				logger.Debug(path, logContent...)
			}
		}

	}
}
