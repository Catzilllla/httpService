package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Catzilllla/httpService/ipocalc/configs"

	cachemod "github.com/Catzilllla/httpService/ipocalc/internal/cache"
	"github.com/Catzilllla/httpService/ipocalc/internal/handlers"
)

func main() {
	// cache initialising
	// здесь проводим инициализацию TTL
	//
	// Инициализируем кэш с TTL (например, 5 минут для кэшируемых данных и 10 минут для очистки устаревших данных)
	cacheStore := cachemod.NewCacheStore(5*time.Minute, 10*time.Minute)

	// read config
	pwd, _ := os.Getwd()
	fmt.Println("Current working directory:", pwd)
	configPath := filepath.Join(pwd, "ipocalc", "configs", "config.yml")

	fmt.Println(configPath)

	config, err := configs.ReadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	// addr := fmt.Sprintf("http://%s:%d", "0.0.0.0", "8080")
	// reg heandlers

	http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleExecute(w, r, cacheStore)
	})
	http.HandleFunc("/cache", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleCache(w, r, cacheStore)
	})
	http.ListenAndServe(addr, nil)
}
