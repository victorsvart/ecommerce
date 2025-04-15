package application

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/user/domain"
)

func (u *userUseCaseImpl) Update(ctx context.Context, user *domain.User) error {
	if err := u.repo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
