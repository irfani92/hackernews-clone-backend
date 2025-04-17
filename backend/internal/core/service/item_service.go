package service

import (
	"context"
	"hackernews-clone-backend/internal/core/domain"
)

type HNAPIAdapter interface {
	GetItem(ctx context.Context, id int) (domain.Item, error)
	GetTopStories(ctx context.Context) ([]int, error)
	GetNewStories(ctx context.Context) ([]int, error)
	GetAskStories(ctx context.Context) ([]int, error)
	GetShowStories(ctx context.Context) ([]int, error)
	GetJobStories(ctx context.Context) ([]int, error)
}

type ItemService struct {
	hnAPIAdapter HNAPIAdapter
}

func NewItemService(hnAPIAdapter HNAPIAdapter) *ItemService {
	return &ItemService{hnAPIAdapter: hnAPIAdapter}
}

func (s *ItemService) GetItem(ctx context.Context, id int) (domain.Item, error) {
	return s.hnAPIAdapter.GetItem(ctx, id)
}

func (s *ItemService) ListTopStories(ctx context.Context, page int, pageSize int) ([]domain.Item, error) {
	ids, err := s.hnAPIAdapter.GetTopStories(ctx)
	if err != nil {
		return nil, err
	}
	return s.fetchPaginatedItems(ctx, ids, page, pageSize)
}

func (s *ItemService) ListNewStories(ctx context.Context, page int, pageSize int) ([]domain.Item, error) {
	ids, err := s.hnAPIAdapter.GetNewStories(ctx)
	if err != nil {
		return nil, err
	}
	return s.fetchPaginatedItems(ctx, ids, page, pageSize)
}

func (s *ItemService) ListAskStories(ctx context.Context, page int, pageSize int) ([]domain.Item, error) {
	ids, err := s.hnAPIAdapter.GetAskStories(ctx)
	if err != nil {
		return nil, err
	}
	return s.fetchPaginatedItems(ctx, ids, page, pageSize)
}

func (s *ItemService) ListShowStories(ctx context.Context, page int, pageSize int) ([]domain.Item, error) {
	ids, err := s.hnAPIAdapter.GetShowStories(ctx)
	if err != nil {
		return nil, err
	}
	return s.fetchPaginatedItems(ctx, ids, page, pageSize)
}

func (s *ItemService) ListJobStories(ctx context.Context, page int, pageSize int) ([]domain.Item, error) {
	ids, err := s.hnAPIAdapter.GetJobStories(ctx)
	if err != nil {
		return nil, err
	}
	return s.fetchPaginatedItems(ctx, ids, page, pageSize)
}

func (s *ItemService) fetchPaginatedItems(ctx context.Context, ids []int, page int, pageSize int) ([]domain.Item, error) {
	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= len(ids) {
		return []domain.Item{}, nil // Empty slice for out-of-bounds page
	}
	if end > len(ids) {
		end = len(ids)
	}

	var items []domain.Item
	for _, id := range ids[start:end] {
		item, err := s.hnAPIAdapter.GetItem(ctx, id)
		if err != nil {
			// Consider logging the error and potentially skipping the item
			continue
		}
		items = append(items, item)
	}
	return items, nil
}
