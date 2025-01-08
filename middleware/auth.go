package middleware

import (
	"context"
	"milliy/auth"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
)

func AuthMiddleware(enforcer *casbin.Enforcer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip auth for login and swagger
			if r.URL.Path == "/v1/user/login" || strings.HasPrefix(r.URL.Path, "/swagger/") {
				next.ServeHTTP(w, r)
				return
			}

			// Get token from header
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Validate token and get user role
			_, role, err := auth.GetUserIdFromRefreshToken(token)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Check permission
			allowed, err := enforcer.Enforce(role, r.URL.Path, r.Method)
			if err != nil {
				http.Error(w, "Authorization error", http.StatusInternalServerError)
				return
			}

			if !allowed {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			// Add user info to context
			ctx := context.WithValue(r.Context(), "user_role", role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
