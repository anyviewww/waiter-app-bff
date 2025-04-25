package main

import (
    "log"
    "net/http"

    "waiter-app-bff/internal/config"
    "waiter-app-bff/internal/handler"
    "waiter-app-bff/internal/middleware"
)

func main() {
    // Загрузка конфигурации
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Инициализация gRPC клиентов
    grpcClient, err := handler.NewGRPCClient(cfg.GRPCServerAddress)
    if err != nil {
        log.Fatalf("Failed to initialize gRPC client: %v", err)
    }

    // Создание обработчиков
    h := handler.NewHandler(grpcClient)

    // Настройка маршрутов
    mux := http.NewServeMux()
    mux.HandleFunc("/api/orders", middleware.AuthMiddleware(h.HandleOrders))

    // Запуск HTTP сервера
    log.Printf("Starting server on %s", cfg.ServerAddress)
    if err := http.ListenAndServe(cfg.ServerAddress, mux); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
