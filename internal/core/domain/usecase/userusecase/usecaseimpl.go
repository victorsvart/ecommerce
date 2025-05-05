package userusecase

import "github.com/victorsvart/egommerce/internal/core/domain"

type userUseCaseImpl struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCases {
	return &userUseCaseImpl{repo}
}
