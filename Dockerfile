FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o app ./ipocalc/cmd/server

EXPOSE 8080

CMD ["./app"]
