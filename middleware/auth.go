package middleware

import (
	"fmt"
	"milliy/auth"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
)

var (
	publicPaths = map[string]map[string]bool{
		"/v1/user/login":        {"POST": true},
		"/v1/twit/{id}":         {"GET": true}, // Allow GET only
		"/v1/twits":             {"GET": true},
		"/v1/twits/type/{type}": {"GET": true},
		"/v1/twits/most-viewed": {"GET": true},
		"/v1/twits/latest":      {"GET": true},
		"/v1/twits/search":      {"GET": true},
		"/v1/twit-count/{id}":   {"POST": true},
	}
)

func isPublicPath(path, method string) bool {
	// Check if the path exists in the publicPaths map
	if methods, exists := publicPaths[path]; exists {
		if methods[method] {
			return true
		}
	}
	// Additional check for Swagger paths
	return strings.HasPrefix(path, "/swagger/")
}

func AuthMiddleware(enforcer *casbin.Enforcer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if the path and method are public
			if isPublicPath(r.URL.Path, r.Method) {
				next.ServeHTTP(w, r)
				return
			}

			// Get token
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			_, role, err := auth.GetUserIdFromRefreshToken(authHeader)
			if err != nil {
				http.Error(w, "Authorization error", http.StatusInternalServerError)
				return
			}

			// Check permission
			allowed, err := enforcer.Enforce(role, r.URL.Path, r.Method)
			fmt.Println(role, r.URL.Path, r.Method)
			if err != nil {
				http.Error(w, "Authorization error", http.StatusInternalServerError)
				return
			}
			if !allowed {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
