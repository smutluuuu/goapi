package utils

import "net/http"

type Middleware func(http.Handler) http.Handler

//Middleware is a function that wraps an http.Handler with additional functionality
func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
