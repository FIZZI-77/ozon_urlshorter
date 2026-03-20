package repository

import (
	"context"
	"database/sql"
	"ozon/models"
	"sync"
)

type MemoryStore struct {
	mu              sync.RWMutex
	originalToShort map[string]string
	shortToOriginal map[string]string
}

func NewMemoryRepo() *MemoryStore {
	return &MemoryStore{
		originalToShort: make(map[string]string),
		shortToOriginal: make(map[string]string),
	}
}

func (m *MemoryStore) Create(ctx context.Context, link *models.Link) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.shortToOriginal[link.Short] = link.URL
	m.originalToShort[link.URL] = link.Short

	return nil

}

func (m *MemoryStore) GetByShort(ctx context.Context, shortUrl string) (*models.Link, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	url, ok := m.shortToOriginal[shortUrl]
	if !ok {
		return nil, sql.ErrNoRows
	}

	link := &models.Link{
		Short: shortUrl,
		URL:   url,
	}
	return link, nil

}

func (m *MemoryStore) GetByURL(ctx context.Context, url string) (*models.Link, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	short, ok := m.originalToShort[url]
	if !ok {
		return nil, sql.ErrNoRows
	}
	link := &models.Link{
		Short: short,
		URL:   url,
	}
	return link, nil
}
