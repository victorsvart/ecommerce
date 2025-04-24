package userusecase

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
)

func (u *userUseCaseImpl) GetById(ctx context.Context, id uint64) (*domain.User, error) {
	user, err := u.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
