package model

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var ResponseCode struct {
	Success         int
	Failure         int
	Unauthorized    int
	Forbidden       int
	Timeout         int
	Exception       int
	TooManyRequests int
}

var ResponseMsg struct {
	Exception       string
	Unauthorized    string
	Forbidden       string
	TooManyRequests string
}

func init() {
	ResponseCode.Success = 200
	ResponseCode.Failure = 400
	ResponseCode.Unauthorized = 401
	ResponseCode.Forbidden = 403
	ResponseCode.Timeout = 408
	ResponseCode.Exception = 500
	ResponseCode.TooManyRequests = 429

	ResponseMsg.Exception = "系统繁忙，请稍后尝试"
	ResponseMsg.Unauthorized = "请重新登录"
	ResponseMsg.Forbidden = "无权操作"
	ResponseMsg.TooManyRequests = "请求过于频繁,请稍后重试"
}
