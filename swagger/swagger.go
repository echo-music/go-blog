package swagger

import (
	_ "github.com/echo-music/go-blog/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(r *gin.Engine) {

	if gin.Mode() == gin.DebugMode || gin.Mode() == gin.TestMode {

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

}
