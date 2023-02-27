package response

import (
	"github.com/echo-music/go-blog/pkg/gerror"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = "ok"
	}

	c.JSON(http.StatusOK, Response{
		Code: gerror.ResponseCode.Success,
		Msg:  msg,
		Data: data,
	})

}

func Error(c *gin.Context, e error) {
	if e == nil {
		return
	}
	code := gerror.Code(e)
	msg := e.Error()
	switch code {
	case 0, -1:
		code = gerror.ResponseCode.Failure
	case gerror.ResponseCode.Exception:
		msg = gerror.ResponseMsg.Exception
	}

	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
	})

}
