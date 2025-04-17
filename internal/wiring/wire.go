package wiring

import (
	"github.com/go-chi/chi/v5"
	"github.com/victorsvart/go-ecommerce/internal/adapter/http/producthandler"
	"github.com/victorsvart/go-ecommerce/internal/adapter/http/userhandler"
	"github.com/victorsvart/go-ecommerce/internal/adapter/postgres/productrepository"
	"github.com/victorsvart/go-ecommerce/internal/adapter/postgres/userrepository"
	"github.com/victorsvart/go-ecommerce/internal/authentication"
	"github.com/victorsvart/go-ecommerce/internal/core/domain/usecase/productusecase"
	"github.com/victorsvart/go-ecommerce/internal/core/domain/usecase/userusecase"
	"gorm.io/gorm"
)

func WireApp(db *gorm.DB, api chi.Router) {
	wireUserAndAuth(db, api)
	wireProduct(db, api)
}

func wireUserAndAuth(db *gorm.DB, api chi.Router) {
	repo := userrepository.NewUserRepository(db)
	usecases := userusecase.NewUserUseCase(repo)
	userhandler.NewUserHandler(api, usecases)
	authentication.NewAuthHandler(api, usecases)
}

func wireProduct(db *gorm.DB, api chi.Router) {
	repo := productrepository.NewProductRepository(db)
	usecases := productusecase.NewProductUseCase(repo)
	producthandler.NewProductHandler(api, usecases)
}
