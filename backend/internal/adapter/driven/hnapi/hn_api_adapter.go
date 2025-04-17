package hnapi

import (
	"context"
	"encoding/json"
	"fmt"
	"hackernews-clone-backend/internal/core/domain"
	"net/http"
	"strings"

	"time"
)

const (
	baseURL = "https://hacker-news.firebaseio.com/v0"
)

type HNAPIAdapter struct {
	client *http.Client
}

func NewHNAPIAdapter() *HNAPIAdapter {
	return &HNAPIAdapter{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (a *HNAPIAdapter) Get(id int) (*domain.Item, error) {
	url := fmt.Sprintf("%s/item/%d.json", baseURL, id)
	resp, err := a.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch item %d from HN API: %w", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HN API returned non-OK status for item %d: %d", id, resp.StatusCode)
	}

	var item domain.Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return nil, fmt.Errorf("failed to decode item %d from HN API: %w", id, err)
	}
	item.CreatedAt = time.Unix(item.Time, 0)
	item.UpdatedAt = time.Unix(item.Time, 0)
	return &item, nil
}

func (a *HNAPIAdapter) GetAllByType(itemType domain.ItemType, limit int) ([]domain.Item, error) {
	var endpoint string
	switch itemType {
	case domain.TypeStory:
		endpoint = "topstories.json" // Default to top stories
	case domain.TypeJob:
		endpoint = "jobstories.json"
	default:
		return nil, fmt.Errorf("unsupported item type: %s for list fetching from HN API", itemType)
	}

	url := fmt.Sprintf("%s/%s", baseURL, endpoint)
	resp, err := a.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s from HN API: %w", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HN API returned non-OK status for %s: %d", endpoint, resp.StatusCode)
	}

	var ids []int
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, fmt.Errorf("failed to decode %s IDs from HN API: %w", endpoint, err)
	}

	var items []domain.Item
	count := 0
	for _, id := range ids {
		if count >= limit {
			break
		}
		item, err := a.Get(id)
		if err != nil {
			fmt.Printf("Error fetching item %d: %v\n", id, err)
			continue
		}

		// Special handling for 'ask' and 'show' stories as they are just stories with specific titles
		if itemType == domain.TypeStory {
			if strings.HasPrefix(*item.Title, "Ask HN") || strings.HasPrefix(*item.Title, "Show HN") {
				items = append(items, *item)
				count++
			}
		} else if item.Type == itemType {
			items = append(items, *item)
			count++
		}
	}
	return items, nil
}

func (a *HNAPIAdapter) GetAskStories(ctx context.Context) ([]int, error) {
	url := fmt.Sprintf("%s/askstories.json", baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ask stories: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch ask stories, status code: %d", resp.StatusCode)
	}

	var ids []int
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, fmt.Errorf("failed to decode ask stories: %w", err)
	}

	return ids, nil
}

func (a *HNAPIAdapter) GetItem(ctx context.Context, id int) (domain.Item, error) {
	url := fmt.Sprintf("%s/item/%d.json", baseURL, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return domain.Item{}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return domain.Item{}, fmt.Errorf("failed to fetch item: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Item{}, fmt.Errorf("failed to fetch item, status code: %d", resp.StatusCode)
	}

	var item domain.Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return domain.Item{}, fmt.Errorf("failed to decode item: %w", err)
	}

	return item, nil
}

func (a *HNAPIAdapter) GetJobStories(ctx context.Context) ([]int, error) {
	url := fmt.Sprintf("%s/jobstories.json", baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch job stories: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch job stories, status code: %d", resp.StatusCode)
	}

	var ids []int
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, fmt.Errorf("failed to decode job stories: %w", err)
	}

	return ids, nil
}

func (a *HNAPIAdapter) GetNewStories(ctx context.Context) ([]int, error) {
	url := fmt.Sprintf("%s/newstories.json", baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch new stories: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch new stories, status code: %d", resp.StatusCode)
	}

	var ids []int
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, fmt.Errorf("failed to decode new stories: %w", err)
	}

	return ids, nil
}

func (a *HNAPIAdapter) GetShowStories(ctx context.Context) ([]int, error) {
	url := fmt.Sprintf("%s/showstories.json", baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch show stories: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch show stories, status code: %d", resp.StatusCode)
	}

	var ids []int
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, fmt.Errorf("failed to decode show stories: %w", err)
	}

	return ids, nil
}

func (a *HNAPIAdapter) GetTopStories(ctx context.Context) ([]int, error) {
	url := fmt.Sprintf("%s/topstories.json", baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch top stories: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch top stories, status code: %d", resp.StatusCode)
	}

	var ids []int
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, fmt.Errorf("failed to decode top stories: %w", err)
	}

	return ids, nil
}
