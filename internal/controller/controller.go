package controller

import (
	demo "kuroko/internal/controller/v1"
	"kuroko/internal/store"
)

type Controller interface {
	Demo() demo.Controller
}

var _ Controller = (*controller)(nil)

type controller struct {
	store store.Factory
}

func (c controller) Demo() demo.Controller {
	return demo.NewDemoController(c.store)
}

// NewController ...
func NewController(factory store.Factory) Controller {
	return &controller{
		store: factory,
	}
}
