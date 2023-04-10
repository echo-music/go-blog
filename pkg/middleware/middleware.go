package middleware

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	mws := []gin.HandlerFunc{
		RequestID(), //客户端请求ID
		Recovery(),  //系统异常
		Logger(),    //请求日志
		Cors(),      //跨域
		Csrf(),      //csrf攻击
		Catch(),     //错误输出
	}
	r.Use(mws...)
}
