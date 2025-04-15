package application

import "github.com/victorsvart/go-ecommerce/internal/user/domain"

type userUseCaseImpl struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCases {
	return &userUseCaseImpl{repo}
}
