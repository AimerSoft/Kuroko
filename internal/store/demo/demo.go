package demo

import "github.com/jinzhu/gorm"

type (
	Store interface {
		Hello()
	}
	demoMethod interface {
		GetMySQL() *gorm.DB
	}
	demoStore struct {
		db *gorm.DB
	}
)

func NewDemo(ds demoMethod) *demoStore {
	return &demoStore{
		db: ds.GetMySQL(),
	}
}

func (ds *demoStore) Hello() {
	print("hello world")
}
