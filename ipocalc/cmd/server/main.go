package main


import {
    "fmt",
    "net/http",
    "log",
    "os",
    "io/ioutil",
    "gopkg.in/yaml.v3"
}

type Config struct {
    Server struct {
        Port int    `yaml:"port"`
        Host string `yaml:"host"`
    } `yaml:"server"`
}

func main() {
    // read config
	configfile := "/configs/config.yml" // Fixed variable declaration
	config, err := readConfig(configfile)
	if err != nil {
		log.Fatal(err)
	}
	if config.Server.Port == 0 {
		config.Server.Port = 8080
	}
	if config.Server.Host == "" {
		config.Server.Host = "0.0.0.0"
	}

	// // регистрация обработчиков
	// http.HandleFunc("/api", handleAPI)

    // //  запуск сервера
    // addr, err := 
    // http.ListenAndServe()
    fmt.Println(config.Server.Host)
	fmt.Println(config.Server.Port)
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
	return &cfg, nil
}
