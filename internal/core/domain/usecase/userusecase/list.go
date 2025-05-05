package userusecase

import (
	"context"

	"github.com/victorsvart/egommerce/internal/core/domain"
)

func (u *userUseCaseImpl) List(ctx context.Context) ([]domain.User, error) {
	return u.repo.List(ctx)
}
