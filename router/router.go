package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petmate/handler"
)

//路由
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	//g.Use(mw...)
	g.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "the incorrect api route.")
	})

	//系统相关
	sysGroup := g.Group("/sysinfo")
	{
		sysGroup.GET("/ping", handler.Ping)
		sysGroup.GET("/disk", handler.Disk)
		sysGroup.GET("/cpu", handler.Cpu)
		sysGroup.GET("/mem", handler.Mem)
	}

	//应用相关
	appGroup := g.Group("/appinfo")
	{
		appGroup.GET("/name", handler.Name)
	}

	//用户相关
	userGroup := g.Group("/user")
	{
		userGroup.POST("/create", handler.Create)
	}

	return g
}
