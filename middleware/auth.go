package middleware

import (
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenParts := strings.Split(r.Header.Get("Authorization"), " ")
		if len(tokenParts) < 2 || tokenParts[1] == "" {
			http.Error(w, "Not authenticated", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
