package productrepository

import (
	"context"
	"errors"

	"github.com/victorsvart/go-ecommerce/internal/core/domain"
	"github.com/victorsvart/go-ecommerce/pkg/appcontext"
	"github.com/victorsvart/go-ecommerce/pkg/rbac"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepositoryImpl{db}
}

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrProductNotFound       = errors.New("product not found")
	ErrUserIsNotProductOwner = errors.New("user is not product owner")
)

func (p *productRepositoryImpl) userExists(ctx context.Context, userID uint64) error {
	var exists bool
	err := p.db.WithContext(ctx).
		Model(&domain.User{}).
		Select("count(*) > 0").
		Where("id = ?", userID).
		Find(&exists).Error

	if err != nil {
		return err
	}

	if !exists {
		return ErrUserNotFound
	}

	return nil
}

func (p *productRepositoryImpl) canUserTakeAction(ctx context.Context, id uint64) error {
	authCtx, err := appcontext.GetAuthContext(ctx)
	if err != nil {
		return err
	}

	var ownerID uint64
	err = p.db.
		Model(&domain.Product{}).
		Select("user_id").
		Where("id = ? AND user_id = ?", id, authCtx.UserID).
		Pluck("user_id", &ownerID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}

		return err
	}

	var canDelete = (ownerID == authCtx.UserID) || authCtx.RoleID == rbac.AdminRoleID
	if !canDelete {
		return ErrUserIsNotProductOwner
	}
	return nil
}
