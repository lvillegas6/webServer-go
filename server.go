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
