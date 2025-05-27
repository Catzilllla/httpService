package main

import (
	"fmt"
	"ipocalc/ipocalc/configs"
	cachemod "ipocalc/ipocalc/internal/cache"
	"ipocalc/ipocalc/internal/handlers"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// sync MAP (время истечения срока действия элементов и интервал очистки)
	cacheStore := cachemod.NewContainer(5*time.Minute, 10*time.Minute)

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

	// Регистрируем обработчик с передачей cacheStore через замыкание
	http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleExecute(w, r, cacheStore)
	})
	http.HandleFunc("/cache", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleCache(w, r, cacheStore)
	})
	http.ListenAndServe(addr, nil)
}
