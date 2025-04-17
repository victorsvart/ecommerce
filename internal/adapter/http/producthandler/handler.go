package producthandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/victorsvart/go-ecommerce/internal/core/domain"
	"github.com/victorsvart/go-ecommerce/pkg/utils"
)

type ProductHandler struct {
	usecases domain.ProductUsecases
}

func NewProductHandler(api chi.Router, usecases domain.ProductUsecases) {
	handler := ProductHandler{usecases}
	api.Route("/products", func(r chi.Router) {
		r.Get("/{id}", handler.GetById)
		r.Post("/", handler.CreateProducts)
		r.Put("/", handler.UpdateProducts)
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

	utils.RespondJSON(w, http.StatusOK, true, product)
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

	utils.RespondJSON(w, http.StatusOK, true, product)
}
