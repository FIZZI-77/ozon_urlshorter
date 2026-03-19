package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ozon/models"
)

type UrlRepo struct {
	db *sql.DB
}

func NewUrlRepo(db *sql.DB) *UrlRepo {
	return &UrlRepo{db: db}
}

func (u *UrlRepo) Create(ctx context.Context, link *models.Link) error {
	const query = `INSERT INTO urls(base_url, short_url) VALUES ($1, $2)`

	_, err := u.db.ExecContext(ctx, query, link.URL, link.Short)

	if err != nil {
		return fmt.Errorf("repo: urlRepo: Create(): cant create short url: %w", err)
	}
	return nil

}
func (u *UrlRepo) GetByShort(ctx context.Context, shortUrl string) (*models.Link, error) {
	const query = `SELECT id, base_url, short_url FROM urls WHERE short_url = $1`

	link := &models.Link{}

	err := u.db.QueryRowContext(ctx, query, shortUrl).Scan(
		&link.ID,
		&link.URL,
		&link.Short,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("repo: urlRepo: GetByShort(): %w", err)
	}
	return link, nil

}
func (u *UrlRepo) GetByURL(ctx context.Context, url string) (*models.Link, error) {
	const query = `SELECT id, base_url, short_url FROM urls WHERE base_url = $1`

	link := &models.Link{}

	err := u.db.QueryRowContext(ctx, query, url).Scan(
		&link.ID,
		&link.URL,
		&link.Short,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("repo: urlRepo: GetByURL(): %w", err)
	}
	return link, nil
}
