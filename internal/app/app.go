package app

import (
	"context"
	"kuroko/internal/option"
	"kuroko/internal/server"
	"kuroko/pkg/shutdown"
	"log"
	"net/http"
	"time"
)

type App struct {
	basename    string
	name        string
	description string
	options     *option.Options
	runFunc     RunFunc
}

// RunFunc ...
type RunFunc func(basename string) error

// Option ...
type Option func(*App)

func NewApp(name, basename string, opts ...Option) *App {
	a := &App{
		name:     name,
		basename: basename,
	}
	for _, o := range opts {
		o(a)
	}
	return a
}
func WithOptions(opt *option.Options) Option {
	return func(a *App) {
		a.options = opt
	}
}

// WithRunFunc ...
func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}

// Run ...
func (a *App) Run() {
	serv, err := server.NewServer(a.options)
	if err != nil {
		panic(err)
	}
	log.Println("")
	a.options.WebOptions.Server.Handler = serv.Web
	if err := a.options.WebOptions.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("http server startup err", err)
	}

	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := a.options.WebOptions.Shutdown(ctx); err != nil {
				log.Println("server shutdown err", err)
			}
		},
	)
}
