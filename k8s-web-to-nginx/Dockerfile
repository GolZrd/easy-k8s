# Этап, на котором выполняется сборка приложения
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o /main main.go

# Финальный этап, копируем собранное приложение
FROM alpine:3 

COPY --from=builder main /main

ENTRYPOINT ["/main"]