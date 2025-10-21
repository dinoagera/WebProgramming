package auth

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type contextKey string

const UserIDKey contextKey = "userID"

func extractBearerToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	return splitToken[1]
}
func New(log *slog.Logger, secretKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := extractBearerToken(r)
			if tokenStr == "" {
				log.Info("authorization token missing")
				http.Error(w, "Authorization token required", http.StatusUnauthorized)
				return
			}
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			}
			token, err := jwt.Parse(tokenStr, keyFunc)
			if err != nil {
				log.Info("token invalid", "err", err)
				http.Error(w, "Authorization token required", http.StatusUnauthorized)
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				log.Info("claim invalid", "err", err)
				http.Error(w, "Authorization token required", http.StatusUnauthorized)
				return
			}
			uid, _ := claims["uid"].(float64)
			ctx := context.WithValue(r.Context(), UserIDKey, int64(uid))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey).(int64)
	return userID, ok
}
