package middleware

import "net/http"

type Middleware struct{
	handler http.Handler
}

type MiddlewareFunc func(http.Handler) http.Handler
func SetHandler(handler http.Handler) *Middleware {
	return &Middleware{
		handler: handler,
	}
}

func (mw *Middleware) Chain(middlewares ...MiddlewareFunc) http.Handler {
	for _,middleware:=range middlewares{
		mw.handler = middleware(mw.handler)
	}
	return mw.handler
}