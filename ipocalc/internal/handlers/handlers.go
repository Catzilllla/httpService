package handlers

import (
	"net/http"
)

func HandleExecute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// 405
		http.Error(w, "Method now allowed", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("execute"))

	// нужно получить json и раскодировать
	aggregates, err := calculateMortgage()
}

func HandleCache(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("Only GET-requests.."))
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("cache"))
}

// func LoggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		next.ServeHTTP(w, r)
// 	})
// }

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("api"))
}
