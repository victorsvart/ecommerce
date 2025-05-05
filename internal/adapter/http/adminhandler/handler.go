package adminhandler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/victorsvart/egommerce/internal/core/domain"
	"github.com/victorsvart/egommerce/pkg/middleware"
	"github.com/victorsvart/egommerce/pkg/rbac"
	"github.com/victorsvart/egommerce/pkg/utils"
)

type AdminHandler struct {
	userCases    domain.UserUseCases
	productCases domain.ProductUsecases
}

func NewAdminHandler(
	api chi.Router,
	userCases domain.UserUseCases,
	productCases domain.ProductUsecases,
) {
	handler := &AdminHandler{userCases, productCases}
	api.With(middleware.Auth).Route("/admin", func(r chi.Router) {
		r.With(middleware.Permission(rbac.ListUser)).Get("/", handler.List)
	})
}

func (a *AdminHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := a.userCases.List(r.Context())
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, true, ToAdminUserPresenterSlice(users))
}
