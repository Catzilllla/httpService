package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	cachemod "github.com/Catzilllla/httpService/ipocalc/internal/cache"
	"github.com/Catzilllla/httpService/ipocalc/internal/models"
	"github.com/Catzilllla/httpService/ipocalc/internal/services"
)

func HandleExecute(w http.ResponseWriter, r *http.Request, storeCache *cachemod.CacheStore) {
	// !!!
	// ПРОВЕРЯЯЕМ ВСЕ СЛУЧАИ
	// if err != nil {
	// 	if apiErr, ok := err.(*models.ApiError); ok {
	// 		switch apiErr.Message {
	// 		case "choose program":
	// 			w.WriteHeader(http.StatusBadRequest)
	// 			json.NewEncoder(w).Encode(map[string]string{"error": "choose program"})
	// 		case "choose only 1 program":
	// 			w.WriteHeader(http.StatusBadRequest)
	// 			json.NewEncoder(w).Encode(map[string]string{"error": "choose only 1 program"})
	// 		case "the initial payment should be more":
	// 			w.WriteHeader(http.StatusBadRequest)
	// 			json.NewEncoder(w).Encode(map[string]string{"error": "the initial payment should be more"})
	// 		default:
	// 			http.Error(w, `{"error":"calculation error"}`, http.StatusInternalServerError)
	// 		}
	// 	} else {
	// 		http.Error(w, `{"error":"calculation error"}`, http.StatusInternalServerError)
	// 	}
	// 	return
	// }
	// !!!

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// чтение и десерелиализация json из тела запроса
	var ourRequest models.JsonRequest
	if err := json.NewDecoder(r.Body).Decode(&ourRequest); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(ourRequest)
	if err != nil {
		log.Fatal(err)
	}

	cacheKey := cachemod.HashRequestBody(jsonData)
	// fmt.Println(cacheKey)
	// fmt.Println(jsonData)
	// fmt.Println(ourRequest)
	// fmt.Println("cacheKey: TYPE:", reflect.TypeOf(cacheKey))
	// fmt.Println("jsonData: TYPE:", reflect.TypeOf(jsonData))
	// fmt.Println("ourRequest: TYPE:", reflect.TypeOf(ourRequest))

	// Проверяем, есть ли уже результат в кэше
	var cachedResponse map[string]interface{}
	if err := storeCache.Get(cacheKey, &cachedResponse); err == nil {
		// Если есть кэшированный результат, возвращаем его
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(cachedResponse); err != nil {
			http.Error(w, "Не удалось вернуть кэшированные данные", http.StatusInternalServerError)
		}
		fmt.Println("CACHED")
		fmt.Println(reflect.TypeOf(cachedResponse))
		fmt.Println(cachedResponse)
		return
	}

	// если кэша нет то идем дальше
	// fmt.Println("no cache: go to backend and calculate this")

	aggregates, err := services.CalculateMortgage(ourRequest.Program, ourRequest)
	if err != nil {
		log.Fatal("back error: cant calculate")
	}
	storeCache.Set(cacheKey, aggregates, 10*time.Minute)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aggregates)
	fmt.Println("BACKEND")
	fmt.Println(reflect.TypeOf(aggregates))
	fmt.Println(aggregates)
}

func HandleCache(w http.ResponseWriter, r *http.Request, storeCache *cachemod.CacheStore) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("Only GET-requests are allowed"))
		return
	}

	allItems := storeCache.GetAll()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(allItems); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(allItems)
}

// func LoggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		next.ServeHTTP(w, r)
// 	})
// }
