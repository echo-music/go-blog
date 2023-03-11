package middleware

import (
	"bytes"
	"context"
	"github.com/echo-music/go-blog/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"time"
)

func Tracing() gin.HandlerFunc {

	return func(c *gin.Context) {
		var (
			newCtx context.Context
			span   opentracing.Span
		)
		start := time.Now()

		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)

		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				opentracing.GlobalTracer(),
				c.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				opentracing.GlobalTracer(),
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{
					Key:   string(ext.Component),
					Value: "HTTP",
				},
			)
		}

		defer span.Finish()

		var (
			traceID     string
			spanID      string
			spanContext = span.Context()
		)
		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerContext := spanContext.(jaeger.SpanContext)
			traceID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
		}
		c.Set("X-Trace-ID", traceID) // 后续取出
		c.Set("X-Span-ID", spanID)

		c.Request = c.Request.WithContext(opentracing.ContextWithSpan(newCtx, span))

		bodyLogWriter := &response.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		c.Next()

		span.SetTag("status", c.Writer.Status())
		span.SetTag("method", c.Request.Method)
		span.SetTag("path", c.Request.URL)
		span.SetTag("query", c.Request.URL.RawQuery)
		span.SetTag("res", bodyLogWriter.Body.String())
		span.SetTag("ip", c.ClientIP())
		span.SetTag("user-agent", c.Request.UserAgent())
		span.SetTag("errors", c.Errors.ByType(gin.ErrorTypePrivate).String())
		span.SetTag("trace_id", traceID)
		span.SetTag("span_id", spanID)
		span.SetTag("cost", time.Since(start))

	}
}
