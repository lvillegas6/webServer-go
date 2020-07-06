package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flag := true
		fmt.Println("Checking authentication")
		time.Sleep(2 * time.Second)
		if flag {
			f(w, r)
		} else {
			return
		}
	}
}

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Println(r.URL.Path, time.Since(start))
		}()
		f(w, r)
	}
}
