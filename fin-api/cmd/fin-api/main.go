package main

import (
	"context"
	"fin-api/internal/config"
	"fin-api/internal/handler"
	"fin-api/internal/repository"
	"fin-api/internal/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
    // Загрузка конфигурации
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v\n", err)
    }

    // Формирование строки подключения к базе данных
    connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBSSLMode, cfg.DBPassword)

    // Подключение к базе данных
    dbpool, err := pgxpool.New(context.Background(), connStr)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer dbpool.Close()

    // Инициализация слоев приложения
    repo := repository.NewRepository(dbpool)
    srv := service.NewService(repo)
    hdl := handler.NewHandler(srv)

    // Настройка маршрутизатора
    r := gin.Default()

    // Middleware для логирования и обработки ошибок
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // Группировка маршрутов
    api := r.Group("/api/v1")
    {
        api.POST("/users/:user_id/topup", hdl.TopUpBalance)
        api.POST("/transactions/transfer", hdl.TransferMoney) // Изменен путь
        api.GET("/users/:user_id/transactions", hdl.GetLastTransactions)
    }

    // Запуск сервера
    log.Println("Starting server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v\n", err)
    }
}