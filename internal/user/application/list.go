package application

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/user/domain"
)

func (u *userUseCaseImpl) List(ctx context.Context) ([]domain.User, error) {
	return u.repo.List(ctx)
}
