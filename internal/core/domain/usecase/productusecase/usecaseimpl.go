package productusecase

import (
	"github.com/victorsvart/go-ecommerce/internal/core/domain"
)

type productUseCase struct {
	repo domain.ProductRepository
}

func NewProductUseCase(repo domain.ProductRepository) domain.ProductUsecases {
	return &productUseCase{repo}
}
