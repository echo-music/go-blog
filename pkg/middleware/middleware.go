package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(r *gin.Engine) {

	r.Use(
		Logger(zap.L()),
		Recovery(zap.L(), true),
		cors.Default(),
		Catch(),
	)

}
