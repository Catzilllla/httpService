package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type Program struct {
	Salary   bool `json:"salary"`
	Military bool `json:"military"`
	Base     bool `json:"base"`
}

type RequestDataPost struct {
	Object_cost     float64 `json:"object_cost"`
	Initial_payment float64 `json:"initial_payment"`
	Months          int     `json:"months"`
	Program         Program `json:"program"`
}

func main() {
	// case 1 - GET | case 2 - POST (request.json)
	var input string
	fmt.Print("case 1: - GET: /cache\n")
	fmt.Print("case 2: - POST (request.json) /execute\n")

	fmt.Scan(&input)

	// read config
	apiUrl := fmt.Sprintf("http://%s:%d/api", "0.0.0.0", 8080)

	// read json for POST
	pwd, _ := os.Getwd()
	jsonPath := filepath.Join(pwd, "request.json")

	var addresForJsonData RequestDataPost
	dataFromJsonFile, err := os.ReadFile(jsonPath)

	if err != nil {
		fmt.Println("problem read files. Error: ", err)
	}
	if err := yaml.Unmarshal(dataFromJsonFile, &addresForJsonData); err != nil {
		fmt.Println("YAML error: %w", err)
	}

	// Обрабатываем ввод с помощью switch case
	switch input {
	case "1":
		// GET
		fmt.Println("start GET request ... ")
		getGetResponse, err := sendGetRequest(apiUrl)

		if err != nil {
			fmt.Println("Error of GET resquest!", err)
			return
		}
		fmt.Println("response: ", getGetResponse)
	case "2":
		// POST
		fmt.Println("start POST request ... ")
		getPostResponse, err := sendPostRequest(apiUrl, addresForJsonData)
		if err != nil {
			fmt.Println("Ошибка POST:", err)
			return
		}
		fmt.Println("Ответ сервера (POST):", getPostResponse)

	default:
		fmt.Printf("0: %s\n", input)
	}
}

func sendGetRequest(url string) (string, error) {
	// client settings
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	//  sending request
	// HTTP-ответы в Go используют системные ресурсы (сетевые соединения, файловые
	// дескрипторы).
	// Если не закрыть Body, может произойти утечка ресурсов (например, исчерпание
	// лимита открытых соединений).
	// defer откладывает выполнение функции до момента, когда завершится текущая
	// функция (в данном случае — main() или та, где был сделан HTTP-запрос).

	// send request
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("problem URL. Error: %v", err)
	}
	defer resp.Body.Close()

	// take response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("problem read response. Error: %v", err)
	}
	return string(body), nil
}

// Отправка POST-запроса с JSON
func sendPostRequest(url string, data interface{}) (string, error) {
	// Сериализация данных в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("ошибка кодирования JSON: %v", err)
	}

	// Создание запроса
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("ошибка создания запроса: %v", err)
	}

	// Установка заголовков
	req.Header.Set("Content-Type", "application/json")

	// Настройка HTTP-клиента с таймаутом
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Отправка запроса
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ответа: %v", err)
	}

	return string(body), nil
}
