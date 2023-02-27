package middleware

import (
	"github.com/gin-gonic/gin"
	"go-blog/pkg/library/gerror"
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
		code := gerror.Code(e)
		msg := e.Error()
		switch code {
		case 0, -1:
			code = gerror.ResponseCode.Failure
		case gerror.ResponseCode.Exception:
			msg = gerror.ResponseMsg.Exception
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": msg,
		})
		return

	}
}
