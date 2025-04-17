package repository

import "hackernews-clone-backend/internal/core/domain"

type ItemRepository interface {
	Get(id int) (*domain.Item, error)
	GetAllByType(itemType domain.ItemType, limit int) ([]domain.Item, error)
	// Save(item *domain.Item) error
	// Update(item *domain.Item) error
}
