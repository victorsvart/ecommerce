package userrepository

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
	"gorm.io/gorm/clause"
)

func (u *userRepositoryImpl) Update(ctx context.Context, user *domain.User) error {
	if u.emailInUse(ctx, &user.ID, user.Email) {
		return ErrEmailInUse
	}

	tx := u.db.Model(&domain.User{}).
		Omit("password").
		Where("id = ?", user.ID).
		Clauses(clause.Returning{}).
		Updates(user).
		Scan(&user)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
