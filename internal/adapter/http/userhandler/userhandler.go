package userhandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/victorsvart/go-ecommerce/internal/core/domain"
	"github.com/victorsvart/go-ecommerce/pkg/middleware"
	"github.com/victorsvart/go-ecommerce/pkg/rbac"
	"github.com/victorsvart/go-ecommerce/pkg/utils"
)

var (
	ErrInvalidPathValue = errors.New("invalid path value parameter")
)

type UserHandler struct {
	usecases domain.UserUseCases
}

func NewUserHandler(api chi.Router, usecases domain.UserUseCases) {
	handler := &UserHandler{usecases}
	api.With(middleware.Auth).Route("/users", func(r chi.Router) {
		r.With(middleware.Permission(rbac.GetUser)).Get("/", handler.List)
		r.With(middleware.Permission(rbac.CreateUser)).Post("/", handler.Create)
		r.With(middleware.Permission(rbac.UpdateUser)).Put("/", handler.Update)
		r.With(middleware.Permission(rbac.DeleteUser)).Delete("/{id}", handler.Delete)
	})
}

func (u *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := u.usecases.List(r.Context())
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, true, ToUserPresenterSlice(users))
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.UserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	}

	user := input.ToUser()
	if err := u.usecases.Create(r.Context(), &user); err != nil {
		utils.RespondJSON(w, http.StatusConflict, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, true, ToUserPresenter(&user))
}

func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input domain.UserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	}

	user := input.ToUser()
	if err := u.usecases.Update(r.Context(), &user); err != nil {
		utils.RespondJSON(w, http.StatusBadGateway, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, true, ToUserPresenter(&user))
}

func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	if idPath == "" {
		utils.RespondJSON(w, http.StatusBadRequest, false, ErrInvalidPathValue)
		return
	}

	id, err := strconv.ParseUint(idPath, 10, 64)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	err = u.usecases.Delete(r.Context(), id)
	if err != nil {
		utils.RespondJSON(w, http.StatusNotFound, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, true, "Deleted successfully")
}
