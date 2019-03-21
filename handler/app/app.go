package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

//应用基本信息

//应用名称
func Name(c *gin.Context) {
	c.String(http.StatusOK, viper.GetString("app.name"))
}
