package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Use(
		RequestID(),
		Logger(),
		Recovery(),
		cors.Default(),
		Catch(),
	)
}
