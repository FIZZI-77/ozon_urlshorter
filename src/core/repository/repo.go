package repository

import (
	"context"
	"database/sql"
	"ozon/models"
)

type UrlRepository interface {
	Create(ctx context.Context, link *models.Link) error
	GetByShort(ctx context.Context, shortUrl string) (*models.Link, error)
	GetByURL(ctx context.Context, url string) (*models.Link, error)
}

type Repository struct {
	UrlRepository
}

func NewPostgresRepository(db *sql.DB) *Repository {
	return &Repository{
		UrlRepository: NewUrlRepo(db),
	}
}
func NewMemoryRepository() *Repository {
	return &Repository{
		UrlRepository: NewMemoryRepo(),
	}
}
