package wiring

import (
	"github.com/go-chi/chi/v5"
	"github.com/victorsvart/go-ecommerce/internal/adapter/http/adminhandler"
	"github.com/victorsvart/go-ecommerce/internal/adapter/http/producthandler"
	"github.com/victorsvart/go-ecommerce/internal/adapter/http/userhandler"
	"github.com/victorsvart/go-ecommerce/internal/adapter/postgres/productrepository"
	"github.com/victorsvart/go-ecommerce/internal/adapter/postgres/userrepository"
	"github.com/victorsvart/go-ecommerce/internal/authentication"
	"github.com/victorsvart/go-ecommerce/internal/core/domain"
	"github.com/victorsvart/go-ecommerce/internal/core/domain/usecase/productusecase"
	"github.com/victorsvart/go-ecommerce/internal/core/domain/usecase/userusecase"
	"gorm.io/gorm"
)

func WireApp(db *gorm.DB, api chi.Router) {
	user := wireUserAndAuth(db, api)
	product := wireProduct(db, api)

	adminhandler.NewAdminHandler(api, user, product)
}

func wireUserAndAuth(db *gorm.DB, api chi.Router) domain.UserUseCases {
	repo := userrepository.NewUserRepository(db)
	usecases := userusecase.NewUserUseCase(repo)
	userhandler.NewUserHandler(api, usecases)
	authentication.NewAuthHandler(api, usecases)
	return usecases
}

func wireProduct(db *gorm.DB, api chi.Router) domain.ProductUsecases {
	repo := productrepository.NewProductRepository(db)
	usecases := productusecase.NewProductUseCase(repo)
	producthandler.NewProductHandler(api, usecases)
	return usecases
}
