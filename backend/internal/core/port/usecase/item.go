package usecase

import "hackernews-clone-backend/internal/core/domain"

type ItemService interface {
	GetItem(id int) (*domain.Item, error)
	ListTopStories() ([]domain.Item, error)
	ListNewStories() ([]domain.Item, error)
	ListAskStories() ([]domain.Item, error)
	ListShowStories() ([]domain.Item, error)
	ListJobStories() ([]domain.Item, error)
	// Add other use cases like submitting stories, comments, etc.
}
