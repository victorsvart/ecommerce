package persistence

import (
	"context"
	"errors"

	"github.com/victorsvart/go-ecommerce/internal/user/domain"
	"gorm.io/gorm"
)

func (u *userRepositoryImpl) List(ctx context.Context) ([]domain.User, error) {
	users := make([]domain.User, 0)
	if err := u.db.Find(&users).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return users, err
	}

	return users, nil
}
