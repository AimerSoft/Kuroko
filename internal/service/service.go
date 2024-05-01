package service

import (
	"kuroko/internal/service/demo"
	"kuroko/internal/store"
)

type Service interface {
	Demo() demo.Service
}

var _ Service = (*service)(nil)

type service struct {
	store store.Factory
}

func (s *service) GetStore() store.Factory {
	return s.store
}

// NewService ...
func NewService(factory store.Factory) Service {
	return &service{
		store: factory,
	}
}
func (s *service) Demo() demo.Service {
	return demo.NewDemo(s)
}
