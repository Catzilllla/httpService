package handlers

import (
	"net/http"
)

func handleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func HandleExecute(w http.ResponseWriter, r *http.Request) {
	// Тут будет обработка /execute
	w.Write([]byte("execute"))
}

func HandleCache(w http.ResponseWriter, r *http.Request) {
	// Тут будет обработка /cache
	w.Write([]byte("cache"))
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Можно добавить логирование
		next.ServeHTTP(w, r)
	})
}
