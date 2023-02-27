package middleware

import (
	"github.com/echo-music/go-blog/pkg/gerror"
	"github.com/echo-music/go-blog/pkg/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
