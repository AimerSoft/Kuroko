package option

import "kuroko/internal/pkg/http"

type Options struct {
	WebOptions *http.WebOption
	//BotOptions *
}

func NewOptions() *Options {
	return &Options{
		WebOptions: http.NewWebOptions(),
	}
}
