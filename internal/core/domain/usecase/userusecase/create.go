package userusecase

import (
	"context"

	"github.com/victorsvart/egommerce/internal/core/domain"
	"github.com/victorsvart/egommerce/pkg/utils"
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
