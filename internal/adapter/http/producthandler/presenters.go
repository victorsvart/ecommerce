package producthandler

import (
	"github.com/victorsvart/egommerce/internal/core/domain"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type ProductPresenter struct {
	ID                 uint64  `json:"id"`
	Name               string  `json:"name"`
	ImageURL           string  `json:"imageUrl"`
	Price              string  `json:"price"`
	DiscountPercentage *uint16 `json:"discountPercentage"`
	Description        string  `json:"description"`
	UserID             uint64  `json:"userId"`
}

func ToProductPresenter(p *domain.Product) ProductPresenter {
	price := p.Price
	if p.DiscountPercentage != nil {
		price = price * (1 - float64(*p.DiscountPercentage)/100)
	}

	return ProductPresenter{
		ID:                 p.ID,
		Name:               p.Name,
		ImageURL:           p.ImageURL,
		Price:              formatBRL(price),
		DiscountPercentage: p.DiscountPercentage,
		Description:        p.Description,
		UserID:             p.UserID,
	}
}

func ToProductPresenterSlice(p []domain.Product) []ProductPresenter {
	s := make([]ProductPresenter, 0)
	for i := range p {
		s = append(s, ToProductPresenter(&p[i]))
	}

	return s
}

func formatBRL(value float64) string {
	p := message.NewPrinter(language.BrazilianPortuguese)
	sf := p.Sprintf("R$ %.2f", value)
	return sf
}
