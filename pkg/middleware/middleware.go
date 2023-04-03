package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Init(r *gin.Engine) {

	r.Use(
		otelgin.Middleware("my-server"),
		Logger(otelzap.L()),
		Recovery(otelzap.L(), true),
		cors.Default(),
		Catch(),

	)
}
