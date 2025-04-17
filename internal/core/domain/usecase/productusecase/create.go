package productusecase

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
)

func (u *productUseCase) Create(ctx context.Context, product *domain.Product) error {
	if err := u.repo.Create(ctx, product); err != nil {
		return err
	}

	return nil
}
