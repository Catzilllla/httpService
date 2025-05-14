# httpService
Http json service deposit calculator

# структура
ipocalc/
├── cmd/
│   └── server/
│       └── main.go             # точка входа в приложение
├── internal/
│   ├── handlers/               # HTTP-обработчики
│   │   └── handlers.go
│   ├── services/               # бизнес-логика, расчет ипотеки
│   │   └── mortgage.go
│   ├── cache/                  # кэш в памяти
│   │   └── cache.go
│   └── models/                 # модели данных
│       └── models.go
├── configs/
│   └── config.yml              # конфигурационный файл
├── scripts/                      # скрипты для сборки/развертывания
│   └── ...
├── test/                         # тестовые файлы и тесты
│   └── handlers_test.go
├── Dockerfile
├── Makefile
├── go.mod
├── go.sum
└── README.md



Общая структура решения
    Основной файл: main.go — содержит сервер, маршруты, middleware, бизнес-логику.
    Конфигурация: config.yml — порт.
    Кэш: реализован в памяти (глобальный []Result с мьютексом).
    Валидация: проверка программ, суммы, ошибок.
    Логика расчетов: расчет ставки, платежа, переплаты, даты.
    API: два обработчика (/execute, /cache).
    Middleware: логирование запросов.
    Тесты: в файле main_test.go.
    Docker: Dockerfile.
    Makefile: для команд.

Итоги
    Вы получите полноценный сервис с REST API, кэшем, middleware, конфигурацией.
    Можно запускать через Docker.
    Можно писать юнит-тесты.
    Все зависимости — "завендорены" (используйте go mod).


Инструкции по запуску
    Создайте папку ipocalc/
    Поместите туда все файлы
    В терминале перейдите в папку ipocalc/
    Инициализируйте модуль:
        go mod init github.com/yourusername/ipocalc
        go mod tidy

    Постройте образ:
        docker build -t ipocalc .

    Запустите контейнер:
        docker run -d -p 8080:8080 --name ipocalc-container ipocalc
