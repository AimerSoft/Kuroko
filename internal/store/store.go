package store

import (
	"github.com/jinzhu/gorm"
	"kuroko/internal/store/demo"
)

type (
	Factory interface {
		Demo() demo.Store
	}
	DataStore struct {
		DB *gorm.DB
	}
)

var _ Factory = (*DataStore)(nil)

func (ds *DataStore) Demo() demo.Store {
	return demo.NewDemo(ds)
}

func (ds DataStore) GetMySQL() *gorm.DB {
	return ds.DB
}
