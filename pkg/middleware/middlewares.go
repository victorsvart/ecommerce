package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/victorsvart/egommerce/pkg/appcontext"
	"github.com/victorsvart/egommerce/pkg/rbac"
	"github.com/victorsvart/egommerce/pkg/token"
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
			UserID: claims.UserID,
			RoleID: claims.RoleID,
		}
		ctx := context.WithValue(r.Context(), appcontext.AuthCtxKey, authCtx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Permission(perm string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authVal := r.Context().Value(appcontext.AuthCtxKey)
			auth, ok := authVal.(appcontext.AuthContext)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			hasPerm := rbac.HasPermission(auth.RoleID, perm)
			if !hasPerm {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
