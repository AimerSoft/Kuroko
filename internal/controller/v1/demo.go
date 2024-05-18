package demo

import (
	"github.com/gin-gonic/gin"
	"kuroko/internal/service"
	"kuroko/internal/store"
	"kuroko/pkg/core"
)

type Controller interface {
	TinyUrl(ctx *gin.Context)
	Redirect(ctx *gin.Context)
}

type demoController struct {
	srv service.Service
}

func NewDemoController(factory store.Factory) *demoController {
	return &demoController{
		srv: service.NewService(factory),
	}
}
func (c demoController) TinyUrl(ctx *gin.Context) {
	form := ctx.PostForm("url")
	day := ctx.PostForm("day")
	if day == "" {
		day = "30"
	}
	tinyUrl, err := c.srv.Demo().TinyUrl(form, day)
	if err != nil {
		core.SendResponse(ctx, err, "error")
		return
	}
	core.SendResponse(ctx, nil, tinyUrl)
}

func (c demoController) Redirect(ctx *gin.Context) {
	form := ctx.Param("code")
	tinyUrl, err := c.srv.Demo().GetTinyUrl(form)
	if err != nil {
		ctx.Err()
		return
	}
	ctx.Redirect(301, tinyUrl)
}
