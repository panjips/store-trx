package middleware

import (
	"context"
	"net/http"
	"store-trx-go/internal/handler/responses"
	"store-trx-go/pkg/utils"
	"strings"
)

type contextKey string

const (
    UserIDKey contextKey = "userID"
    EmailKey  contextKey = "email"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			responses.HTTPResponse(w, "error", http.StatusUnauthorized, "Missing token", nil)
			return
		}

		requestToken := strings.TrimPrefix(tokenHeader, "Bearer ")


		claims, err := utils.ValidateToken(requestToken)

		if err != nil {
			responses.HTTPResponse(w, "error", http.StatusUnauthorized, "Invalid token", nil)
			return
		}
		
		ctx := context.WithValue(r.Context(), UserIDKey, claims.DataClaims.ID)
		ctx = context.WithValue(ctx, EmailKey, claims.DataClaims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}