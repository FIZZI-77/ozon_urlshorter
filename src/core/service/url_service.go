package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"ozon/models"
	"ozon/src/core/repository"
)

type UrlShorterService struct {
	repo *repository.Repository
}

func NewUrlShorterService(repo *repository.Repository) *UrlShorterService {
	return &UrlShorterService{repo: repo}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const length = 10

func (u *UrlShorterService) CreateShortUrl(ctx context.Context, url string) (string, error) {

	existing, err := u.repo.GetByURL(ctx, url)
	if err == nil {
		return existing.Short, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	shortUrl, err := generateShortUrl()
	if err != nil {
		return "", err
	}
	link := &models.Link{
		URL:   url,
		Short: shortUrl,
	}

	err = u.repo.Create(ctx, link)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (u *UrlShorterService) GetOriginalUrl(ctx context.Context, shortUrl string) (string, error) {
	link, err := u.repo.GetByShort(ctx, shortUrl)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", models.ErrNotFound
		}
		return "", err
	}
	return link.URL, nil

}

func generateShortUrl() (string, error) {
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("service: url_service: generateShortUrl(): error generating random number: %v", err)
		}
		b[i] = charset[n.Int64()]
	}
	return string(b), nil
}
