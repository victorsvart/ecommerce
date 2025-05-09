package productusecase

import (
	"context"

	"github.com/victorsvart/egommerce/internal/core/domain"
)

func (p *productUseCase) GetAll(ctx context.Context, filterText string) ([]domain.Product, error) {
	return p.repo.GetAll(ctx, filterText)
}
