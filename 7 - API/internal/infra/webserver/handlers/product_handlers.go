package handlers

import (
	"encoding/json"
	"net/http"
	"renebizelli/api/internal/dtos"
	"renebizelli/api/internal/entities"
	"renebizelli/api/internal/infra/database"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var dto dtos.CreateProductInput

	e := json.NewDecoder(r.Body).Decode(&dto)

	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, e := entities.NewProduct(dto.Name, dto.Price)

	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e = ph.ProductDB.Create(product)

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description  Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Param page path string true "page number"
// @Param limit path string true "limit number"
// @Param sort path string true "sort by"
// @Success 200 {array} entities.Product
// @Failure 500
// @Failure 400
// @Router /products/{page}/{limit}/{sort} [get]
// @security ApiKeyAuth
func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	page := chi.URLParam(r, "page")
	limit := chi.URLParam(r, "limit")
	sort := chi.URLParam(r, "sort")

	if page == "" || limit == "" || sort == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pageInt, e := strconv.Atoi(page)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limitInt, e := strconv.Atoi(limit)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	products, e := ph.ProductDB.FindAll(pageInt, limitInt, sort)

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, e := ph.ProductDB.FindById(id)

	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto dtos.UpdateProductInput

	e := json.NewDecoder(r.Body).Decode(&dto)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	product, e := ph.ProductDB.FindById(id)
	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	product.Name = dto.Name
	product.Price = dto.Price

	e = ph.ProductDB.Update(product)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, e := ph.ProductDB.FindById(id)
	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	e = ph.ProductDB.Delete(id)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
