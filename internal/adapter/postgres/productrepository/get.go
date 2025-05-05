package productrepository

import (
	"context"
	"errors"

	"github.com/victorsvart/egommerce/internal/core/domain"
	"gorm.io/gorm"
)

func (p *productRepositoryImpl) Get(ctx context.Context, id uint64) (*domain.Product, error) {
	var product domain.Product
	if err := p.db.Where("id = ?", id).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}

		return nil, err
	}

	return &product, nil
}
