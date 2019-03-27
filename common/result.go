package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SendResponse(c *gin.Context, code ErrCode, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: code.Code,
		Msg:  code.Msg,
		Data: data,
	})
}
