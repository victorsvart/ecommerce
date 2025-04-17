package productrepository

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
)

func (p *productRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	if err := p.canUserTakeAction(ctx, id); err != nil {
		return err
	}

	tx := p.db.Delete(&domain.Product{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}
