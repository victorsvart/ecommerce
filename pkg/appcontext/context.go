package appcontext

import (
	"context"
	"net/http"
)

type AuthContext struct {
	UserID uint64
	// RoleID uint64
}

type ctxKey string

const AuthCtxKey ctxKey = "auth"

func GetAuthContextFromRequest(r *http.Request) (AuthContext, bool) {
	auth, ok := r.Context().Value(AuthCtxKey).(AuthContext)
	return auth, ok
}

func GetAuthContext(ctx context.Context) (AuthContext, bool) {
	auth, ok := ctx.Value(AuthCtxKey).(AuthContext)
	return auth, ok
}
