BINARY_NAME=app
SOURCE_DIR=.
DOCKERFILE_DIR=.
DOCKER_IMAGE_NAME=ipocalc
PORT=8080
.DEFAULT_GOAL := help

build:
	@echo "Сборка бинарного файла..."
	go build -o $(BINARY_NAME) ./ipocalc/cmd/server

test:
	@echo "Запуск тестов..."
	go test ./... -cover

docker-build:
	@echo "Сборка Docker-образа..."
	docker build -t $(DOCKER_IMAGE_NAME) $(DOCKERFILE_DIR)

docker-run:
	@echo "Запуск Docker-контейнера..."
	docker run -p $(PORT):8080 --rm $(DOCKER_IMAGE_NAME)

docker-stop:
	@echo "Остановка всех работающих контейнеров..."
	docker stop $(shell docker ps -q)

lint:
	@echo "Проверка линтера..."
	golangci-lint run

clean:
	@echo "Очистка..."
	rm -f $(BINARY_NAME)

help:
	@echo "Доступные цели:"
	@echo "  build       - Сборка бинарного файла"
	@echo "  test        - Запуск тестов"
	@echo "  docker-build- Сборка Docker-образа"
	@echo "  docker-run  - Запуск контейнера Docker"
	@echo "  docker-stop - Остановка всех работающих контейнеров"
	@echo "  clean       - Удаление бинарного файла"



