package middleware

import (
	"github.com/echo-music/go-blog/pkg/logs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.Use(
		Recovery(logs.ZapLog, true),
		Logger(logs.ZapLog),
		cors.Default(),
		Catch(),
	)

	//r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
	//	// 程序panic需要报警
	//	fmt.Println(err)
	//	response.Error(c, gerror.Exception(""))
	//}))

}
