package productusecase

import (
	"context"

	"github.com/victorsvart/egommerce/internal/core/domain"
)

func (p *productUseCase) Get(ctx context.Context, id uint64) (*domain.Product, error) {
	product, err := p.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
