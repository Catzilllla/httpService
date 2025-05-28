package test

import (
	"bytes"
	"encoding/json"
	cachemod "ipocalc/ipocalc/internal/cache"
	"ipocalc/ipocalc/internal/handlers"
	"ipocalc/ipocalc/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCache(t *testing.T) {
	tests := []struct {
		name           string
		requestData    models.JsonRequest
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "No program selected",
			requestData: models.JsonRequest{
				ObjectCost:     100000,
				InitialPayment: 25000,
				Months:         12,
				Program:        models.JsonProgram{},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error": "choose program"}`,
		},
		{
			name: "More than one program selected",
			requestData: models.JsonRequest{
				ObjectCost:     100000,
				InitialPayment: 25000,
				Months:         12,
				Program:        models.JsonProgram{Salary: true, Military: true},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error": "choose only 1 program"}`,
		},
		{
			name: "Initial payment too low",
			requestData: models.JsonRequest{
				ObjectCost:     100000,
				InitialPayment: 15000, // меньше 20% от стоимости
				Months:         12,
				Program:        models.JsonProgram{Salary: true},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error": "the initial payment should be more"}`,
		},
		{
			name: "Valid request",
			requestData: models.JsonRequest{
				ObjectCost:     100000,
				InitialPayment: 30000,
				Months:         12,
				Program:        models.JsonProgram{Salary: true},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"params":{"object_cost":100000,"initial_payment":30000,"months":12,"program":{"salary":true,"military":false,"base":false}},"program":{"salary":true,"military":false,"base":false},"aggregates":{"rate":0.05,"loan_sum":70000,"monthly_payment":2916.6666666666665,"overpayment":50000,"last_payment_date":"2026-05-28"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Преобразуем данные в JSON
			requestBody, err := json.Marshal(tt.requestData)
			if err != nil {
				t.Fatalf("could not marshal request data: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/handle", bytes.NewReader(requestBody))
			rr := httptest.NewRecorder()

			// Вызываем обработчик
			handlers.HandleCache(rr, req, &cachemod.Cache{})

			// Проверяем статус-код
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, status)
			}

			// Проверяем тело ответа
			if body := rr.Body.String(); body != tt.expectedBody {
				t.Errorf("expected body %s, got %s", tt.expectedBody, body)
			}
		})
	}
}
