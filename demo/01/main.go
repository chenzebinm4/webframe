package main

import (
	"github.com/chenzebinm4/webframe"
	"net/http"
)

func main() {
	server := &http.Server{
		Handler: webframe.NewCore(),
		Addr:    "localhost:8080",
	}
	server.ListenAndServe()
}
