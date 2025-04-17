package productrepository

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
)

func (p *productRepositoryImpl) GetAll(ctx context.Context) ([]domain.Product, error) {
	products := make([]domain.Product, 0)
	if err := p.db.Model(&domain.Product{}).Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}
