package productrepository

import (
	"context"
	"errors"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepositoryImpl{db}
}

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrProductNotFound = errors.New("product not found")
)

func (p *productRepositoryImpl) userExists(ctx context.Context, userID uint64) error {
	var exists bool
	err := p.db.WithContext(ctx).
		Model(&domain.User{}).
		Select("count(*) > 0").
		Where("id = ?", userID).
		Find(&exists).Error

	if err != nil {
		return err
	}

	if !exists {
		return ErrUserNotFound
	}

	return nil
}
