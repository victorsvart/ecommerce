package productusecase

import (
	"context"

	"github.com/victorsvart/egommerce/internal/core/domain"
)

func (p *productUseCase) GetAll(ctx context.Context) ([]domain.Product, error) {
	return p.repo.GetAll(ctx)
}
