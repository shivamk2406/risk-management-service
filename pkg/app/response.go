package app

import (
	"github.com/gin-gonic/gin"
	err "github.com/shivamk2406/risk-management-service/pkg/err"
)
type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  err.GetMsg(errCode),
		Data: data,
	})
	return
}