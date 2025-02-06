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
	AdminKey contextKey = "isAdmin"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			responses.HTTPResponse(w, "error", http.StatusUnauthorized, "missing token", nil)
			return
		}

		requestToken := strings.TrimPrefix(tokenHeader, "Bearer ")
		claims, err := utils.ValidateToken(requestToken)
		if err != nil {
			responses.HTTPResponse(w, "error", http.StatusUnauthorized, "invalid token", nil)
			return
		}
		
		ctx := context.WithValue(r.Context(), UserIDKey, claims.DataClaims.ID)
		ctx = context.WithValue(ctx, EmailKey, claims.DataClaims.Email)
		ctx = context.WithValue(ctx, AdminKey, claims.DataClaims.Admin)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthenticationAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		var isAdmin bool = r.Context().Value(AdminKey).(bool)

		if !isAdmin {
			responses.HTTPResponse(w, "error", http.StatusUnauthorized, "unauthorized action", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}