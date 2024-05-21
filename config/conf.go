package conf

import (
	"log"
	"os"
)

var (
	Port        = "8080"
	Mode        = "debug"
	ProjectName = "Kuroko"
	RedisHost   = "101.42.18.21:6379"
	RedisPwd    = "umbink"
	Domain      = "https://umb.ink/k/"
	//Domain = "http://127.0.0.1:8080/tinyurl/"
)

func LoadConfig() {
	Domain = os.Getenv("DOMAIN")
	Port = os.Getenv("PORT")
	RedisHost = os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	RedisPwd = os.Getenv("REDIS_PWD")
	log.Println("load config", Domain, Port, RedisHost, RedisPwd)
}
