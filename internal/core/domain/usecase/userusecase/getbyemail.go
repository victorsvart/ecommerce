package userusecase

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
)

func (u *userUseCaseImpl) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
