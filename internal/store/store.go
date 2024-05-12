package store

import (
	"github.com/jinzhu/gorm"
	"github.com/redis/go-redis/v9"
	"kuroko/internal/store/demo"
)

type (
	Factory interface {
		Cache() demo.Store
	}
	DataStore struct {
		DB    *gorm.DB
		Redis *redis.Client
	}
)

var _ Factory = (*DataStore)(nil)

func (ds *DataStore) Cache() demo.Store {
	return demo.NewDemo(ds)
}

func (ds DataStore) GetMySQL() *gorm.DB {
	return ds.DB
}
func (ds DataStore) GetRedis() *redis.Client {
	return ds.Redis
}
