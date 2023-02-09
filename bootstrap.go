package wss

import (
	"log"
	"net/http"
)

type Bootstrap struct {
	addr string

	// Is it a reused body
	reusedBody bool
}

type Options struct {
	// Addr eg :8080
	Addr string

	// Is it a reused body
	ReusedBody bool
}

// New function can be used to create a startup
func New(opt *Options) *Bootstrap {
	return &Bootstrap{
		reusedBody: opt.ReusedBody,
	}
}

func (b *Bootstrap) AddResponder(resp Responder) {
	initResponder(resp, b)
}

func (b *Bootstrap) Bind(addr string) {
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Panicln(err)
		return
	}
}
