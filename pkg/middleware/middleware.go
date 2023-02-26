package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.Use(
		gin.Logger(),
		cors.Default(),
		gin.Recovery(),
		Catch(),
	)

}
