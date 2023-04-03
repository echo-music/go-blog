package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

func Init(r *gin.Engine) {

	r.Use(
		Logger(zap.L()),
		Recovery(zap.L(), true),
		cors.Default(),
		Catch(),
		otelgin.Middleware("my-server"),
	)
}
