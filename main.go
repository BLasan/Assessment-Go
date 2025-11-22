package main

import (
	"net/http"
	"github.com/BLasan/Assessment-Go/handlers"
)

func main() {
	http.HandleFunc("/hello-world", handlers.HelloWorldHandler)

	port := ":8080"
	println("Server starting on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
