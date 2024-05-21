package http

import (
	"github.com/gin-gonic/gin"
	conf "kuroko/config"
	"net/http"
)

type WebOption struct {
	Mode            string
	Middlewares     []gin.HandlerFunc
	EnableProfiling bool
	EnableMetrics   bool
	*gin.Engine
	*http.Server
}

func NewWebOptions() *WebOption {
	port := conf.Port
	gin.SetMode(conf.Mode)
	return &WebOption{
		Mode:          conf.Mode,
		EnableMetrics: true,
		Middlewares:   []gin.HandlerFunc{},
		Engine:        gin.New(),
		Server: &http.Server{
			Addr: ":" + port,
		},
	}
}
