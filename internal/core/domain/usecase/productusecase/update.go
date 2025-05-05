package productusecase

import (
	"context"

	"github.com/victorsvart/egommerce/internal/core/domain"
)

func (u *productUseCase) Update(ctx context.Context, product *domain.Product) error {
	if err := u.repo.Update(ctx, product); err != nil {
		return err
	}

	return nil
}
