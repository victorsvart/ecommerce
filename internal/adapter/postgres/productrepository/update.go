package productrepository

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
)

func (p *productRepositoryImpl) Update(ctx context.Context, product *domain.Product) error {
	if err := p.userExists(ctx, product.UserID); err != nil {
		return err
	}

	if err := p.canUserTakeAction(ctx, product.ID); err != nil {
		return err
	}

	tx := p.db.
		Where("id = ?", product.ID).
		Updates(product).
		Scan(&product)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}
