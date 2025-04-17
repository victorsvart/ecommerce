package domain

import "context"

type Product struct {
	ID       uint64
	Name     string
	ImageURL string
	UserID   uint64
}

type ProductRepository interface {
	Get(context.Context, uint64) (*Product, error)
	Create(context.Context, *Product) error
	Update(context.Context, *Product) error
}

type ProductUsecases interface {
	Get(context.Context, uint64) (*Product, error)
	Create(context.Context, *Product) error
	Update(context.Context, *Product) error
}

type ProductInput struct {
	ID       *uint64 `json:"id"`
	Name     string  `json:"name"`
	ImageURL string  `json:"imageUrl"`
	UserID   uint64  `json:"userId"`
}

func (p *ProductInput) ToProduct() Product {
	id := uint64(0)
	if p.ID != nil {
		id = *p.ID
	}

	return Product{
		ID:       id,
		Name:     p.Name,
		ImageURL: p.ImageURL,
		UserID:   p.UserID,
	}
}
