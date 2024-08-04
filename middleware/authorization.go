package middleware

import (
	"context"
	"net/http"

	"github.com/luqxus/spaces/tokens"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signedToken := r.Header.Get("authorization")
		if signedToken == "" {
			http.Error(w, "authorization header not provided", http.StatusUnauthorized)
			return
		}

		uid, err := tokens.VerifyJwt(signedToken)
		if err != nil {
			http.Error(w, "invalid authorization header", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "uid", uid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
