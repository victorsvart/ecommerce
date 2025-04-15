package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/victorsvart/go-ecommerce/pkg/appcontext"
	"github.com/victorsvart/go-ecommerce/pkg/token"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := token.ParseJWT(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		userIDFloat, err := strconv.ParseUint(claims.RegisteredClaims.ID, 10, 64)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid user ID in token", http.StatusUnauthorized)
			return
		}

		authCtx := appcontext.AuthContext{
			UserID: uint64(userIDFloat),
			//	RoleID: uint64(roleIDFloat),
		}
		ctx := context.WithValue(r.Context(), appcontext.AuthCtxKey, authCtx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
