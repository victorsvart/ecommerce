package userrepository

import (
	"context"

	"github.com/victorsvart/egommerce/internal/core/domain"
	"github.com/victorsvart/egommerce/pkg/utils"
)

func (u *userRepositoryImpl) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, utils.CheckNotFoundErr(err, ErrUserNotFound)
	}

	return &user, nil
}
