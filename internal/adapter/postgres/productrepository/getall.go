package productrepository

import (
	"context"
	"strings"

	"github.com/victorsvart/egommerce/internal/core/domain"
)

func (p *productRepositoryImpl) GetAll(ctx context.Context, filterText string) ([]domain.Product, error) {
	filterText = strings.TrimSpace(filterText)
	products := make([]domain.Product, 0)
	tx := p.db.Model(&domain.Product{})

	if filterText != "" {
		tx = tx.Where("LOWER(name) LIKE LOWER(?)", "%"+filterText+"%")
	}

	if err := tx.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
