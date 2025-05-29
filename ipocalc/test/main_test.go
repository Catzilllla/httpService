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
	"time"
)

func TestHandleCache(t *testing.T) {

	cacheStore := cachemod.NewContainer(5*time.Minute, 10*time.Minute)

	tests := []struct {
		name           string
		requestData    models.JsonRequest
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "No program selected",
			requestData: models.JsonRequest{
				ObjectCost:     5000000,
				InitialPayment: 1000000,
				Months:         240,
				Program:        models.JsonProgram{},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error": "choose program"}`,
		},
		{
			name: "More than one program selected",
			requestData: models.JsonRequest{
				ObjectCost:     5000000,
				InitialPayment: 1000000,
				Months:         240,
				Program:        models.JsonProgram{Salary: true, Military: true},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error": "choose only 1 program"}`,
		},
		{
			name: "Initial payment too low",
			requestData: models.JsonRequest{
				ObjectCost:     5000000,
				InitialPayment: 99000, // меньше 20% от стоимости
				Months:         240,
				Program:        models.JsonProgram{Salary: true},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error": "the initial payment should be more"}`,
		},
		{
			name: "Valid request",
			requestData: models.JsonRequest{
				ObjectCost:     5000000,
				InitialPayment: 1000000,
				Months:         240,
				Program:        models.JsonProgram{Salary: true},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":0,"params":{"object_cost":5000000,"initial_payment":1000000,"months":240,"program":{"salary":true,"military":false,"base":false}},"program":{"salary":true,"military":false,"base":false},"aggregates":{"rate":8,"loan_sum":4000000,"monthly_payment":2666666.6666666665,"overpayment":636000000,"last_payment_date":"29-05-2045 11:12:30"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/* получаем байты JSON'a */
			requestBody, err := json.Marshal(tt.requestData)
			if err != nil {
				t.Fatalf("could not marshal request data: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/execute", bytes.NewReader(requestBody))
			rr := httptest.NewRecorder()

			handlers.HandleExecute(rr, req, cacheStore)

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
