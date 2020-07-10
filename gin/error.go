package gin

import "net/http"

func ServerError() *APIException {
	return &APIException{
		Code: http.StatusInternalServerError,
		Msg:  "服务器异常",
	}
}

func LogicError(message string) *APIException {
	return &APIException{
		Code: http.StatusBadRequest,
		Msg:  message,
	}
}
