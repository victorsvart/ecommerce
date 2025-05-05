package userrepository

import (
	"context"

	"github.com/victorsvart/egommerce/internal/core/domain"
)

func (u *userRepositoryImpl) GetById(ctx context.Context, id uint64) (*domain.User, error) {
	var user domain.User
	if err := u.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
