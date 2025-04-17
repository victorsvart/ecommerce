package productusecase

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
)

func (p *productUseCase) GetAll(ctx context.Context) ([]domain.Product, error) {
	return p.repo.GetAll(ctx)
}
