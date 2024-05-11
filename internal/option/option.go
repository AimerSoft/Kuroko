package option

import (
	"github.com/redis/go-redis/v9"
	"kuroko/config"
	"kuroko/internal/pkg/http"
)

type Options struct {
	WebOptions *http.WebOption
	//BotOptions *
	RedisOptions *redis.Client
}

func NewOptions() *Options {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.RedisHost,
		Password: "umbink", // no password set
		DB:       1,        // use default DB
	})
	redisClient.Options()
	//redisClient.Set(context.Background(), "test", "test", time.Hour)
	//res := redisClient.Get(context.Background(), "test")
	//log.Println(res)
	return &Options{
		WebOptions:   http.NewWebOptions(),
		RedisOptions: redisClient,
	}
}
