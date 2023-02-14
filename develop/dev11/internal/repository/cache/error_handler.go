package cache

type ErrorHandler struct {
	Err        error `json:"err"`
	StatusCode int   `json:"statusCode"`
}

func NewErrorHandler(err error, statusCode int) ErrorHandler {
	return ErrorHandler{Err: err, StatusCode: statusCode}
}

func (m ErrorHandler) Error() string {
	return m.Err.Error()
}
