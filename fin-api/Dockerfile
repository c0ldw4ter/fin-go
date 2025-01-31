# Этап установки Goose
FROM golang:1.23.4 AS goose-installer
WORKDIR /goose
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.11.0

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

# Установка совместимости с glibc
RUN apk add --no-cache bash curl gcompat

# Копируем Goose из первого этапа
COPY --from=goose-installer /go/bin/goose /usr/local/bin/

# Копируем бинарник приложения
COPY --from=builder /app/fin-api .

# Копируем директорию migrations
COPY migrations /migrations

CMD ["./fin-api"]