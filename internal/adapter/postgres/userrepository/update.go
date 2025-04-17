package userrepository

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
)

func (u *userRepositoryImpl) Update(ctx context.Context, user *domain.User) error {
	if u.emailInUse(ctx, &user.ID, user.Email) {
		return ErrEmailInUse
	}

	tx := u.db.Model(&domain.User{}).
		Omit("password").
		Where("id = ?", user.ID).
		Updates(user).
		Scan(&user)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
