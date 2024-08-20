package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/idalgo-2021/example_swag/internal/models"
	"github.com/idalgo-2021/example_swag/internal/repository"
)

type ItemHandler struct {
	repo repository.ItemRepository
}

func NewItemHandler(repo repository.ItemRepository) *ItemHandler {
	return &ItemHandler{repo: repo}
}

// @Summary Create a new item
// @Description Create a new item by providing a name
// @Tags items
// @Accept json
// @Produce json
// @Param body body models.ItemRequest true "Item details"
// @Success 201 {object} models.ItemResponse "ID of the created item"
// @Failure 400 {object} models.Error "Invalid input"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /items [post]
func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var requestBody models.ItemRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.repo.CreateItem(requestBody.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.ItemResponse{ID: id}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// @Summary Get an item by its ID
// @Description Get the item with the specified ID
// @Tags items
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} models.Item "Item description"
// @Failure 404 {object} models.Error "Item not found"
// @Router /items/{id} [get]
func (h *ItemHandler) GetItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	item, err := h.repo.GetItemByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

// @Summary Get all items
// @Description Get a list of all items
// @Tags items
// @Accept json
// @Produce json
// @Success 200 {array} models.Item "List of items"
// @Failure 404 {object} models.Error "No items available"
// @Router /items [get]
func (h *ItemHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.repo.GetAllItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(items)
}
