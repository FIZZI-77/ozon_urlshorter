package repository

import (
	"context"
	"ozon/models"
	"testing"
)

func TestMemoryRepo_CreateAndGet(t *testing.T) {
	repo := NewMemoryRepo()

	link := &models.Link{
		URL:   "https://google.com",
		Short: "abc1234567",
	}

	err := repo.Create(context.Background(), link)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	byShort, _ := repo.GetByShort(context.Background(), link.Short)
	if byShort.URL != link.URL {
		t.Fatalf("expected %s, got %s", link.URL, byShort.URL)
	}

	byURL, _ := repo.GetByURL(context.Background(), link.URL)
	if byURL.Short != link.Short {
		t.Fatalf("expected %s, got %s", link.Short, byURL.Short)
	}
}

func TestMemoryRepo_NotFound(t *testing.T) {
	repo := NewMemoryRepo()

	_, err := repo.GetByShort(context.Background(), "nope")
	if err == nil {
		t.Fatalf("expected error")
	}

	_, err = repo.GetByURL(context.Background(), "nope")
	if err == nil {
		t.Fatalf("expected error")
	}
}
