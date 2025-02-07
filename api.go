package main

import (
	"fmt"
	"net/http"

	"github.com/sebomancien/api/logger"
	"github.com/sebomancien/api/middleware"
	"github.com/sebomancien/api/storage"
)

type APIServer struct {
	addr  string
	store storage.Store
}

func NewAPIServer(addr string, store *storage.Store) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

func (s APIServer) Start() error {
	router := http.NewServeMux()

	router.Handle("/", middleware.Log(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	}))

	logger.LogInfo("API server started at ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
