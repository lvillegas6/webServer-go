package main

import (
	"fmt"
	"net/http"
)

//Server ...
type Server struct {
	port string
}

//NewServer ...
func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

//Listen ...
func (s *Server) Listen() error {
	fmt.Println("Server on port:", s.port)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
