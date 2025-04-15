package wiring

import (
	"github.com/go-chi/chi/v5"
	"github.com/victorsvart/go-ecommerce/internal/adapter/http/userhandler"
	"github.com/victorsvart/go-ecommerce/internal/adapter/postgres/userrepository"
	"github.com/victorsvart/go-ecommerce/internal/authentication"
	"github.com/victorsvart/go-ecommerce/internal/core/domain/usecase/userusecase"
	"gorm.io/gorm"
)

func WireApp(db *gorm.DB, chi chi.Router) {
	wireUserAndAuth(db, chi)
}

func wireUserAndAuth(db *gorm.DB, api chi.Router) {
	repo := userrepository.NewUserRepository(db)
	usecases := userusecase.NewUserUseCase(repo)
	userhandler.NewUserHandler(api, usecases)
	authentication.NewAuthHandler(api, usecases)
}
