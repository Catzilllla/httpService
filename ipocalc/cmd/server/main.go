package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Catzilllla/httpService/ipocalc/internal/handlers"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

type jsonProgram struct {
	Salary   bool `json:"salary"`
	Military bool `json:"military"`
	Base     bool `json:"base"`
}

type jsonResult struct {
	ObjectCost     float64     `json:"object_cost"`
	InitialPayment float64     `json:"initial_payment"`
	Months         int         `json:"months"`
	Program        jsonProgram `json:"program"`
}

type jsonAggregate struct {
	Rate            float64 `json:"rate"`
	LoanSum         float64 `json:"loan_sum"`
	MonthlyPayment  float64 `json:"monthly_payment"`
	Overpayment     float64 `json:"overpayment"`
	LastPaymentDate string  `json:"last_payment_date"`
}

func main() {
	// read config
	pwd, _ := os.Getwd()
	fmt.Println("Current working directory:", pwd)
	configPath := filepath.Join(pwd, "ipocalc", "configs", "config.yml")

	fmt.Println(configPath)

	config, err := readConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	// reg heandlers
	http.HandleFunc("/api", handlers.HandleAPI)
	http.HandleFunc("/execute", handlers.HandleExecute)
	http.HandleFunc("/cache", handlers.HandleCache)
	http.ListenAndServe(addr, nil)
}

func readConfig(filename string) (*Config, error) {
	var cfg Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("YAML error: %w", err)
	}

	// default config
	if cfg.Server.Port == 0 {
		cfg.Server.Port = 8080
	}
	if cfg.Server.Host == "" {
		cfg.Server.Host = "0.0.0.0"
	}
	return &cfg, nil
}
