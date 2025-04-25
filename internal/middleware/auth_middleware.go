package middleware

import (
    "encoding/json"
    "net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var tokenData map[string]string
        err := json.NewDecoder(r.Body).Decode(&tokenData)
        if err != nil || tokenData["token"] == "" {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Проверка токена (можно добавить логику валидации JWT)
        if tokenData["token"] != "valid-token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next(w, r)
    }
}
