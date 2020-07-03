package main

import (
	"net/http"
)

//Router ...
type Router struct {
	rules map[string]http.HandlerFunc
}

//NewRouter ...
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

//FindHandler ...
func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}

//ServeHTTP ...
func (r *Router) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	handler, exist := r.FindHandler(request.URL.Path)

	if !exist {
		write.WriteHeader(http.StatusNotFound)
		return
	}
	handler(write, request)
}
