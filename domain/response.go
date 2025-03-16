package domain

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewErrorResponse(msg string) Response {
	return Response{
		Code: -1,
		Msg:  msg,
	}
}
