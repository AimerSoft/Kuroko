package demo

import (
	"github.com/gin-gonic/gin"
	"kuroko/internal/service"
	"kuroko/internal/store"
	"kuroko/pkg/core"
)

type Controller interface {
	Hello(ctx *gin.Context)
}

type demoController struct {
	srv service.Service
}

func NewDemoController(factory store.Factory) *demoController {
	return &demoController{
		srv: service.NewService(factory),
	}
}
func (c demoController) Hello(ctx *gin.Context) {
	// todo: 处理相关逻辑，调用service方法
	c.srv.Demo().Hello()
	core.SendResponse(ctx, nil, "ok")
}
