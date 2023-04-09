package middleware

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	mws := []gin.HandlerFunc{
		RequestID(), //客户端请求ID
		Logger(),    //请求日志
		Recovery(),  //系统异常
		Cors(),      //跨域
		Csrf(),      //csrf
		Catch(),     //错误输出
	}
	r.Use(mws...)
}
