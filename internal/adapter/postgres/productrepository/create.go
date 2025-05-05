package productrepository

import (
	"context"
	"errors"

	"github.com/victorsvart/egommerce/internal/core/domain"
	"gorm.io/gorm"
)

func (p *productRepositoryImpl) Create(ctx context.Context, product *domain.Product) error {
	if err := p.userExists(ctx, product.UserID); err != nil {
		return err
	}

	if err := p.db.WithContext(ctx).Create(product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductNotFound
		}

		return err
	}

	return nil
}
