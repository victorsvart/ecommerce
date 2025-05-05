package userrepository

import (
	"context"

	"github.com/victorsvart/egommerce/internal/core/domain"
)

func (u *userRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	result := u.db.WithContext(ctx).Delete(&domain.User{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
