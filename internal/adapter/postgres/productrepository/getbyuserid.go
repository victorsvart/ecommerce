package productrepository

import (
	"context"
	"errors"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
	"gorm.io/gorm"
)

func (p *productRepositoryImpl) GetByUserID(ctx context.Context, id uint64) (*domain.Product, error) {
	var product domain.Product
	tx := p.db.Where("user_id", id).
		Find(&product)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, tx.Error
	}

	return &product, nil
}
