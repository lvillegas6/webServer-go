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

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	fmt.Println("Test")
	s.router.rules[path] = handler
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
