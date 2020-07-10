package gin

import "github.com/gin-gonic/gin"

type HandlerFunc func(c *gin.Context) error

// api错误的结构体
type APIError struct {
	Code int    `json:"-"`
	Msg  string `json:"msg"`
}

func (e *APIError) Error() string {
	return e.Msg
}

func ErrorWrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
		)

		err = handler(c)
		if err != nil {
			var apiException *APIError
			if h, ok := err.(*APIError); ok {
				apiException = h
			} else {
				apiException = ServerError()
			}
			c.JSON(apiException.Code, apiException)
			return
		}
	}
}
