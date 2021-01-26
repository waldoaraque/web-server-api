package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true // booleano para bloquear middleware
			fmt.Println("Middleware: Checking Authentication...")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { //defer permite ejecutar una funcion al final del scope
				log.Println(r.URL.Path, time.Since(start))
			}() //llamado a la funcion anonima

			f(w, r)
		}
	}
}
