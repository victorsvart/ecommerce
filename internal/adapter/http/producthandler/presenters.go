package producthandler

import "github.com/victorsvart/egommerce/internal/core/domain"

type ProductPresenter struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
	UserID   uint64 `json:"userId"`
}

func ToProductPresenter(p *domain.Product) ProductPresenter {
	return ProductPresenter{
		ID:       p.ID,
		Name:     p.Name,
		ImageURL: p.ImageURL,
		UserID:   p.UserID,
	}
}

func ToProductPresenterSlice(p []domain.Product) []ProductPresenter {
	s := make([]ProductPresenter, 0)
	for i := range p {
		s = append(s, ToProductPresenter(&p[i]))
	}

	return s
}
