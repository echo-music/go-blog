package router

import (
	"github.com/echo-music/go-blog/internal/api"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	// order router
	api.User.Router(r)

	// goods router

	// trade router

}
