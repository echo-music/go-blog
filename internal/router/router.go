package router

import (
	"github.com/gin-gonic/gin"
	"go-blog/internal/api"
)

func Register(r *gin.Engine) {

	// order router
	api.Order.Router(r)

	// goods router

	// trade router

}
