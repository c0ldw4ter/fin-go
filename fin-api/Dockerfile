# Этап сборки
FROM golang:1.23.4 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o fin-api ./cmd/fin-api/

# Этап запуска
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/fin-api .
CMD ["./fin-api"]