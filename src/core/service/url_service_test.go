package service

import (
	"context"
	"ozon/src/core/repository"
	"testing"
)

func TestCreateShortUrl_New(t *testing.T) {
	repo := repository.NewMemoryRepository()
	service := NewUrlShorterService(repo)

	url := "https://google.com"

	short, err := service.CreateShortUrl(context.Background(), url)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(short) != 10 {
		t.Fatalf("expected short url length 10, got %d", len(short))
	}
}

func TestCreateShortUrl_SameURL(t *testing.T) {
	repo := repository.NewMemoryRepository()
	service := NewUrlShorterService(repo)

	url := "https://google.com"

	short1, _ := service.CreateShortUrl(context.Background(), url)
	short2, _ := service.CreateShortUrl(context.Background(), url)

	if short1 != short2 {
		t.Fatalf("expected same short url, got %s and %s", short1, short2)
	}
}

func TestGetOriginalUrl(t *testing.T) {
	repo := repository.NewMemoryRepository()
	service := NewUrlShorterService(repo)

	url := "https://google.com"

	short, _ := service.CreateShortUrl(context.Background(), url)

	original, err := service.GetOriginalUrl(context.Background(), short)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if original != url {
		t.Fatalf("expected %s, got %s", url, original)
	}
}

func TestGetOriginalUrl_NotFound(t *testing.T) {
	repo := repository.NewMemoryRepository()
	service := NewUrlShorterService(repo)

	_, err := service.GetOriginalUrl(context.Background(), "not_exist")

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
