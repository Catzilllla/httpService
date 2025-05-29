FROM golang:1.24.3-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o app ./ipocalc/cmd/server

FROM alpine:latest
WORKDIR /mapp
COPY --from=builder /app .
COPY ./ipocalc/configs/config.yml /app/config.yml
EXPOSE 8080

CMD ["./app"]
