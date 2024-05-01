package demo

import "kuroko/internal/store"

type (
	Service interface {
		// todo: service方法
		Hello()
	}
	StoreMethod interface {
		GetStore() store.Factory
	}
	demoService struct {
		store store.Factory
	}
)

func NewDemo(srv StoreMethod) *demoService {
	return &demoService{store: srv.GetStore()}
}

func (d demoService) Hello() {
	d.store.Demo().Hello()
}
