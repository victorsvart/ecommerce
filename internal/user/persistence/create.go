package persistence

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/user/domain"
)

func (u *userRepositoryImpl) Create(ctx context.Context, user *domain.User) error {
	if u.emailInUse(ctx, nil, user.Email) {
		return ErrEmailInUse
	}

	if err := u.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}

	return nil
}
