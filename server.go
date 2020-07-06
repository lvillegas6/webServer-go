package main

import (
	"fmt"
	"net/http"
)

//Server ...
type Server struct {
	port   string
	router *Router
}

//NewServer ...
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {

	_, exist := s.router.rules[path]

	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.rules[path][method] = handler
}

//AddMiddleware ...
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		f = middleware(f) //Middleware retorna un handler
	}
	return f
}

//Listen ...
func (s *Server) Listen() error {
	http.Handle("/", s.router)
	fmt.Println("Server on port:", s.port)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
