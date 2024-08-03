package middleware

import (
	"net/http"

	"github.com/sebomancien/api/logger"
)

func Log(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.LogInfo(r.Method, r.RequestURI)
		handler.ServeHTTP(w, r)
	}
}
