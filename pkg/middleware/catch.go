package middleware

import (
	"github.com/gin-gonic/gin"
	gerror2 "go-blog/pkg/gerror"
	"net/http"
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
		code := gerror2.Code(e)
		msg := e.Error()
		switch code {
		case 0, -1:
			code = gerror2.ResponseCode.Failure
		case gerror2.ResponseCode.Exception:
			msg = gerror2.ResponseMsg.Exception
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": msg,
		})
		return

	}
}
