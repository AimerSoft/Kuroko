package demo

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/redis/go-redis/v9"
	"time"
)

type (
	Store interface {
		SetCache(key string, value string, ttl time.Duration) error
		GetCache(key string) (string, error)
	}
	demoMethod interface {
		GetMySQL() *gorm.DB
		GetRedis() *redis.Client
	}
	demoStore struct {
		db    *gorm.DB
		redis *redis.Client
	}
)

func NewDemo(ds demoMethod) *demoStore {
	return &demoStore{
		db:    ds.GetMySQL(),
		redis: ds.GetRedis(),
	}
}

func (ds *demoStore) SetCache(key string, value string, ttl time.Duration) error {
	err := ds.redis.Set(context.Background(), key, value, ttl).Err()
	return err
}

func (ds *demoStore) GetCache(key string) (string, error) {
	return ds.redis.Get(context.Background(), key).Result()
}
