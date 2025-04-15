package userusecase

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
	"github.com/victorsvart/go-ecommerce/pkg/utils"
)

func (u *userUseCaseImpl) Create(ctx context.Context, user *domain.User) error {
	err := utils.HashPassword(&user.Password)
	if err != nil {
		return err
	}

	err = u.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
