package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controller2 "kuroko/internal/controller"
	"kuroko/internal/store"
)

func ApiServer(g *gin.Engine, sro *store.DataStore) *gin.Engine {
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	controller := controller2.NewController(sro)
	// 使用CORS中间件
	g.Use(cors.Default())
	g.GET("/tinyurl/:code", controller.Demo().Redirect)
	g.POST("/tinyurl", controller.Demo().TinyUrl)
	return g
}
