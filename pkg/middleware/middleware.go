package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	mws := []gin.HandlerFunc{
		RequestID(),    //客户端请求ID
		Recovery(),     //系统异常
		Logger(),       //请求日志
		cors.Default(), //跨域
		Catch(),        //错误输出
	}
	r.Use(mws...)
}
