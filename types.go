package main

import "net/http"

//Middleware ...
type Middleware func(http.HandlerFunc) http.HandlerFunc
