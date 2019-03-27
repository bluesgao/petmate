package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petmate/handler/sys"
	"petmate/handler/user"
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
		sysGroup.GET("/ping", sys.Ping)
		sysGroup.GET("/disk", sys.Disk)
		sysGroup.GET("/cpu", sys.Cpu)
		sysGroup.GET("/mem", sys.Mem)
		sysGroup.GET("/name", sys.Name)

	}

	//用户相关
	userGroup := g.Group("/user")
	{
		userGroup.POST("", user.Create)
	}

	return g
}
