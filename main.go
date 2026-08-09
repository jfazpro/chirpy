package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{Handler: mux, Addr: ":8080"}

	http.ListenAndServe(server.Addr, server.Handler)
}
