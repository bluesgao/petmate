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

func SendResult(c *gin.Context, bizErrCode BizErrCode, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: bizErrCode.Code,
		Msg:  bizErrCode.Msg,
		Data: data,
	})
}
