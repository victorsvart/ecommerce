package persistence

import (
	"context"
	"errors"

	"github.com/victorsvart/go-ecommerce/internal/user/domain"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{db}
}

func (u *userRepositoryImpl) emailInUse(ctx context.Context, id *uint64, address string) bool {
	tx := u.db.WithContext(ctx).
		Model(&domain.User{}).
		Select("1").
		Where("email = ?", address)

	if id != nil {
		tx = tx.Not("id = ?", *id)
	}

	rows, err := tx.Limit(1).Rows()
	if err != nil {
		return false
	}

	defer rows.Close()
	return rows.Next()
}

var (
	ErrEmailInUse   = errors.New("email is in use")
	ErrUserNotFound = errors.New("user not found")
	ErrIDRequired   = errors.New("id is required")
)
