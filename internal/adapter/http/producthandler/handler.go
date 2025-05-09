package producthandler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/victorsvart/egommerce/internal/core/domain"
	"github.com/victorsvart/egommerce/pkg/middleware"
	"github.com/victorsvart/egommerce/pkg/rbac"
	"github.com/victorsvart/egommerce/pkg/utils"
)

type ProductHandler struct {
	usecases domain.ProductUsecases
}

func NewProductHandler(api chi.Router, usecases domain.ProductUsecases) {
	handler := ProductHandler{usecases}
	api.With(middleware.Auth).Route("/products", func(r chi.Router) {
		r.With(middleware.Permission(rbac.GetProduct)).Get("/{id}", handler.GetById)
		r.With(middleware.Permission(rbac.GetProduct)).Get("/", handler.GetAll)
		r.With(middleware.Permission(rbac.GetProduct)).Get("/{userId}", handler.GetByUserId)
		r.With(middleware.Permission(rbac.CreateProduct)).Post("/", handler.CreateProducts)
		r.With(middleware.Permission(rbac.UpdateProduct)).Put("/", handler.UpdateProducts)
		r.With(middleware.Permission(rbac.DeleteProduct)).Delete("/{id}", handler.Delete)
	})
}

func (p *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	if idPath == "" {
		utils.RespondJSON(w, http.StatusBadRequest, false, errors.New("id is required"))
		return
	}

	id, err := strconv.ParseUint(idPath, 10, 64)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, errors.New("error getting id"))
		return
	}

	product, err := p.usecases.Get(r.Context(), id)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, true, ToProductPresenter(product))
}

func (p *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	filterText := r.URL.Query().Get("filterText")
	log.Println(filterText)
	products, err := p.usecases.GetAll(r.Context(), filterText)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, true, ToProductPresenterSlice(products))
}

func (p *ProductHandler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("userId")
	if idPath == "" {
		utils.RespondJSON(w, http.StatusBadRequest, false, errors.New("id is required"))
		return
	}

	id, err := strconv.ParseUint(idPath, 10, 64)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, errors.New("error getting id"))
		return
	}

	product, err := p.usecases.GetByUserID(r.Context(), id)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, true, ToProductPresenter(product))
}

func (p *ProductHandler) CreateProducts(w http.ResponseWriter, r *http.Request) {
	var input domain.ProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	}

	product := input.ToProduct()
	if err := p.usecases.Create(r.Context(), &product); err != nil {
		utils.RespondJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, true, ToProductPresenter(&product))
}

func (p *ProductHandler) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	var input domain.ProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondJSON(w, http.StatusUnprocessableEntity, false, err.Error())
		return
	}

	product := input.ToProduct()
	if err := p.usecases.Update(r.Context(), &product); err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, true, ToProductPresenter(&product))
}

func (p *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	if idPath == "" {
		utils.RespondJSON(w, http.StatusBadRequest, false, errors.New("id is required"))
		return
	}

	id, err := strconv.ParseUint(idPath, 10, 64)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, errors.New("error getting id"))
		return
	}

	if err := p.usecases.Delete(r.Context(), id); err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, false, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, true, "Product deleted successfully")
}
