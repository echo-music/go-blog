package gerror

var ResponseCode struct {
	Success         int
	Failure         int
	Unauthorized    int
	Forbidden       int
	Timeout         int
	Exception       int
	TooManyRequests int
}

func init() {
	ResponseCode.Success = 200
	ResponseCode.Failure = 400
	ResponseCode.Unauthorized = 401
	ResponseCode.Forbidden = 403
	ResponseCode.Timeout = 408
	ResponseCode.Exception = 500
	ResponseCode.TooManyRequests = 429

}
