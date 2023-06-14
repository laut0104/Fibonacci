package main

import (
	"net/http"

	"github.com/laut0104/Fibonacci/handler"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/fib", handler.Fib)
	server.ListenAndServe()
}
