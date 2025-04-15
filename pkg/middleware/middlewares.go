package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/victorsvart/go-ecommerce/pkg/appcontext"
	"github.com/victorsvart/go-ecommerce/pkg/token"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		claims, err := token.ParseJWT(cookie.Value)
		if err != nil {
			http.Error(w, ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		authCtx := appcontext.AuthContext{
			UserID: uint64(claims.UserID),
			//	RoleID: uint64(roleIDFloat),
		}
		ctx := context.WithValue(r.Context(), appcontext.AuthCtxKey, authCtx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
