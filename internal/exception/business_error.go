package exception

type BusinessError struct {
	Code int
	Msg  string
}

func (e *BusinessError) Error() string {
	return e.Msg
}

func New(code int, msg string) *BusinessError {
	return &BusinessError{Code: code, Msg: msg}
}
