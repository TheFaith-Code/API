package data

import (
	"sync"

	"API/models"
)

var (
	mu    sync.RWMutex
	items = []models.Item{
		{ID: 1, Name: "Alpha", Description: "First demo item"},
		{ID: 2, Name: "Beta", Description: "Second demo item"},
		{ID: 3, Name: "Gamma", Description: "Third demo item"},
	}
)

// GetItems returns a copy of the current items slice (read-safe).
func GetItems() []models.Item {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]models.Item, len(items))
	copy(out, items)
	return out
}

// GetItemByID returns the item with the given id and a boolean whether it was found.
func GetItemByID(id int) (models.Item, bool) {
	mu.RLock()
	defer mu.RUnlock()
	for _, it := range items {
		if it.ID == id {
			return it, true
		}
	}
	return models.Item{}, false
}
