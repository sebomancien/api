package main

import (
	"fmt"
	"net/http"

	"github.com/sebomancien/api/logger"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s APIServer) Start() error {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	logger.LogInfo("API server started at ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
