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
	urlPrefix := tinyurl.CalculateTinyUrlPrefix(md5Str, 0)
	// 使用计算出的key查询是否有冲突
	cache, err := d.store.Demo().GetCache(urlPrefix)
	if err != nil {
		log.Println("redis error", err)
	}
	if len(cache) != 0 {
		// 即该url已计算过
		if cache == url {
			return conf.Domain + urlPrefix, nil
		} else {
			// 发生key冲突
			oldPreFix := urlPrefix
			log.Println("pre hash error", "发生前缀hash冲突")
			randFix := uint64(1)
			for true {
				if urlPrefix != oldPreFix {
					urlPrefix = tinyurl.CalculateTinyUrlPrefix(md5Str, randFix)
					randFix++
				}
			}
		}
	}
	err = d.store.Demo().SetCache(urlPrefix, url, time.Hour*24*30)
	if err != nil {
		return "", err
	}
	return conf.Domain + urlPrefix, nil
}

func (d demoService) GetTinyUrl(url string) (string, error) {
	res, err := d.store.Demo().GetCache(url)
	if err != nil {
		log.Println("error： get cache", err)
		return "", nil
	}
	return res, err
}
