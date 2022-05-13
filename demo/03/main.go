package main

import (
	"github.com/chenzebinm4/webframe"
	"net/http"
)

func main() {
	core := webframe.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
