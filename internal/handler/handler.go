package handler

import (
    "encoding/json"
    "net/http"

    "waiter-app-bff/internal/grpc_client"
    pb "waiter-app-bff/proto"
)

type Handler struct {
    grpcClient *grpc_client.GRPCClient
}

func NewHandler(grpcClient *grpc_client.GRPCClient) *Handler {
    return &Handler{grpcClient: grpcClient}
}

func (h *Handler) HandleOrders(w http.ResponseWriter, r *http.Request) {
    // Получение данных из запроса
    var orderRequest pb.OrderRequest
    err := json.NewDecoder(r.Body).Decode(&orderRequest)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Вызов gRPC сервиса
    ctx := r.Context()
    orderResponse, err := h.grpcClient.GetOrders(ctx, &orderRequest)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Отправка ответа
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orderResponse)
}
