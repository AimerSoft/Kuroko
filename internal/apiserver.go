package apiserver

import (
	conf "kuroko/config"
	"kuroko/internal/app"
	"kuroko/internal/option"
)

func NewAPIServer() *app.App {
	conf.LoadConfig()
	opts := option.NewOptions()
	application := app.NewApp("Bus API Server",
		conf.ProjectName,
		app.WithOptions(opts),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *option.Options) app.RunFunc {
	return func(basename string) error {
		return nil
	}
}
