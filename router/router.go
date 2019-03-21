package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petmate/handler/appinfo"
	"petmate/handler/sysinfo"
)

//路由
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	//g.Use(mw...)
	g.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "the incorrect api route.")
	})

	sysGroup := g.Group("/sysinfo")
	{
		sysGroup.GET("/ping", sysinfo.Ping)
		sysGroup.GET("/disk", sysinfo.Disk)
		sysGroup.GET("/cpu", sysinfo.Cpu)
		sysGroup.GET("/mem", sysinfo.Mem)
	}

	appGroup := g.Group("/appinfo")
	{
		appGroup.GET("/name", appinfo.Name)
	}

	return g
}
