package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(r *gin.Engine) {

	r.Use(
		Recovery(zap.L(), true),
		Logger(zap.L()),
		cors.Default(),
		Catch(),
	)

}
