package productrepository

import (
	"context"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
)

func (p *productRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	tx := p.db.Delete(&domain.Product{}, id)
	if err := tx.Error; err != nil {
		return err
	}

	if tx.RowsAffected == 0 {
		return ErrProductNotFound
	}

	return nil
}
