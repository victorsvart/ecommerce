package wiring

import (
	"github.com/go-chi/chi/v5"
	"github.com/victorsvart/go-ecommerce/internal/authentication"
	"github.com/victorsvart/go-ecommerce/internal/user/application"
	"github.com/victorsvart/go-ecommerce/internal/user/handler"
	"github.com/victorsvart/go-ecommerce/internal/user/persistence"
	"gorm.io/gorm"
)

func WireApp(db *gorm.DB, chi chi.Router) {
	wireUserAndAuth(db, chi)
}

func wireUserAndAuth(db *gorm.DB, api chi.Router) {
	repo := persistence.NewUserRepository(db)
	usecases := application.NewUserUseCase(repo)
	handler.NewUserHandler(api, usecases)
	authentication.NewAuthHandler(api, usecases)
}
