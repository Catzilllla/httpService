FROM golang:1.24.3-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o app ./ipocalc/cmd/server

EXPOSE 8080

CMD ["./app"]

# # --- Этап сборки (builder) ---
# # используем 1.21 (1.24.3 не существует)
# FROM golang:1.24.3-alpine AS builder  
# WORKDIR /app

# # Копируем зависимости и скачиваем их
# COPY go.mod go.sum ./
# RUN go mod tidy

# # Копируем исходный код и собираем бинарник
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app ./ipocalc/cmd/server

# # --- Этап запуска (финальный образ) ---
# FROM alpine:3.19
# WORKDIR /app

# # Копируем бинарник из builder-этапа
# COPY --from=builder /app/app .
# COPY --from=builder /app.

# # Уменьшаем размер Alpine (можно удалить кэш apk, если он не нужен)
# # RUN apk add --no-cache tzdata && \
# #     rm -rf /var/cache/apk/*

# EXPOSE 8080
# CMD ["./app"]