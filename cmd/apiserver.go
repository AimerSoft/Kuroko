package main

import (
	apiserver "kuroko/internal"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
	apiserver.NewAPIServer().Run()

}
