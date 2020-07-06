package main

import (
	"net/http"
)

//Router ...
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

//NewRouter ...
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

//FindHandler ...
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, exist, methodExist
}

//ServeHTTP ...
func (r *Router) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	handler, exist, methodExist := r.FindHandler(request.URL.Path, request.Method)

	if !exist {
		write.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		write.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(write, request)
}
