package domain

import "context"

type Product struct {
	ID                 uint64
	Name               string
	ImageURL           string
	Price              float64
	DiscountPercentage *uint16
	Description        string
	UserID             uint64
}

type ProductRepository interface {
	Get(context.Context, uint64) (*Product, error)
	GetAll(context.Context, string) ([]Product, error)
	GetByUserID(ctx context.Context, id uint64) (*Product, error)
	Create(context.Context, *Product) error
	Update(context.Context, *Product) error
	Delete(context.Context, uint64) error
}

type ProductUsecases interface {
	Get(context.Context, uint64) (*Product, error)
	GetAll(context.Context, string) ([]Product, error)
	GetByUserID(ctx context.Context, id uint64) (*Product, error)
	Create(context.Context, *Product) error
	Update(context.Context, *Product) error
	Delete(context.Context, uint64) error
}

type ProductInput struct {
	ID                 *uint64 `json:"id"`
	Name               string  `json:"name"`
	Price              float64 `json:"price"`
	DiscountPercentage *uint16 `json:"discountPercentage"`
	ImageURL           string  `json:"imageUrl"`
	Description        string  `json:"description"`
	UserID             uint64  `json:"userId"`
}

func (p *ProductInput) ToProduct() Product {
	id := uint64(0)
	if p.ID != nil {
		id = *p.ID
	}

	return Product{
		ID:                 id,
		Name:               p.Name,
		ImageURL:           p.ImageURL,
		Price:              p.Price,
		DiscountPercentage: p.DiscountPercentage,
		Description:        p.Description,
		UserID:             p.UserID,
	}
}
