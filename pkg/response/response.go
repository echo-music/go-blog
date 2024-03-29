package response

import (
	"bytes"
	"github.com/echo-music/go-blog/pkg/gerror"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = "ok"
	}

	c.JSON(http.StatusOK, Result{
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
	c.Abort()
	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
	})

}

type BodyLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w BodyLogWriter) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
