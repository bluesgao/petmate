package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petmate/handler/sys"
)

//路由
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	//g.Use(mw...)
	g.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "the incorrect api route.")
	})

	sysGroup := g.Group("/sys")
	{
		sysGroup.GET("/ping", sys.Ping)
		sysGroup.GET("/disk", sys.Disk)
		sysGroup.GET("/cpu", sys.Cpu)
		sysGroup.GET("/mem", sys.Mem)
	}

	return g
}
