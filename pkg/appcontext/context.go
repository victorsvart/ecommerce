package appcontext

import (
	"context"
	"errors"
	"net/http"
)

type AuthContext struct {
	UserID uint64
	RoleID uint
}

type ctxKey string

const AuthCtxKey ctxKey = "auth"

func GetAuthContextFromRequest(r *http.Request) (AuthContext, bool) {
	auth, ok := r.Context().Value(AuthCtxKey).(AuthContext)
	return auth, ok
}

func GetAuthContext(ctx context.Context) (AuthContext, error) {
	auth, ok := ctx.Value(AuthCtxKey).(AuthContext)
	if !ok {
		return AuthContext{}, ErrContextFetch
	}
	return auth, nil
}

var (
	ErrContextFetch = errors.New("error fetching context")
)
