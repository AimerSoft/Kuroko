package route

import (
	"github.com/gin-gonic/gin"
	controller2 "kuroko/internal/controller"
	"kuroko/internal/store"
)

func ApiServer(g *gin.Engine, sro *store.DataStore) *gin.Engine {
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	controller := controller2.NewController(sro)
	g.GET("/init", controller.Demo().Hello)
	return g
}
