package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/idalgo-2021/example_swag/internal/models"
)

type ItemRepository interface {
	CreateItem(name string) (string, error)
	GetItemByID(id string) (*models.Item, error)
	GetAllItems() ([]*models.Item, error)
}

type itemRepo struct {
	items map[string]models.Item
}

func NewItemRepository() ItemRepository {
	return &itemRepo{
		items: map[string]models.Item{
			"1": {ID: "1", Name: "Table", CreateTime: time.Now()},
			"2": {ID: "2", Name: "Lamp", CreateTime: time.Now()},
			"3": {ID: "3", Name: "Chair", CreateTime: time.Now()},
		},
	}
}

func (r *itemRepo) CreateItem(name string) (string, error) {
	id := fmt.Sprintf("%d", len(r.items)+1)
	item := models.Item{
		ID:         id,
		Name:       name,
		CreateTime: time.Now(),
	}
	r.items[id] = item
	return id, nil
}

func (r *itemRepo) GetItemByID(id string) (*models.Item, error) {
	item, exists := r.items[id]
	if !exists {
		return nil, errors.New("item not found")
	}
	return &item, nil
}

func (r *itemRepo) GetAllItems() ([]*models.Item, error) {

	if len(r.items) == 0 {
		return nil, errors.New("no items available")
	}

	var allItems []*models.Item
	for _, item := range r.items {
		allItems = append(allItems, &item)
	}
	return allItems, nil
}
