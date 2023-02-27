package middleware

import (
	"github.com/gin-gonic/gin"
	"go-blog/pkg/response"
)

// Catch 捕获异常
func Catch() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		length := len(c.Errors)
		if length <= 0 {
			return
		}

		e := c.Errors[length-1]
		response.Error(c, e)
		return

	}
}
