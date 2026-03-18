package service

import (
	"context"
	"ozon/src/core/repository"
)

type UlrService interface {
	CreateShortUrl(ctx context.Context, url string) (string, error)
	GetOriginalUrl(ctx context.Context, shortUrl string) (string, error)
}

type Service struct {
	UlrService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UlrService: NewUrlShorterService(repo),
	}
}
