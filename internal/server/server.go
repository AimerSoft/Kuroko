package server

import (
	"github.com/gin-gonic/gin"
	"kuroko/internal/option"
	"kuroko/internal/route"
	"kuroko/internal/store"
)

type Server struct {
	Web *gin.Engine
}

func NewServer(opt *option.Options) (*Server, error) {
	sro := &store.DataStore{
		DB: nil, // todo: DB初始化
	}
	g := route.ApiServer(opt.WebOptions.Engine, sro)
	return &Server{
		Web: g,
	}, nil
}
