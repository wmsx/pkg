package gin

import "github.com/gin-gonic/gin"

type HandlerFunc func(c *gin.Context) error

// api错误的结构体
type APIException struct {
	Code int    `json:"-"`
	Msg  string `json:"msg"`
}

func (e *APIException) Error() string {
	return e.Msg
}

func Wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
		)

		err = handler(c)
		if err != nil {
			var apiException *APIException
			if h, ok := err.(*APIException); ok {
				apiException = h
			} else {
				apiException = ServerError()
			}
			c.JSON(apiException.Code, apiException)
			return
		}
	}
}
