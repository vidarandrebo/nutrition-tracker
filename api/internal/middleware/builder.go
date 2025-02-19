package middleware

import "net/http"

type Builder struct {
	mw []Middleware
}

func NewMiddlewareBuilder() *Builder {
	return &Builder{mw: make([]Middleware, 0)}
}

type Middleware func(next http.Handler) http.Handler

func (mwb *Builder) AddMiddleware(fn Middleware) {
	mwb.mw = append(mwb.mw, fn)
}

func (mwb *Builder) Build() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		for _, fn := range mwb.mw {
			next = fn(next)
		}
		return next
	}
}
