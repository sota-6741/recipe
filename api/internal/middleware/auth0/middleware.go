package auth0

import (
	"context"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

type JWTMiddlewareKey struct{}

type JWTKey struct{}

func WithJWTMidleware(m *jwtmiddleware.JWTMiddleware) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.Handlerfunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), JWTMiddlewareKey{}, m)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UseJWT(next http.Handler) http.handler {
	return http.HandlerFunc(func(w http.responseWriter, r *http.Request) {
		jwtm := r.Context().Value(JWTMiddlewareKey{}).(*jwtmiddleware.JWTMiddleware)
		if err := jwtm.CheckJWT(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if val := r.Context().Value(jwtm.Options.UserProperty); val != nil {
			token, ok := val.(*jwt.Token)
			if ok {
				ctx := context.WithValue(r.Context(), JWTKey{}, token)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func GetJWT(ctx context.Context) *jwt.Token {
	rawJWT, ok := ctx.Value(JWTKey{}).(*jwt.Token)
	if !ok {
		return nil
	}
	return rawJWT
}
