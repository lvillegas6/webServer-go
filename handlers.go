package main

import (
	"fmt"
	"net/http"
)

//HandleRoot ...
func HandleRoot(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "Hola mundo")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the very best server!")
}
