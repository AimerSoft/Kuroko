package demo

import (
	conf "kuroko/config"
	"kuroko/internal/pkg/tinyurl"
	"kuroko/internal/store"
	"log"
	"time"
)

type (
	Service interface {
		// todo: service方法
		TinyUrl(url string) (string, error)
		GetTinyUrl(url string) (string, error)
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

func (d demoService) TinyUrl(url string) (string, error) {
	log.Println("TinyUrl", url)
	md5Str := tinyurl.Md5Transform(url)
	tinyUrl := tinyurl.CalculateTinyUrlPrefix(md5Str)
	err := d.store.Demo().SetCache(tinyUrl, url, time.Hour*24*30)
	if err != nil {
		return "", err
	}
	return conf.Domain + tinyUrl, nil
}

func (d demoService) GetTinyUrl(url string) (string, error) {
	res, err := d.store.Demo().GetCache(url)
	if err != nil {
		log.Println("error： get cache", err)
		return "", nil
	}
	return res, err
}
