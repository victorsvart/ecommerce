package productusecase

import (
	"context"

	"github.com/victorsvart/egommerce/internal/core/domain"
)

func (p *productUseCase) GetByUserID(ctx context.Context, id uint64) (*domain.Product, error) {
	product, err := p.repo.GetByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
