package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-blog/pkg/gerror"
	"go-blog/pkg/response"
)

func Register(r *gin.Engine) {

	r.Use(
		gin.Logger(),
		cors.Default(),
		Catch(),
	)

	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		// 程序panic需要报警
		response.Error(c, gerror.Exception(err.(string)))
	}))

}
