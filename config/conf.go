package conf

import (
	"kuroko/internal/pkg/util"
	"log"
	"os"
)

var (
	Port        = "8080"
	Mode        = "debug"
	ProjectName = "Kuroko"
	RedisHost   = ""
	RedisPwd    = ""
	Domain      = ""
	//Domain = "http://127.0.0.1:8080/tinyurl/"
)

func LoadConfig() {
	Domain = "http://" + util.GetIP() + ":" + Port + "/k/"
	Domain = os.Getenv("DOMAIN")
	RedisHost = os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	RedisPwd = os.Getenv("REDIS_PWD")
	log.Println("load config", Domain, Port, RedisHost, RedisPwd)
}
