package handlers

import (
	"encoding/json"
	"fmt"
	cachemod "ipocalc/ipocalc/internal/cache"
	"ipocalc/ipocalc/internal/models"
	"ipocalc/ipocalc/internal/services"
	"log"
	"net/http"
	"time"
)

// func httpRequestInformation(r *http.Request) {
// 	fmt.Printf("Method: %s\n", r.Method)
// 	fmt.Printf("URL: %s\n", r.URL)
// 	fmt.Println("Headers:")
// 	for key, values := range r.Header {
// 		for _, value := range values {
// 			fmt.Printf("%s: %s\n", key, value)
// 		}
// 	}

// 	if r.Body != nil {
// 		body, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			log.Println("Error reading body:", err)
// 		} else {
// 			fmt.Println("Body: ")
// 			fmt.Println(string(body))
// 			fmt.Println(" ")
// 		}
// 	}
// }

func HandleExecute(w http.ResponseWriter, r *http.Request, storeCache *cachemod.Cache) {
	/* инфа о запросе */
	// httpRequestInformation(r)

	/* проверка метода */
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	/* чтение и десерелиализация json из тела запроса */
	/*
		type JsonRequest struct {
			ObjectCost     float64     `json:"object_cost"`
			InitialPayment float64     `json:"initial_payment"`
			Months         int         `json:"months"`
			Program        JsonProgram `json:"program"`
		}
	*/
	var requestMap models.JsonRequest
	if err := json.NewDecoder(r.Body).Decode(&requestMap); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}

	/* обработка случаев */
	selectedCount := 0
	if requestMap.Program.Salary {
		selectedCount++
	}
	if requestMap.Program.Military {
		selectedCount++
	}
	if requestMap.Program.Base {
		selectedCount++
	}
	if selectedCount == 0 {
		http.Error(w, `{"error": "choose program"}`, http.StatusBadRequest)
		return
	}
	if selectedCount > 1 {
		http.Error(w, `{"error": "choose only 1 program"}`, http.StatusBadRequest)
		return
	}

	/* первоначальный взнос */
	if requestMap.InitialPayment < requestMap.ObjectCost*0.2 {
		http.Error(w, `{"error": "the initial payment should be more"}`, http.StatusBadRequest)
		return
	}

	/* requestMapa в  байты | потом в  HASH */
	jsonData, err := json.Marshal(requestMap)
	if err != nil {
		log.Fatal("Error marshaling data Json to bytes")
		return
	}
	/* jsonData в models.JsonRequest | потом в BACK если нужно */
	var req models.JsonRequest
	err = json.Unmarshal(jsonData, &req)
	if err != nil {
		log.Fatal("Error Unmarshaling data Json to JsonRequest")
		return
	}

	fmt.Println("ourRequest:   ", req)
	/* хэшируем и используем как ключ */
	keyCacheId := cachemod.HashRequestBody(req)

	/* ищем по ключу */
	if body, found := storeCache.Get(keyCacheId); found {
		fmt.Println("body from GET:  ", body)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(body); err != nil {
			http.Error(w, "Не удалось вернуть кэшированные данные", http.StatusInternalServerError)
		}
		fmt.Println("CACHED")
		return
		/*
			type JsResult struct {
				ID         int           `json:"id"`
				Params     JsonRequest   `json:"params"`
				Program    JsonProgram   `json:"program"`
				Aggregates JsonAggregate `json:"aggregates"`
			}
		*/
	}

	/* задействуем бэк если не нашли выше */
	newId := storeCache.Counter
	resultJson, err := services.CalculateMortgage(req, newId)
	if err != nil {
		log.Fatal("cant calculating")
	}
	storeCache.Counter++
	storeCache.Set(keyCacheId, resultJson, 5*time.Minute)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultJson)
	fmt.Println("BACKEND")
	// fmt.Println(reflect.TypeOf(calculating))
	// fmt.Println(calculating)
}

func HandleCache(w http.ResponseWriter, r *http.Request, storeCache *cachemod.Cache) {
	/* инфа о запросе */
	// httpRequestInformation(r)

	if r.Method != http.MethodGet {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	result := storeCache.GetAll()

	if len(result) == 0 {
		http.Error(w, `{"error": "empty cache"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode result", http.StatusInternalServerError)
	}
}
