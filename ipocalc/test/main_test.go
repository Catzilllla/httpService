package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExecuteAndCache(t *testing.T) {
	// Инициализация
	cache.result = nil

	// Создаем сервер
	mux := http.NewServeMux()
	mux.Handle("/execute", LoggingMiddleware(http.HandlerFunc(HandleExecute)))
	mux.Handle("/cache", LoggingMiddleware(http.HandlerFunc(HandleCache)))

	ts := httptest.NewServer(mux)
	defer ts.Close()

	// Тест успешного запроса
	reqBody := `{
		"object_cost": 5000000,
		"initial_payment": 1000000,
		"months": 240,
		"program": {"salary": true}
	}`

	resp, err := http.Post(ts.URL+"/execute", "application/json", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	// Проверка /cache
	resp2, err := http.Get(ts.URL + "/cache")
	if err != nil {
		t.Fatal(err)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp2.StatusCode)
	}
	var cacheRes []Result
	if err := json.NewDecoder(resp2.Body).Decode(&cacheRes); err != nil {
		t.Fatal(err)
	}
	if len(cacheRes) != 1 {
		t.Fatalf("expected 1 result, got %d", len(cacheRes))
	}
}
