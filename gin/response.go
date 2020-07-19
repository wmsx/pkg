package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode           = 0    // 一切正常
	LogicErrorCode        = 1000 // 业务逻辑异常
	UnauthorizedErrorCode = 1001 // 未登陆
	ServerErrorCode       = 9999 // 服务器未知异常
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) ServerErrorResponse() {
	g.C.JSON(http.StatusOK, Response{
		Code:    ServerErrorCode,
		Message: "服务器异常",
	})
	return
}

func (g *Gin) LogicErrorResponse(message string) {
	g.C.JSON(http.StatusOK, Response{
		Code:    LogicErrorCode,
		Message: message,
	})
	return
}

func (g *Gin) Response(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code:    SuccessCode,
		Message: "success",
		Data:    data,
	})
	return
}
