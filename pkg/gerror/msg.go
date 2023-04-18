package gerror

var ResponseMsg struct {
	Exception       string
	Unauthorized    string
	Forbidden       string
	TooManyRequests string
	Failure         string
}

func init() {
	ResponseMsg.Exception = "系统繁忙，请稍后尝试"
	ResponseMsg.Unauthorized = "请重新登录"
	ResponseMsg.Forbidden = "无权操作"
	ResponseMsg.TooManyRequests = "请求过于频繁,请稍后重试"
	ResponseMsg.Failure = "操作失败"
}
