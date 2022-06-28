package response

type ResponseMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type PageResponse[T any] struct {
	CurrentPage int64 `json:"currentPage"`
	PageSize    int64 `json:"pageSize"`
	Total       int64 `json:"total"`
	Pages       int64 `json:"pages"` // 总页数
	Data        []T   `json:"data"`
}

const (
	success         = 200
	failed          = 500
	validate_failed = 404
	unauthorized    = 401
	forbidden       = 403
)

func SuccessMsg(data interface{}) *ResponseMsg {
	msg := &ResponseMsg{
		Code: success,
		Msg:  "SUCCESS",
		Data: data,
	}
	return msg
}

func FailedMsg(data interface{}) *ResponseMsg {
	msg := &ResponseMsg{
		Code: failed,
		Msg:  "FAILED",
		Data: data,
	}
	return msg
}

func ValidateFailedMsg(data interface{}) *ResponseMsg {
	msg := &ResponseMsg{
		Code: validate_failed,
		Msg:  "VALIDATE_FAILED",
		Data: data,
	}
	return msg
}

func UnauthorizedMsg(data interface{}) *ResponseMsg {
	msg := &ResponseMsg{
		Code: unauthorized,
		Msg:  "UNAUTHORIZED",
		Data: data,
	}
	return msg
}

func ForbiddenMsg(data interface{}, msg ...interface{}) *ResponseMsg {

	m := &ResponseMsg{
		Code: failed,
		Msg:  "FORBIDDEN",
		Data: data,
	}
	return m
}
