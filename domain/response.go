package domain

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ErrorResponse(msg string) Response {
	return Response{
		Code: -1,
		Msg:  msg,
	}
}

func SuccessResponse(data ...any) Response {
	return Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}
